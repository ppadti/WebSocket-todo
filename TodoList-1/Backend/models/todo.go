package models

type TODO struct {
    Id          int    `json:"id"`
    Action      string `json:"action"`
    Todo        string `json:"todo"`
    IsDone      bool   `json:"isDone,omitempty"`
}