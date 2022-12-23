package main

import (
	"UserServices/UserService/Databases"
	"UserServices/UserService/Model"
	"UserServices/UserService/MsgQueue"
	"UserServices/UserService/utils"
	pb "UserServices/proto"
	"bytes"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
	"time"
)

const maxImageSize = 1 << 20

type Server struct {
	pb.UserServiceServer
}

//dbCollection
var (
	userCollection  = Databases.OpenCollection(Databases.Client, "user")
	verifCollection = Databases.OpenCollection(Databases.Client, "verif")
)

func (*Server) RegisterUser(ctx context.Context, in *pb.User) (*pb.UserId, error) {
	log.Println("User Register Invoke")

	count, err := userCollection.CountDocuments(ctx, bson.M{"email": in.Email})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	if count > 0 {
		return nil, status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf("Email Already Exists"),
		)
	}
	user := Model.UserModel{
		ID:          primitive.NewObjectID().Hex(),
		Email:       in.Email,
		Username:    in.Username,
		Image:       "",
		PhoneNumber: in.PhoneNumber,
		About:       in.About,
		Birthdate:   in.Birthdate,
		Password:    utils.HashPassword(in.Password),
		Role:        in.Role,
		CreatedAt:   time.Now().UnixNano(),
		Validate:    false,
	}
	//fmt.Println(user)
	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	return &pb.UserId{
		Id: user.ID,
	}, nil
}

func (*Server) GetUserByID(_ context.Context, in *pb.UserId) (*pb.User, error) {
	log.Println("Get UserByID Invoke")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user Model.UserModel
	res := userCollection.FindOne(ctx, bson.M{"_id": in.Id})
	defer cancel()
	if err := res.Decode(&user); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find user with specified ID: %v", err),
		)
	}
	return Model.DocumentToUser(&user), nil
}

func (*Server) GetUserByEmail(_ context.Context, in *pb.UserEmail) (*pb.User, error) {
	log.Println("Get UserByID Invoke")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user Model.UserModel
	res := userCollection.FindOne(ctx, bson.M{"email": in.Email})
	defer cancel()
	if err := res.Decode(&user); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find user with specified email: %v", err),
		)
	}
	return Model.DocumentToUser(&user), nil
}

func (*Server) ListUser(pagination *pb.Pagination, stream pb.UserService_ListUserServer) error {
	log.Println("List User was invoked")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	recordPerPage := pagination.RecordPerPage
	page := pagination.Page
	startIndex := (page - 1) * recordPerPage
	matchStage := bson.D{{"$match", bson.D{{}}}}
	groupStage := bson.D{{"$group", bson.D{
		{"_id", bson.D{{"_id", "null"}}},
		{"total_count", bson.D{{"$sum", 1}}},
		{"data", bson.D{{"$push", "$$ROOT"}}}}}}
	projectStage := bson.D{
		{"$project", bson.D{
			{"_id", 0},
			{"total_count", 1},
			{"user_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}}}}}
	result, err := userCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage, groupStage, projectStage})
	defer cancel()
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	var allusers []bson.M
	if err = result.All(ctx, &allusers); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(allusers[0]["user_items"].(primitive.A)[0].(primitive.M))
	var user Model.UserModel
	for _, x := range allusers[0]["user_items"].(primitive.A) {
		bsonBytes, _ := bson.Marshal(x)
		if err := bson.Unmarshal(bsonBytes, &user); err != nil {
			return err
		}
		//fmt.Println(i, user)
		err := stream.Send(Model.DocumentToUser(&user))
		if err != nil {
			return err
		}
	}

	if err = result.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

	return nil
}

func (*Server) DeleteUser(ctx context.Context, in *pb.UserId) (*empty.Empty, error) {
	log.Printf("Delete User was invoked with %v\n", in)
	res, err := userCollection.DeleteOne(ctx, bson.M{"_id": in.Id})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot delete object in MongoDB: %v", err),
		)
	}
	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"User was not found",
		)
	}
	return &emptypb.Empty{}, nil
}

func (*Server) GetVerification(_ context.Context, in *pb.UserEmail) (*empty.Empty, error) {
	log.Println("User Verification Invoke")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	res := userCollection.FindOne(ctx, bson.M{"email": in.Email})
	var user Model.UserModel
	defer cancel()
	if err := res.Decode(&user); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot validate user with specified ID: %v", err),
		)
	}
	if user.Validate {
		return nil, status.Errorf(
			codes.Aborted,
			fmt.Sprintf("User Already Validated"),
		)
	}
	var tokenData Model.Verification
	check := verifCollection.FindOne(ctx, bson.M{"email": user.Email})
	if err := check.Decode(&tokenData); err == nil {
		_, err := verifCollection.DeleteOne(ctx, bson.M{"_id": tokenData.UserID})
		if err != nil {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Cannot delete object in MongoDB: %v", err),
			)
		}
	}
	now := time.Now()
	newToken, _ := utils.RandToken(5)
	VerifDetail := Model.Verification{
		UserID:    user.ID,
		Email:     user.Email,
		Token:     newToken,
		CreatedAt: now.UnixNano(),
		ExpiredAt: now.Add(time.Minute * time.Duration(5)).UnixNano(),
	}
	_, err := verifCollection.InsertOne(ctx, VerifDetail)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	MsgQueue.SendMSG(VerifDetail)

	return &emptypb.Empty{}, nil
}

