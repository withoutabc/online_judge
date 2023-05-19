package model

import (
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
	"log"
	"time"
)

type Problem struct {
	ProblemId         int64  `json:"problem_id" form:"problem_id" binding:"-" gorm:"primarykey"`
	UserId            int64  `json:"user_id" form:"user_id" binding:"required" gorm:"not null"`
	Title             string `json:"title" form:"title" binding:"required" gorm:"type:varchar(60);not null"`
	Description       string `json:"description" form:"description" binding:"required" gorm:"type:longtext;not null"`
	DescriptionInput  string `json:"description_input" form:"description_input" binding:"required" gorm:"type:longtext;not null"`
	DescriptionOutput string `json:"description_output" form:"description_output" binding:"required" gorm:"type:longtext;not null"`
	SampleInput       string `json:"sample_input" form:"sample_input" binding:"required" gorm:"type:text;not null"`
	SampleOutput      string `json:"sample_output" form:"sample_output" binding:"required" gorm:"type:text;not null"`
	Level             string `json:"level" form:"level" binding:"required;oneof='极易','容易','中等','困难','极难'" gorm:"check_constraint:level IN('极易','容易','中等','困难','极难');not null;type:varchar(20)"`
	UpdateTime        string `json:"update_time" form:"update_time" binding:"-" gorm:"type:varchar(100)"`
	Submit            int64  `json:"submit" form:"submit" binding:"-" gorm:"default:0"`
	Correct           int64  `json:"correct" form:"correct" binding:"-" gorm:"default:0"`
}

type ReqSearchProblem struct {
	UserId    int64  `json:"user_id" form:"user_id"`
	ProblemId int64  `json:"problem_id" form:"problem_id"`
	Keyword   string `json:"keyword" form:"keyword"`
	Level     string `json:"level" form:"level"`
	From      string `json:"from" form:"from"`
	To        string `json:"to" form:"to"`
}

// BeforeCreate uses snowflake to generate an ID.
func (p *Problem) BeforeCreate(_ *gorm.DB) (err error) {
	// skip if the accountID already set.
	if p.ProblemId != 0 {
		return nil
	}
	sf, err := snowflake.NewNode(0)
	if err != nil {
		log.Fatalf("generate id failed: %s", err.Error())
		return err
	}
	p.ProblemId = sf.Generate().Int64()
	p.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	return nil
}

func (p *Problem) BeforeUpdate(_ *gorm.DB) (err error) {
	p.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	return nil
}
