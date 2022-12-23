package models

type FormLoginData struct {
	Email    string `form:"email" binding:"required" json:"email"`
	Password string `form:"password" binding:"required" json:"password"`
}

type RegisterData struct {
	Email       string `json:"email" validate:"required"`
	Username    string `json:"username" validate:"required"`
	PhoneNumber string `json:"phoneNumber"`
	About       string `json:"about"`
	Birthdate   int64  `json:"birthdate"`
	Password    string `json:"password" validate:"required,min=6"`
	Role        string `json:"role"`
}

type AccountVerificationData struct {
	Email string `json:"email"`
	Token string `json:"token" validate:"required"`
}
