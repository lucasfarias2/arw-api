package models

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type FullUser struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
