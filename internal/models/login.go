package models

type Login struct {
	Token  string `json:"token"`
	UserID int    `json:"user_id"`
}
