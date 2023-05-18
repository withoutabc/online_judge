package model

type Testcase struct {
	TestId    int    `json:"test_id" form:"test_id"`
	ProblemId string `json:"problem_id" form:"test_id"`
	UserId    string `json:"user_id" form:"user_id"`
	Input     string `json:"input" form:"input"`
	Output    string `json:"output" form:"output"`
}
