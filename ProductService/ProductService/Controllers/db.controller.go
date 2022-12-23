package Controllers

import (
	"ProductService/ProductService/Config"
	"ProductService/ProductService/Databases"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ProductCollection = Databases.OpenCollection("product")
	EClient           = Databases.ElastiClient
	buf               *bytes.Buffer
)

func ElasticUpdate(pID string, body any, ctx context.Context) error {
	buf = &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error Encode"),
		)
	}
	req := esapi.UpdateRequest{
		Index:      Config.GetEnv("ELASTICSEARCH_INDEX"),
		DocumentID: pID,
		Body:       bytes.NewReader(buf.Bytes()),
		Refresh:    "true",
	}
	res, err := req.Do(ctx, EClient)
	defer res.Body.Close()
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Failed"),
		)
	}
	return nil
}
