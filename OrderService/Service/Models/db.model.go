package Models

type UserCart struct {
	UserID    string
	ProductID string
	QTY       int32
	Total     int64
	CreatedAt int64
}

type Purchase struct {
	TransactionID string `gorm:"primaryKey"`
	UserID        string
	Total         int64
	CreatedAt     int64
}

type PurchaseDetail struct {
	TransactionID string `gorm:"primaryKey;autoIncrement:false"`
	ProductID     string `gorm:"primaryKey;autoIncrement:false"`
	QTY           int32
	Total         int64
	Rating        float64
	Desc          string
}
