package application

import (
	"context"

	"github.com/kenboo0426/franky_assessment/application/input"
	"github.com/kenboo0426/franky_assessment/application/output"
	"github.com/kenboo0426/franky_assessment/common"
	"github.com/kenboo0426/franky_assessment/domain/dto"
	"github.com/kenboo0426/franky_assessment/domain/model"
	"github.com/kenboo0426/franky_assessment/domain/repository"
	"github.com/kenboo0426/franky_assessment/domain/service"
)

type IProductUsecase interface {
	GetAllProduct(ctx context.Context) (*output.GetAllProductDTO, error)
	SearchProduct(ctx context.Context, input *input.SearchProductQueryDTO) (*output.SearchProductDTO, error)
	CreateProduct(ctx context.Context, input *input.CreateProductDTO) (*output.CreateProductDTO, error)
}

type productUsecase struct {
	repository.IBrandRepository
	repository.IProductRepository
	repository.IProductImageRepository
	service.IProductService
}

func NewProductUsecase(
	brandRepo repository.IBrandRepository,
	productRepo repository.IProductRepository,
	productImageRepo repository.IProductImageRepository,
	productService service.IProductService,
) IProductUsecase {
	return &productUsecase{
		brandRepo,
		productRepo,
		productImageRepo,
		productService,
	}
}

func (r productUsecase) GetAllProduct(ctx context.Context) (*output.GetAllProductDTO, error) {
	products, err := r.IProductService.FindAllWithDetail(ctx)
	if err != nil {
		return nil, err
	}

	return (*output.GetAllProductDTO)(common.Ptr(dto.NewProducts(products))), nil
}

func (r productUsecase) SearchProduct(ctx context.Context, input *input.SearchProductQueryDTO) (*output.SearchProductDTO, error) {
	products, err := r.IProductService.FindByBrandIDWithDetail(ctx, input.Brand)
	if err != nil {
		return nil, err
	}

	return (*output.SearchProductDTO)(common.Ptr(dto.NewProducts(products))), nil
}

func (r productUsecase) CreateProduct(ctx context.Context, input *input.CreateProductDTO) (*output.CreateProductDTO, error) {
	brand, _ := r.IBrandRepository.FindByName(ctx, input.Brand)
	if brand == nil {
		if newBrand, err := r.IBrandRepository.Create(ctx, &model.Brand{
			Name: input.Brand,
		}); err != nil {
			return nil, err
		} else {
			brand = newBrand
		}
	}

	product, err := r.IProductRepository.Create(ctx, &model.Product{
		BrandID: brand.ID,
		Name:    input.Name,
	})
	if err != nil {
		return nil, err
	}

	images := make([]model.ProductImage, len(input.Images))
	for i, image := range input.Images {
		images[i] = model.ProductImage{
			ProductID: product.ID,
			URL:       image,
		}
	}

	productImages, err := r.IProductImageRepository.CreateInBatch(ctx, images)
	if err != nil {
		return nil, err
	}

	product.Brand = brand
	product.ProductImages = productImages

	return (*output.CreateProductDTO)(dto.NewProduct(product)), nil
}
