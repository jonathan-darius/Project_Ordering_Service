package models

type CartItemSend struct {
	UserID    string `json:"userID" swaggerignore:"true"`
	ProductID string `json:"productID"`
	QTY       int32  `json:"qty"`
	Total     int64  `json:"total" swaggerignore:"true"`
}

type CartItem struct {
	ProductID string `json:"productID"`
	QTY       int32  `json:"qty"`
	Total     int64  `json:"total"`
}

type Cart struct {
	UserID string     `json:"userID"`
	Item   []CartItem `json:"item"`
	Total  int64      `json:"total"`
}
