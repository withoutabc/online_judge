package model

type Problem struct {
	Pid               int     `json:"pid"`
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	DescriptionInput  string  `json:"description_input"`
	DescriptionOutput string  `json:"description_output"`
	SampleInput       string  `json:"sample_input"`
	SampleOutput      string  `json:"sample_output"`
	TimeLimit         float64 `json:"time_limit"`
	MemoryLimit       float64 `json:"memory_limit"`
	Uid               string  `json:"uid"`
}
