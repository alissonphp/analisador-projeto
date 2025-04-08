package models

type Measure struct {
	Metric    string `json:"metric"`
	Value     string `json:"value"`
	BestValue *bool  `json:"bestValue,omitempty"`
}
