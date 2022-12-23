package utils

type Pagination struct {
	Page   int32 `json:"page"`
	Record int32 `json:"record"`
	Data   any   `json:"data"`
}
