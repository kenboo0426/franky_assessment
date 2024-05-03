package service

import (
	"context"

	"github.com/kenboo0426/franky_assessment/domain/model"
)

type IProductService interface {
	FindAllWithDetail(ctx context.Context) ([]model.Product, error)
	FindByBrandIDWithDetail(ctx context.Context, brandID string) ([]model.Product, error)
}