func (*Server) VerificationUser(_ context.Context, in *pb.ValidateRequest) (*empty.Empty, error) {
	log.Println("User Verification Invoke")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	res := userCollection.FindOne(ctx, bson.M{"email": in.Email})
	var user Model.UserModel
	defer cancel()
	if err := res.Decode(&user); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot validate user with specified Email: %v", err),
		)
	}
	if user.Validate {
		return nil, status.Errorf(
			codes.Aborted,
			fmt.Sprintf("User Already Validated"),
		)
	}
	var tokenData Model.Verification
	check := verifCollection.FindOne(ctx, bson.M{"email": user.Email})
	if err := check.Decode(&tokenData); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("User Token Invalid"),
		)
	}

	if tokenData.Token == in.Token {
		if tokenData.ExpiredAt < time.Now().UnixNano() {
			return nil, status.Errorf(
				codes.DeadlineExceeded,
				fmt.Sprintf("Token Expired"),
			)
		}
		filter := bson.D{bson.E{
			Key:   "_id",
			Value: user.ID,
		}}
		updateData := bson.D{bson.E{Key: "$set", Value: bson.D{
			bson.E{Key: "validate", Value: true},
		}}}
		if _, err := userCollection.UpdateOne(ctx, filter, updateData); err != nil {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Failed Validate"),
			)
		}
		if _, err := verifCollection.DeleteOne(ctx, bson.M{"_id": user.ID}); err != nil {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Cannot delete object in MongoDB: %v", err),
			)
		}
	} else {
		return nil, status.Errorf(
			codes.Canceled,
			fmt.Sprintf("User Token Invalid"),
		)
	}
	return &emptypb.Empty{}, nil
}

func (*Server) Login(_ context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Println("Login Invoke")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user Model.UserModel
	res := userCollection.FindOne(ctx, bson.M{"email": in.Email})
	defer cancel()

	if err := res.Decode(&user); err != nil {
		return nil, status.Errorf(
			codes.PermissionDenied,
			fmt.Sprintf("Wrong Email / Password"),
		)
	}
	if check, _ := utils.VerifyPassword(in.Password, user.Password); !check {
		return nil, status.Errorf(
			codes.PermissionDenied,
			fmt.Sprintf("Wrong Email or Password"),
		)
	}
	token, refresh, err := utils.GenerateToken(
		user.Email,
		user.ID,
		user.Role,
	)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Failed Generate"),
		)
	}

	return &pb.LoginResponse{
		Token:   token,
		Refresh: refresh,
	}, nil
}

func (*Server) ValidateToken(_ context.Context, in *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	log.Println("Validate Token Invoke")
	claims, err := utils.ValidateToken(in.Token)
	if err != "" {
		return &pb.ValidateTokenResponse{
			Status:   false,
			Messages: err,
		}, nil
	}
	return &pb.ValidateTokenResponse{
		Status:   true,
		Messages: "Success",
		UserId:   claims.Uid,
		Role:     claims.Role,
	}, nil
}

func (*Server) UpdateUser(_ context.Context, in *pb.User) (*empty.Empty, error) {
	log.Println("User Update Invoke")
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var user Model.UserModel
	res := userCollection.FindOne(ctx, bson.M{"_id": in.Id})
	defer cancel()
	if err := res.Decode(&user); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("User Not Found"),
		)
	}
	filter := bson.D{bson.E{
		Key:   "_id",
		Value: user.ID,
	}}
	updateVar := bson.D{}
	if in.Username != "" {
		updateVar = append(updateVar, bson.E{Key: "username", Value: in.Username})
	}
	if in.PhoneNumber != "" {
		updateVar = append(updateVar, bson.E{Key: "phoneNumber", Value: in.PhoneNumber})
	}
	if in.About != "" {
		updateVar = append(updateVar, bson.E{Key: "about", Value: in.About})
	}
	if in.Birthdate != 0 {
		updateVar = append(updateVar, bson.E{Key: "birthdate", Value: in.Birthdate})
	}

	updateData := bson.D{bson.E{Key: "$set", Value: updateVar}}

	if _, err := userCollection.UpdateOne(ctx, filter, updateData); err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Failed Validate"),
		)
	}

	return &emptypb.Empty{}, nil
}

func (*Server) UploadImage(stream pb.UserService_UploadImageServer) error {
	req, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot receive image info")
	}
	userID := req.GetInfo().UserId
	imageType := req.GetInfo().GetImageType()
	log.Printf("receive an upload-image request for user %s with image type %s", userID, imageType)
	imageData := bytes.Buffer{}
	imageSize := 0
	for {
		log.Print("waiting to receive more data")

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err)
		}
		chunk := req.GetChunkData()
		size := len(chunk)

		log.Printf("received a chunk with size: %d", size)

		imageSize += size
		if imageSize > maxImageSize {
			return status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", imageSize, maxImageSize)
		}
		_, err = imageData.Write(chunk)
		if err != nil {
			return status.Errorf(codes.Internal, "cannot write chunk data: %v", err)
		}
	}
	res, err := utils.SaveToStorage(userID, imageData)
	if err != nil {
		return status.Errorf(codes.Internal, "Cannot Save: %v", err)
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	filter := bson.D{bson.E{
		Key:   "_id",
		Value: userID,
	}}

	updateData := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "image", Value: res.Path}}}}
	defer cancel()
	if _, err := userCollection.UpdateOne(ctx, filter, updateData); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Failed Validate"),
		)
	}
	return nil
}
