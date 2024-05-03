package model

type ProductImage struct {
	BaseModel
	ProductID uint   `bun:"product_id,notnull"`
	URL       string `bun:"url,notnull"`

	Product *Product `bun:"rel:belongs-to"`
}
