package dao

import (
	"context"

	"github.com/kenboo0426/franky_assessment/domain/model"
	"github.com/kenboo0426/franky_assessment/domain/repository"
	"github.com/uptrace/bun"
)

type productImageRepository struct {
	db *bun.DB
}

func NewProductImageRepository(db *bun.DB) repository.IProductImageRepository {
	return &productImageRepository{
		db: db,
	}
}

func (r productImageRepository) Create(ctx context.Context, productImage *model.ProductImage) (*model.ProductImage, error) {
	_, err := r.db.
		NewInsert().
		Model(productImage).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return productImage, nil
}

func (r productImageRepository) CreateInBatch(ctx context.Context, productImages []model.ProductImage) ([]model.ProductImage, error) {
	_, err := r.db.
		NewInsert().
		Model(&productImages).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return productImages, nil
}
