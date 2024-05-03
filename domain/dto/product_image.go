package dto

import (
	"github.com/kenboo0426/franky_assessment/application/output"
	"github.com/kenboo0426/franky_assessment/domain/model"
)

func NewProductImages(r []model.ProductImage) []output.ProductImageDTO {
	result := make([]output.ProductImageDTO, 0)

	for _, v := range r {
		result = append(result, output.ProductImageDTO{
			ID:  v.ID,
			URL: v.URL,
		})
	}

	return result
}
