package model

type Testcase struct {
	Tid    int    `json:"tid"`
	Pid    string `json:"pid"`
	Input  string `json:"input"`
	Output string `json:"output"`
	Score  string `json:"score"`
}
