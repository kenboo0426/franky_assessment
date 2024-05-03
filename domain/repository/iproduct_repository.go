package repository

import (
	"context"

	"github.com/kenboo0426/franky_assessment/domain/model"
)

type IProductRepository interface {
	Create(ctx context.Context, product *model.Product) (*model.Product, error)
}
