package model

import (
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
	"log"
)

type Testcase struct {
	TestId    int64  `json:"test_id" form:"test_id" binding:"-" gorm:"primarykey"`
	ProblemId int64  `json:"problem_id" form:"problem_id" binding:"required" gorm:"not null"`
	UserId    int64  `json:"user_id" form:"user_id" binding:"required" gorm:"not null"`
	Input     string `json:"input" form:"input" binding:"required" gorm:"not null"`
	Output    string `json:"output" form:"output" binding:"required" gorm:"not null"`
}

// BeforeCreate uses snowflake to generate an ID.
func (t *Testcase) BeforeCreate(_ *gorm.DB) (err error) {
	// skip if the accountID already set.
	if t.TestId != 0 {
		return nil
	}
	sf, err := snowflake.NewNode(0)
	if err != nil {
		log.Fatalf("generate id failed: %s", err.Error())
		return err
	}
	t.TestId = sf.Generate().Int64()
	return nil
}
