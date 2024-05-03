package dao

import (
	"context"

	"github.com/kenboo0426/franky_assessment/domain/model"
	"github.com/kenboo0426/franky_assessment/domain/repository"
	"github.com/uptrace/bun"
)

type productRepository struct {
	db *bun.DB
}

func NewProductRepository(db *bun.DB) repository.IProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r productRepository) Create(ctx context.Context, product *model.Product) (*model.Product, error) {
	_, err := r.db.
		NewInsert().
		Model(product).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return product, nil
}
