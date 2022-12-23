package Controllers

import (
	"ProductService/ProductService/Config"
	pb "ProductService/proto"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func Search(record int32, page int32, param *pb.SearchQuery) (map[string]interface{}, error) {
	startIndex := (page - 1) * record
	var (
		varMust   []map[string]interface{}
		varShould []map[string]interface{}
		buf       bytes.Buffer
		r         map[string]interface{}
	)

	if param.Keyword != "" {
		varShould = append(varShould, map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":     param.Keyword,
				"fuzziness": "5",
				"fields":    []string{"name", "desc"},
			},
		})
		varShould = append(varShould, map[string]interface{}{
			"match": map[string]interface{}{
				"name": param.Keyword,
			},
		})
		varShould = append(varShould, map[string]interface{}{
			"match": map[string]interface{}{
				"title": param.Keyword,
			},
		})
	}

	if param.Category != "" {
		varMust = append(varMust, map[string]interface{}{
			"match": map[string]interface{}{
				"category": param.Category,
			},
		})
	}

	if param.PriceLow != 0 || param.PriceHigh != 0 {
		tmp := map[string]interface{}{
			"gte": param.PriceLow,
		}
		if param.PriceHigh != 0 {
			tmp["lte"] = param.PriceHigh
		}
		varMust = append(varMust, map[string]interface{}{
			"range": map[string]interface{}{
				"price": tmp,
			},
		})
	}

	if param.Rating != 0 {
		varMust = append(varMust, map[string]interface{}{
			"match": map[string]interface{}{
				"rating": param.Rating,
			},
		})
	}
	sField := "_score"
	sOrder := "asc"

	if param.SortBy != "" {
		sField = param.SortBy
	}
	if param.Order != "" {
		sOrder = param.Order
	}

	sort := map[string]interface{}{
		sField: sOrder,
	}
	query := map[string]interface{}{
		"size": record,
		"from": startIndex,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must":   varMust,
				"should": varShould,
			},
		},
		"sort": []any{sort, "_score"},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	searchService, err := EClient.Search(
		EClient.Search.WithContext(context.Background()),
		EClient.Search.WithIndex(Config.GetEnv("ELASTICSEARCH_INDEX")),
		EClient.Search.WithBody(&buf),
		EClient.Search.WithTrackTotalHits(true),
		EClient.Search.WithPretty(),
	)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("API Fail"),
		)
	}

	if searchService.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(searchService.Body).Decode(&e); err != nil {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Fail Decode"),
			)
		} else {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Fail"),
			)
		}
	}
	ans := searchService.Body

	if err := json.NewDecoder(ans).Decode(&r); err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Fail Parsing"),
		)
	}
	if err = ans.Close(); err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Fail Close"),
		)
	}

	return r, nil
}
