package model

type Product struct {
	BaseModel
	BrandID uint   `bun:"brand_id,notnull"`
	Name    string `bun:"name,notnull"`

	Brand         *Brand         `bun:"rel:belongs-to"`
	ProductImages []ProductImage `bun:"rel:has-many,join:id=product_id"`
}
