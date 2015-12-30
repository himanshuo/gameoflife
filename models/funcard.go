package models

type FunCard struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Progress Progress `json:"progress"`
}
