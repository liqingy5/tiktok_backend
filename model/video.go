package model

type Video struct {
	ID           uint   `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	UserId       uint   `gorm:"type:uint;not null" json:"user_id"`
	User         User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	VideoUrl     string `gorm:"type:varchar(200);not null" json:"video_url"`
	CoverUrl     string `gorm:"type:varchar(200);not null" json:"cover_url"`
	VideoDesc    string `gorm:"type:TEXT;not null" json:"video_desc"`
	CreateTime   int64  `gorm:"autoCreateTime" json:"create_time"`
	LikeCount    int    `gorm:"type:int;not null;default:0" json:"like_count"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	Title        string `gorm:"type:TEXT;not null" json:"title"`
}

func (Video) CreateFakeVideo() Video {
	video := Video{
		UserId:    1,
		VideoUrl:  "./public/video/bear.mp4",
		CoverUrl:  "./public/cover/bear.jpg",
		VideoDesc: "This is a bear",
	}
	return video
}
