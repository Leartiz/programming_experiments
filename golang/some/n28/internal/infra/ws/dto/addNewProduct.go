package dto

type AddNewProduct struct {
	Name  string  `json:"name"`
	Desc  string  `json:"description"`
	Price float64 `json:"price"`
}
