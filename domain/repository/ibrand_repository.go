package repository

import (
	"context"

	"github.com/kenboo0426/franky_assessment/domain/model"
)

type IBrandRepository interface {
	FindByName(ctx context.Context, name string) (*model.Brand, error)
	Create(ctx context.Context, brand *model.Brand) (*model.Brand, error)
}
