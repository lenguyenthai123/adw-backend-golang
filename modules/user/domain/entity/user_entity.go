package entity

import "backend-golang/core"

type UserEntity struct {
	core.BaseEntity `json:",inline"`
	FullName        string `json:"full_name" gorm:"not null;"`
	Email           string `json:"email" gorm:"uniqueIndex;not null;"`
	Phone           string `json:"phone" gorm:"uniqueIndex;default:null"`
	Status          string `json:"status" gorm:"not null;default:UNVERIFIED"`
	OTP             string `json:"-" gorm:"default:null"`
	Password        string `json:"-" gorm:"not null;"`
	AccessToken     string `json:"-" gorm:"-"`
	RefreshToken    string `json:"-" gorm:"-"`
}

func (UserEntity) TableName() string {
	return "users"
}
