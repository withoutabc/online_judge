package model

type RespProblem struct {
	Status int       `json:"status"`
	Info   string    `json:"info"`
	Data   []Problem `json:"problem"`
}

type RespSubmission struct {
	Status int          `json:"status"`
	Info   string       `json:"info"`
	Data   []Submission `json:"submission"`
}

type RespTestcase struct {
	Status int        `json:"status"`
	Info   string     `json:"info"`
	Data   []Testcase `json:"testcase"`
}

type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type RespToken struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
	Data   Token  `json:"data"`
}
