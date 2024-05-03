package service

import (
	"context"

	"github.com/kenboo0426/franky_assessment/domain/model"
	"github.com/kenboo0426/franky_assessment/domain/service"
	"github.com/kenboo0426/franky_assessment/infrastructure/external/mysql"
)

type productService struct {
	db mysql.MysqlClient
}

func NewProductService(db mysql.MysqlClient) service.IProductService {
	return &productService{
		db: db,
	}
}

func (r productService) FindAllWithDetail(ctx context.Context) ([]model.Product, error) {
	return nil, nil
}

func (r productService) FindByBrandIDWithDetail(ctx context.Context, brandID string) ([]model.Product, error) {
	return nil, nil
}
