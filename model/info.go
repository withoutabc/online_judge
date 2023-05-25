package model

import (
	"gorm.io/gorm"
	"time"
)

type Info struct {
	Name     string `json:"name"  form:"name" binding:"required" gorm:"not null;default:''"`
	UserId   int64  `json:"user_id" form:"user_id" binding:"required" gorm:"not null"`
	Nickname string `json:"nickname," form:"nickname" binding:"required" gorm:"not null;default:''"`
	Gender   int    `json:"gender" form:"gender" binding:"required" gorm:"not null;default:0"`
	Year     int    `json:"year" form:"year" binding:"required" gorm:"not null;default:0"`
	Month    int    `json:"month" form:"month" binding:"required" gorm:"not null;default:0"`
	Day      int    `json:"day" form:"day" binding:"required" gorm:"not null;default:0"`
	Email    string `json:"email" form:"email" binding:"required" gorm:"not null;default:''"`
	Correct  int    `json:"correct" form:"correct" binding:"-" gorm:"not null;default:0"`
	Score    int    `json:"score" form:"score" binding:"-" gorm:"not null;default:0"`
}

func (i *Info) BeforeCreate(_ *gorm.DB) (err error) {
	i.Gender = 0
	i.Year = time.Now().Year()
	i.Month = int(time.Now().Month())
	i.Day = time.Now().Day()
	i.Email = ""
	i.Score = 0
	i.Correct = 0
	i.Nickname = ""
	i.Name = ""
	return nil
}
