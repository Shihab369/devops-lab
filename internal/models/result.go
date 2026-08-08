package models

type Result struct {
	Name string                 `json:"name"`
	Data map[string]interface{} `json:"data"`
}
