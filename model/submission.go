package model

type Submission struct {
	Sid      int    `json:"sid"`
	Pid      string `json:"pid"`
	Uid      string `json:"uid"`
	Code     string `json:"code"`
	Language string `json:"language"`
	Status   string `json:"status"`
}
