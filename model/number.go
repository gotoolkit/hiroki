package model

import "github.com/jinzhu/gorm"

type Number struct {
    gorm.Model
    Reds      string `gorm:"type:char(17);not null"`
    Beginning string `gorm:"type:char(2);not null"`
    Ending    string `gorm:"type:char(2);not null"`
}

func (Number) TableName() string {
    return "number"
}
