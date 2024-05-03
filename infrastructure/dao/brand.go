package dao

import (
	"context"

	"github.com/kenboo0426/franky_assessment/domain/model"
	"github.com/kenboo0426/franky_assessment/domain/repository"
	"github.com/uptrace/bun"
)

type brandRepository struct {
	db *bun.DB
}

func NewBrandRepository(db *bun.DB) repository.IBrandRepository {
	return &brandRepository{
		db: db,
	}
}

func (r brandRepository) FindByName(ctx context.Context, name string) (*model.Brand, error) {
	var brand model.Brand
	err := r.db.
		NewSelect().
		Model(&brand).
		Where("name = ?", name).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &brand, nil
}

func (r brandRepository) Create(ctx context.Context, brand *model.Brand) (*model.Brand, error) {
	_, err := r.db.
		NewInsert().
		Model(brand).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	return brand, nil
}
