package Model

import pb "UserServices/proto"

type UserModel struct {
	ID          string `bson:"_id"`
	Email       string `bson:"email"`
	Username    string `bson:"username"`
	Image       string `bson:"image"`
	PhoneNumber string `bson:"phoneNumber"`
	About       string `bson:"about"`
	Birthdate   int64  `bson:"birthdate"`
	Password    string `bson:"password"`
	Role        string `bson:"role"`
	CreatedAt   int64  `bson:"createdAt"`
	Validate    bool   `bson:"validate"`
}

func DocumentToUser(data *UserModel) *pb.User {
	return &pb.User{
		Id:          data.ID,
		Email:       data.Email,
		Username:    data.Username,
		Image:       data.Image,
		PhoneNumber: data.PhoneNumber,
		About:       data.About,
		Birthdate:   data.Birthdate,
		Password:    data.Password,
		Role:        data.Role,
		CreatedAt:   data.CreatedAt,
		Validate:    data.Validate,
	}
}
