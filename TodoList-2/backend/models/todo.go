package models

type TODO struct {
	Id     int    `json:"id"`
	Todo   string `json:"todo"`
	IsDone bool   `json:"isDone"`
}