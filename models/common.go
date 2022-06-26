package models

type Response struct {
	Error   string
	Message string
	Data    interface{}
}

// ApplicationQueryParamModel ...
type ApplicationQueryParamModel struct {
	Search string `json:"search"`
	Offset int    `json:"offset" default:"0"`
	Limit  int    `json:"limit" default:"10"`
}
