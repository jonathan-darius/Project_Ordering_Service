package models

type UserShow struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Image       string `json:"image"`
	PhoneNumber string `json:"phoneNumber"`
	About       string `json:"about"`
	Birthdate   int64  `json:"birthdate"`
	Role        string `json:"role"`
}

type UpdateModel struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phoneNumber"`
	About       string `json:"about"`
	Birthdate   int64  `json:"birthdate"`
}
