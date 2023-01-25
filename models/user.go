package models

type User struct {
	ID                uint   `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Username          string `gorm:"type:varchar(20);not null;unique" json:"username"`
	Password          string `gorm:"type:varchar(255);not null" json:"password"`
	LikesReceiveCount int    `gorm:"type:int;not null;default:0" json:"likes_receive_count"`
	FollowerCount     int    `gorm:"type:int;not null;default:0" json:"follower_count"`
	FollowingCount    int    `gorm:"type:int;not null;default:0" json:"following_count"`
}

func (User) CreateFakeUser() []User {
	users := []User{
		{
			Username: "zhanglei",
			Password: "douyin",
		},
		{
			Username: "icc",
			Password: "996007",
		},
	}
	return users
}
