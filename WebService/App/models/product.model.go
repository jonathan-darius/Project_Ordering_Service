package models

type Product struct {
	ID        string   `json:"id" swaggerignore:"true"`
	Name      string   `json:"name" example:"Buku"`
	Price     int64    `json:"price" example:"2000"`
	Stock     int32    `json:"stock" example:"5000"`
	Sold      int32    `json:"sold" swaggerignore:"true"`
	Rating    float64  `json:"rating" swaggerignore:"true"`
	Rated     int32    `json:"rated" swaggerignore:"true"`
	Desc      string   `json:"desc" example:"Ini Buku Ajaib"`
	Category  []string `json:"category" example:"buku,alat tulis,keren,sekolah"`
	Image     []string `json:"image" swaggerignore:"true"`
	CreatedBy string   `json:"createdBy" swaggerignore:"true"`
	CreatedAt int64    `json:"createdAt" swaggerignore:"true"`
	UpdatedAt int64    `json:"updatedAt" swaggerignore:"true"`
	Deleted   bool     `json:"deleted" swaggerignore:"true"`
}

type UpdateProduct struct {
	ID       string    `json:"id" example:"9232131"`
	Name     *string   `json:"name" example:"Buku"`
	Price    *int64    `json:"price" example:"2000"`
	Stock    *int32    `json:"stock" example:"5000"`
	Desc     *string   `json:"desc" example:"Ini Buku Ajaib"`
	Category *[]string `json:"category" example:"buku,alat tulis,keren,sekolah"`
	Image    *[]string `json:"image"`
}

type ProductArr struct {
	ID       string   `json:"id"`
	Category []string `json:"category" example:"buku,alat tulis,keren,sekolah"`
}

type ProductStock struct {
	ID    string `json:"id"`
	Stock int64  `json:"stock"`
}

type SearchProduct struct {
	Keyword   string  `json:"keyword" validate:"min=3"`
	Category  string  `json:"category"`
	PriceLow  int64   `json:"priceLow"`
	PriceHigh int64   `json:"priceHigh"`
	Rating    float64 `json:"rating"`
	SortBy    string  `json:"sortBy" validate:"eq=price|eq=rating|eq=name"`
	Order     string  `json:"order" validate:"eq=asc|eq=desc"`
}

type ProductShow struct {
	ID        string   `json:"id" swaggerignore:"true"`
	Name      string   `json:"name" example:"Buku"`
	Price     int64    `json:"price" example:"2000"`
	Stock     int32    `json:"stock" example:"5000"`
	Sold      int32    `json:"sold" swaggerignore:"true"`
	Rating    float64  `json:"rating" swaggerignore:"true"`
	Rated     int32    `json:"rated" swaggerignore:"true"`
	Desc      string   `json:"desc" example:"Ini Buku Ajaib"`
	Category  []string `json:"category" example:"buku,alat tulis,keren,sekolah"`
	CreatedAt int64    `json:"createdAt" swaggerignore:"true"`
}
