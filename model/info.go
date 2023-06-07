package model

import (
	"gorm.io/gorm"
	"time"
)

type Info struct {
	Name     string `json:"name"  form:"name" gorm:"not null;default:''"`
	UserId   int64  `json:"user_id" form:"user_id" binding:"required" gorm:"not null"`
	Nickname string `json:"nickname," form:"nickname"  gorm:"not null;default:''"`
	Gender   int    `json:"gender" form:"gender" gorm:"not null"`
	Year     int    `json:"year" form:"year"  gorm:"not null;default:0"`
	Month    int    `json:"month" form:"month"  gorm:"not null;default:0"`
	Day      int    `json:"day" form:"day" gorm:"not null;default:0"`
	Email    string `json:"email" form:"email"  gorm:"not null;default:''"`
	Correct  int    `json:"correct" form:"correct" binding:"-" gorm:"not null;default:0"`
	Score    int    `json:"score" form:"score" binding:"-" gorm:"not null;default:0"`
}

func (i *Info) BeforeCreate(_ *gorm.DB) (err error) {
	i.Gender = 3
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
