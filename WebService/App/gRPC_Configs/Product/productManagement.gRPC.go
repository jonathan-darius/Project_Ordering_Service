package gRPCFunc

import (
	"WebService/App/gRPC_Configs/proto/ProductService"
	"WebService/App/models"
	"bufio"
	"context"
	"io"
	"mime/multipart"
	"path/filepath"
	"time"
)

func AddProduct(data models.Product) (*ProductService.ProductId, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	res, err := Management.RegisterProduct(ctx, &ProductService.Product{
		Id:        data.ID,
		Name:      data.Name,
		Price:     data.Price,
		Stock:     data.Stock,
		Desc:      data.Desc,
		Category:  data.Category,
		CreatedBy: data.CreatedBy,
	})
	defer cancel()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func ProductUpdate(data *models.UpdateProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	dataUpdate := &ProductService.UpdateProductMSG{
		Id: data.ID,
		Name: &ProductService.StrUpdate{
			Check: false,
			Value: "",
		},
		Price: &ProductService.NumUpdate{
			Check: false,
			Value: 0,
		},
		Stock: &ProductService.NumUpdate{
			Check: false,
			Value: 0,
		},
		Desc: &ProductService.StrUpdate{
			Check: false,
			Value: "",
		},
		Category: &ProductService.ArrUpdate{
			Check: false,
			Value: []string{},
		},
		Image: &ProductService.ArrUpdate{
			Check: false,
			Value: []string{},
		},
	}

	if data.Name != nil {
		dataUpdate.Name.Check = true
		dataUpdate.Name.Value = *data.Name
	}

	if data.Price != nil {
		dataUpdate.Price.Check = true
		dataUpdate.Price.Value = *data.Price
	}

	if data.Stock != nil {
		dataUpdate.Stock.Check = true
		dataUpdate.Stock.Value = int64(*data.Stock)
	}

	if data.Desc != nil {
		dataUpdate.Desc.Check = true
		dataUpdate.Desc.Value = *data.Desc
	}

	if data.Category != nil {
		dataUpdate.Category.Check = true
		dataUpdate.Category.Value = *data.Category
	}

	if data.Image != nil {
		dataUpdate.Image.Check = true
		dataUpdate.Image.Value = *data.Image
	}

	_, err := Management.UpdateProduct(ctx, dataUpdate)
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}

func DelProduct(pID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_, err := Management.DeleteProduct(ctx, &ProductService.ProductId{Id: pID})
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}

func AddCategory(data *models.ProductArr) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_, err := Management.AddCategory(ctx, &ProductService.UpdateProductMSG{Id: data.ID, Category: &ProductService.ArrUpdate{
		Check: true,
		Value: data.Category,
	}})
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}

func RemoveCategory(data *models.ProductArr) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_, err := Management.RemoveCatogory(ctx, &ProductService.UpdateProductMSG{Id: data.ID, Category: &ProductService.ArrUpdate{
		Check: true,
		Value: data.Category,
	}})
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}

func AddStock(data *models.ProductStock) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_, err := Management.AddStock(ctx, &ProductService.Stock{
		Id:    data.ID,
		Stock: int32(data.Stock),
	})
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}

func ProductImage(file *multipart.FileHeader, uid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stream, err := Management.UploadProductImage(ctx)
	if err != nil {
		return err
	}

	req := &ProductService.UploadImageRequest{
		Data: &ProductService.UploadImageRequest_Info{
			Info: &ProductService.ImageInfo{
				Id:        uid,
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

		req := &ProductService.UploadImageRequest{
			Data: &ProductService.UploadImageRequest_ChunkData{
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
