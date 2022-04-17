package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	IsActive bool   `json:"isActive"`
}
