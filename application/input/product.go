package input

type SearchProductURIDTO struct {
	Brand string `uri:"brand" binding:"required"`
}

type CreateProductDTO struct {
	Name   string   `form:"name" json:"name" binding:"required"`
	Brand  string   `form:"brand" json:"brand" binding:"required"`
	Images []string `form:"images" json:"images" binding:"dive,required"`
}
