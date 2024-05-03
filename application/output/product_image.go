package output

type ProductImageDTO struct {
	ID      uint        `json:"id"`
	URL     string      `json:"url"`
	Product *ProductDTO `json:"product"`
}
