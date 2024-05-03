package service

import (
	"context"

	"github.com/kenboo0426/franky_assessment/domain/model"
	"github.com/kenboo0426/franky_assessment/domain/service"
	"github.com/uptrace/bun"
)

type productService struct {
	db *bun.DB
}

func NewProductService(db *bun.DB) service.IProductService {
	return &productService{
		db: db,
	}
}

func (r productService) FindAllWithDetail(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	if err := r.db.
		NewSelect().
		Model(&products).
		Relation("Brand").
		Relation("ProductImages").
		Scan(ctx); err != nil {
		return nil, err
	}
	return products, nil
}

func (r productService) FindByBrandIDWithDetail(ctx context.Context, brandName string) ([]model.Product, error) {
	var products []model.Product
	if err := r.db.
		NewSelect().
		Model(&products).
		Relation("Brand", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("brand.name = ?", brandName)
		}).
		Relation("ProductImages").
		Scan(ctx); err != nil {
		return nil, err
	}
	return products, nil
}
