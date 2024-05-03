package dao

import (
	"context"

	"github.com/kenboo0426/franky_assessment/domain/model"
	"github.com/kenboo0426/franky_assessment/domain/repository"
	"github.com/kenboo0426/franky_assessment/infrastructure/external/mysql"
)

type productRepository struct {
	db mysql.MysqlClient
}

func NewProductRepository(db mysql.MysqlClient) repository.IProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r productRepository) Create(ctx context.Context, product *model.Product) (*model.Product, error) {
	return nil, nil
}
