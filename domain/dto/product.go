package dto

import (
	"github.com/kenboo0426/franky_assessment/application/output"
	"github.com/kenboo0426/franky_assessment/domain/model"
)

func NewProducts(r []model.Product) []output.ProductDTO {
	result := make([]output.ProductDTO, 0)

	for _, v := range r {
		result = append(result, *NewProduct(&v))
	}

	return result
}

func NewProduct(r *model.Product) *output.ProductDTO {
	result := &output.ProductDTO{
		ID:   r.ID,
		Name: r.Name,
	}

	if r.Brand != nil {
		result.Brand = NewBrand(r.Brand)
	}

	if len(r.ProductImages) > 0 {
		result.Images = NewProductImages(r.ProductImages)
	}

	return result
}
