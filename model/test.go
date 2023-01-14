package model

type Testcase struct {
	Tid    int    `json:"tid"`
	Pid    string `json:"pid"`
	Uid    string `json:"uid"`
	Input  string `json:"input"`
	Output string `json:"output"`
}
