package entity

type TaskNumberByStatusEntity struct {
	Status string `json:"status"`
	Number int    `json:"number"`
}
