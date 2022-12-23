package utils

import (
	gRPCFunc "WebService/App/gRPC_Configs/User"
	pb "WebService/App/gRPC_Configs/proto/UserService"
	"WebService/App/models"
)

func CheckAuthUser(email string, uid string) (bool, error) {
	getUser, err := gRPCFunc.SearchUserByID(uid)
	if err != nil {
		return false, err
	}
	if getUser.Email == email {
		return true, nil
	}
	return false, nil
}

func CheckAdmin(role string) bool {
	if role == "admin" {
		return true
	}
	return false
}

func CheckVerifyAccount(uID string) bool {
	getUser, err := gRPCFunc.SearchUserByID(uID)
	if err != nil {
		return false
	}
	return getUser.Validate
}

func ResponseUser(raw *pb.User) *models.UserShow {
	return &models.UserShow{
		ID:          raw.Id,
		Email:       raw.Email,
		Image:       raw.Image,
		Username:    raw.Username,
		PhoneNumber: raw.PhoneNumber,
		About:       raw.About,
		Birthdate:   raw.Birthdate,
		Role:        raw.Role,
	}
}
