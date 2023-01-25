package models

type UserFan struct {
	ID     uint `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	UserId uint `gorm:"type:uint;not null" json:"user_id"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	FanId  uint `gorm:"type:uint;not null" json:"fan_id"`
	Fan    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (UserFan) CreateFakeUserFan() UserFan {
	userFan := UserFan{
		UserId: 2,
		FanId:  1,
	}
	return userFan
}
