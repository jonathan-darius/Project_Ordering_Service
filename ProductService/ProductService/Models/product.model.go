package Models

type ProductModel struct {
	Id        string   `bson:"_id" json:"id"`
	Name      string   `bson:"name" json:"name"`
	Stock     int32    `bson:"stock" json:"stock"`
	Price     int64    `bson:"price" json:"price"`
	Sold      int32    `bson:"sold" json:"sold"`
	Rating    float64  `bson:"rating" json:"rating"`
	Rated     int32    `bson:"rated" json:"rated"`
	Desc      string   `bson:"desc" json:"desc"`
	Category  []string `bson:"category" json:"category"`
	Image     []string `bson:"image" json:"image"`
	CreatedBy string   `bson:"createdBy" json:"createdBy"`
	CreatedAt int64    `bson:"createdAt" json:"createdAt"`
	UpdatedAt int64    `bson:"updatedAt" json:"updatedAt"`
	Deleted   bool     `bson:"deleted" json:"deleted"`
}

type ProductModelElastic struct {
	Id        string   `json:"id"`
	Type      string   `json:"type"`
	Name      string   `json:"name"`
	Stock     int32    `json:"stock"`
	Price     int64    `json:"price"`
	Sold      int32    `json:"sold"`
	Rating    float64  `json:"rating"`
	Rated     int32    `json:"rated"`
	Desc      string   `json:"desc"`
	Category  []string `json:"category"`
	Image     []string `json:"image"`
	CreatedBy string   `json:"createdBy"`
	CreatedAt int64    `json:"createdAt"`
	UpdatedAt int64    `json:"updatedAt"`
	Deleted   bool     `json:"deleted"`
}
