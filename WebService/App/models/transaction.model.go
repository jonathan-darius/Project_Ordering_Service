package models

type ProductRating struct {
	UserID        string `json:"user_id" swaggerignore:"true"`
	TransactionID string `json:"transaction_id"`
	ProductID     string `json:"product_id"`
	Rating        int64  `json:"rating"`
	Desc          string `json:"desc"`
}

type TransactionDetail struct {
	ProductID string
	QTY       int32
	Rating    float64
	Desc      string
	Total     int64
}

type Transaction struct {
	TransactionID string
	UserID        string
	Total         int64
	CreatedAt     int64
	Detail        []TransactionDetail
}
