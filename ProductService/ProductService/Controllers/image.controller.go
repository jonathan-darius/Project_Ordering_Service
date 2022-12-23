package Controllers

import (
	"ProductService/ProductService/Utils"
	pb "ProductService/proto"
	"bytes"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
)

const maxImageSize = 1 << 20

func StreamImage(stream pb.ProductManagementService_UploadProductImageServer) error {
	req, err := stream.Recv()
	ctx := context.Background()
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot receive image info")
	}
	productID := req.GetInfo().GetId()
	imageType := req.GetInfo().GetImageType()

	_, err = GetProduct(ctx, productID)
	if err != nil {
		return err
	}

	log.Printf("receive an upload-image request for user %s with image type %s", productID, imageType)
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

	res, err := Utils.SaveToStorage(productID, imageData)
	if err != nil {
		return status.Errorf(codes.Internal, "Cannot Save: %v", err)
	}
	if err := AddImage(ctx, productID, res); err != nil {
		return err
	}
	return nil
}
