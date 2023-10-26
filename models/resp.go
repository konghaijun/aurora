package models

import "auroralab/repository"

type SelectResponse struct {
	StatusCode int                    `json:"status_code"`
	StatusMsg  string                 `json:"status_msg"`
	Dep        repository.Departments `json:"dep"`
}

type BaseResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}
type AnswerResponse struct {
	Resp    BaseResponse `json:"resp"`
	Answers string       `json:"answers"`
}

type Answer struct {
	Prompt         string `json:"prompt"`
	UserID         string `json:"userId"`
	Network        bool   `json:"network"`
	System         string `json:"system"`
	WithoutContext bool   `json:"withoutContext"`
	Stream         bool   `json:"stream"`
}
