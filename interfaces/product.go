package interfaces

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenboo0426/franky_assessment/application"
	"github.com/kenboo0426/franky_assessment/interfaces/request"
)

type IProductHandler interface {
	GetAll(c *gin.Context)
	Search(c *gin.Context)
	Create(c *gin.Context)
}

type productHandler struct {
	application.IProductUsecase
}

func NewProductHandler(productUsecase application.IProductUsecase) IProductHandler {
	return &productHandler{
		IProductUsecase: productUsecase,
	}
}

func (r productHandler) GetAll(c *gin.Context) {
	products, err := r.IProductUsecase.GetAllProduct(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (r productHandler) Search(c *gin.Context) {
	request := request.SearchProductDTO{}
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	products, err := r.IProductUsecase.SearchProduct(c, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (r productHandler) Create(c *gin.Context) {
	request := request.PostProductDTO{}
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	product, err := r.IProductUsecase.CreateProduct(c, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}
