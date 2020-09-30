package model

type NLP struct {
	Active bool   `json:"active" yml:"active"`
	Name   string `json:"name" yml:"name"`
	Value  int64  `json:"value" yml:"value"`
        ID string      `json:"id" yml:"id"` 
}
