package repository

import (
	"context"

	"github.com/kenboo0426/franky_assessment/domain/model"
)

type IProductImageRepository interface {
	Create(ctx context.Context, productImage *model.ProductImage) (*model.ProductImage, error)
	CreateInBatch(ctx context.Context, productImages []model.ProductImage) ([]model.ProductImage, error)
}
