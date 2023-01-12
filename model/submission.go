package model

type Submission struct {
	Sid    int    `json:"sid" form:"sid" `
	Pid    string `json:"pid" form:"pid"`
	Uid    string `json:"uid" form:"uid" `
	Status string `json:"status" form:"status"`
	Code   string `json:"code" form:"code" `
}
