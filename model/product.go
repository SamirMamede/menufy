package model

type Product struct {
	ID          int     `json:"id_product"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImagePath   string  `json:"image_path"`
}
