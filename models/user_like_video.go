package models

type UserLikeVideo struct {
	ID      uint  `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	UserId  uint  `gorm:"type:uint;not null" json:"user_id"`
	User    User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	VideoId uint  `gorm:"type:uint;not null" json:"video_id"`
	Video   Video `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (UserLikeVideo) CreateFakeUserLikeVideo() UserLikeVideo {
	userLikeVideo := UserLikeVideo{
		UserId:  2,
		VideoId: 1,
	}
	return userLikeVideo
}
