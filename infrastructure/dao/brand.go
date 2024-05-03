package dao

import (
	"context"

	"github.com/kenboo0426/franky_assessment/domain/model"
	"github.com/kenboo0426/franky_assessment/domain/repository"
	"github.com/kenboo0426/franky_assessment/infrastructure/external/mysql"
)

type brandRepository struct {
	db mysql.MysqlClient
}

func NewBrandRepository(db mysql.MysqlClient) repository.IBrandRepository {
	return &brandRepository{
		db: db,
	}
}

func (r brandRepository) FindByName(ctx context.Context, name string) (*model.Brand, error) {
	return nil, nil
}

func (r brandRepository) Create(ctx context.Context, brand *model.Brand) (*model.Brand, error) {
	return nil, nil
}
