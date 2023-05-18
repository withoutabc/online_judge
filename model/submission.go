package model

import (
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
	"log"
)

type Submission struct {
	SubmissionId int64  `json:"submission_id" form:"submission_id" gorm:"primarykey"`
	ProblemId    int64  `json:"problem_id" form:"problem_id" binding:"required" gorm:"type:int;not null"`
	UserId       int64  `json:"user_id" form:"user_id" binding:"required" gorm:"type:int;not null"`
	Code         string `json:"code" form:"code" binding:"required" gorm:"type:longblob;not null"`
	Language     string `json:"language" form:"language" binding:"required" gorm:"type:varchar(20);not null"`
	Status       string `json:"status" form:"status" gorm:"check_constraint:status IN('待测评','编译错误','答案错误','正确','运行超时');not null;type:varchar(20)"`
}

// BeforeCreate uses snowflake to generate an ID.
func (s *Submission) BeforeCreate(_ *gorm.DB) (err error) {
	// skip if the accountID already set.
	if s.SubmissionId != 0 {
		return nil
	}
	sf, err := snowflake.NewNode(0)
	if err != nil {
		log.Fatalf("generate id failed: %s", err.Error())
		return err
	}
	s.SubmissionId = sf.Generate().Int64()
	s.Status = "待评测"
	return nil
}
