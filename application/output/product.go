package output

type ProductDTO struct {
	ID     uint              `json:"id"`
	Name   string            `json:"name"`
	Brand  *BrandDTO         `json:"brand"`
	Images []ProductImageDTO `json:"images"`
}

type GetAllProductDTO []ProductDTO

type SearchProductDTO []ProductDTO

type CreateProductDTO ProductDTO
