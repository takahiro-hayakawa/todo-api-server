package model

type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	LimitDate int    `json:"limitDate"`
	Status    int    `json:"status"`
}
