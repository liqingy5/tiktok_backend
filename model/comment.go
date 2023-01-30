package model

type Comment struct {
	ID          uint   `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	UserId      uint   `gorm:"type:uint;not null" json:"user_id"`
	User        User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	VideoId     uint   `gorm:"type:uint;not null" json:"video_id"`
	Video       Video  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CommentText string `gorm:"type:TEXT;not null" json:"comment_text"`
	CreateTime  int64  `gorm:"autoCreateTime" json:"create_time"`
}

func (Comment) CreateFakeComment() Comment {
	comment := Comment{
		UserId:      2,
		VideoId:     1,
		CommentText: "This is a fake comment",
	}
	return comment
}
