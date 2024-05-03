package dto

import (
	"github.com/kenboo0426/franky_assessment/application/output"
	"github.com/kenboo0426/franky_assessment/domain/model"
)

func NewBrand(r *model.Brand) *output.BrandDTO {
	return &output.BrandDTO{
		ID:   r.ID,
		Name: r.Name,
	}
}
