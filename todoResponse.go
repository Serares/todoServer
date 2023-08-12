package main

import (
	"encoding/json"
	"time"

	todo "github.com/Serares/todoCli"
)

type todoResponse struct {
	Results todo.List `json:"results"`
}

// associate a custom marchal to the response
// this helps with adding fields dynamically at the time of using them
func (r *todoResponse) MarshalJSON() ([]byte, error) {
	resp := struct {
		Results       todo.List `json:"results"`
		Date          int64     `json:"date"`
		TotalResulsts int       `json:"total_results"`
	}{
		Results:       r.Results,
		Date:          time.Now().Unix(),
		TotalResulsts: len(r.Results),
	}

	return json.Marshal(resp)
}
