package config

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"tiktok_backend/model"
)

type Configuration struct {
	DB struct {
		User     string `json:"db_user"`
		Password string `json:"db_pass"`
		Host     string `json:"db_host"`
		Port     int    `json:"db_port"`
		HostName string `json:"db_name"`
	} `json:"db"`
}

// global db object
var _db *gorm.DB

func Connect() {
	file, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	// Unmarshal the JSON data into a Config struct
	var config Configuration
	json.Unmarshal(file, &config)

	// Build the connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.HostName)
	db, err := gorm.Open(mysql.Open(dsn))

	_db = db

	if err != nil {
		panic("SQL connection failed: " + err.Error())
	}
	sqlDB, _ := _db.DB()
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Could not establish connection with the database", err)
	}
	fmt.Println("Connection success")

	//设置数据库连接池参数 reference from https://www.tizi365.com/archives/15.html
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭

}

func GetDB() *gorm.DB {
	return _db
}

func InitTables() {
	if _db == nil {
		log.Fatal("db is nil")
	}
	err := _db.AutoMigrate(&model.User{}, &model.Comment{}, &model.Video{}, &model.UserFan{}, &model.UserLikeVideo{})
	if err != nil {
		log.Fatal("create tables failed", err)
	}
	fmt.Println("create tables success")
}

func InitFakeData() {
	if _db == nil {
		log.Fatal("db is nil")
	}
	// create fake user
	users := model.User{}.CreateFakeUser()
	_db.Create(&users)
	// create fake video
	videos := model.Video{}.CreateFakeVideo()
	_db.Create(&videos)
	// create fake comment
	comments := model.Comment{}.CreateFakeComment()
	_db.Create(&comments)
	// create fake user fan
	userFans := model.UserFan{}.CreateFakeUserFan()
	_db.Create(&userFans)
	// create fake user like video
	userLikeVideos := model.UserLikeVideo{}.CreateFakeUserLikeVideo()
	_db.Create(&userLikeVideos)
}

func CheckFakeData() {
	if _db == nil {
		log.Fatal("db is nil")
	}
	var users []model.User
	_db.Find(&users)
	for _, user := range users {
		fmt.Println(user)
	}
	var videos []model.Video
	_db.Find(&videos)
	for _, video := range videos {
		fmt.Println(video.UserId)
	}
	var comments []model.Comment
	_db.Find(&comments)
	for _, comment := range comments {
		fmt.Println(comment.UserId)
	}

}
