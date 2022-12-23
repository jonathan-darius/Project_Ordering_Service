package Model

type VerificationMSG struct {
	Email string `json:"email" bson:"email"`
	Token string `json:"token" bson:"token"`
}

type Verification struct {
	UserID    string `json:"id" bson:"_id"`
	Email     string `json:"email" bson:"email"`
	Token     string `json:"token" bson:"token"`
	CreatedAt int64  `json:"createdAt" bson:"createdAt"`
	ExpiredAt int64  `json:"expiredAt" bson:"expiredAt"`
}
