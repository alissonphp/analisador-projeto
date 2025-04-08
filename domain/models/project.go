package models

import "github.com/alissonphp/analisador-projeto/domain/values"

type Project struct {
	Name       string `json:"name"`
	Squad      string `json:"squad"`
	Identifier string `json:"identifier"`
	Source     string `json:"source"`
}

func (p Project) GetApiCall() ApiCall {
	return ApiCall{
		Host: values.SONAR_HOST[p.Source],
		Key:  values.SONAR_KEY[p.Source],
	}
}
