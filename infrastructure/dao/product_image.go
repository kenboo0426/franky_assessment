package dao

import (
	"context"

	"github.com/kenboo0426/franky_assessment/domain/model"
	"github.com/kenboo0426/franky_assessment/domain/repository"
	"github.com/kenboo0426/franky_assessment/infrastructure/external/mysql"
)

type productImageRepository struct {
	db mysql.MysqlClient
}

func NewProductImageRepository(db mysql.MysqlClient) repository.IProductImageRepository {
	return &productImageRepository{
		db: db,
	}
}

func (r productImageRepository) Create(ctx context.Context, productImage *model.ProductImage) (*model.ProductImage, error) {
	return nil, nil
}

func (r productImageRepository) CreateInBatch(ctx context.Context, productImages []model.ProductImage) ([]model.ProductImage, error) {
	return nil, nil
}
