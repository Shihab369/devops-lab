package models

import "time"

type Result struct {
	Name      string      `json:"name"`
	Status    string      `json:"status"`
	Timestamp time.Time   `json:"timestamp"`
	Data      interface{} `json:"data"`
}
