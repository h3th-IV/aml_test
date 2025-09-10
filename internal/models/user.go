package models

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"firstname"`
	Email   string `json:"email"`
	Gender  string `json:"gender"`
	Dob     string `json:"dob"`
	Address string `json:"address"`
}
