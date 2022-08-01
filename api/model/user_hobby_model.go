package model

//
type UserHobby struct {
	Base
	UserID  string `gorm:"type:size:36;index"`
	User    User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	HobbyID string `gorm:"type:size:36;index"`
	Hobby   Hobby  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
