package gRPCFunc

import (
	"WebService/App/gRPC_Configs/proto/UserService"
	"WebService/App/models"
	"bufio"
	"context"
	"io"
	"mime/multipart"
	"path/filepath"
	"time"
)

func UserLogin(data models.FormLoginData) (*UserService.LoginResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	loginReq, err := Client.Login(ctx, &UserService.LoginRequest{
		Email:    data.Email,
		Password: data.Password,
	})
	defer cancel()
	if err != nil {
		return nil, err
	}
	return loginReq, nil
}

func RegisterUser(data models.RegisterData) (*UserService.UserId, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	registerRes, err := Client.RegisterUser(ctx, &UserService.User{
		Email:       data.Email,
		Username:    data.Username,
		PhoneNumber: data.PhoneNumber,
		About:       data.About,
		Birthdate:   data.Birthdate,
		Password:    data.Password,
		Role:        data.Role,
	})
	defer cancel()
	if err != nil {
		return nil, err
	}
	return registerRes, nil
}

func GetAccountVerification(email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_, err := Client.GetVerification(ctx, &UserService.UserEmail{Email: email})
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}

func ValidateToken(token string) (*UserService.ValidateTokenResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	res, err := Client.ValidateToken(ctx,
		&UserService.ValidateTokenRequest{
			Token: token,
		})
	defer cancel()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func SearchUserByID(id string) (*UserService.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	res, err := Client.GetUserByID(ctx, &UserService.UserId{Id: id})
	defer cancel()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func AccountVerification(mail string, verifyToken string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_, err := Client.VerificationUser(ctx, &UserService.ValidateRequest{
		Email: mail,
		Token: verifyToken,
	})
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}

func ListUser(page int32, recordPerPage int32) ([]*models.UserShow, error) {
	ctx := context.Background()
	stream, err := Client.ListUser(ctx, &UserService.Pagination{
		Page:          page,
		RecordPerPage: recordPerPage,
	})
	if err != nil {
		return nil, err
	}
	var listUser []*models.UserShow
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		listUser = append(listUser, &models.UserShow{
			ID:          res.Id,
			Email:       res.Email,
			Username:    res.Username,
			PhoneNumber: res.PhoneNumber,
			About:       res.About,
			Birthdate:   res.Birthdate,
			Role:        res.Role,
		})
	}
	return listUser, nil
}

func DelUser(uid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_, err := Client.DeleteUser(ctx, &UserService.UserId{Id: uid})
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(data *models.UpdateModel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_, err := Client.UpdateUser(ctx, &UserService.User{
		Id:          data.ID,
		Username:    data.Username,
		PhoneNumber: data.PhoneNumber,
		About:       data.About,
		Birthdate:   data.Birthdate,
	})
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}

func SendImage(file *multipart.FileHeader, uid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stream, err := Client.UploadImage(ctx)
	if err != nil {
		return err
	}

	req := &UserService.UploadImageRequest{
		Data: &UserService.UploadImageRequest_Info{
			Info: &UserService.ImageInfo{
				UserId:    uid,
				ImageType: filepath.Ext(filepath.Base(file.Filename)),
			},
		},
	}

	if err = stream.Send(req); err != nil {
		return err
	}

	rawData, err := file.Open()
	reader := bufio.NewReader(rawData)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		req := &UserService.UploadImageRequest{
			Data: &UserService.UploadImageRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			return err
		}
	}

	if _, closeErr := stream.CloseAndRecv(); closeErr != nil {
		return err
	}

	return nil
}
