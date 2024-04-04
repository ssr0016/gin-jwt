package model

type Users struct {
	Id       int    `gorm:"type:int;primaryKey"`
	Username string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}
