package models

type ApiCall struct {
	Host string `json:"host"`
	Key  string `json:"key"`
}

type APIResponse struct {
	Component Component `json:"component"`
}
