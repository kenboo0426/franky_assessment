package model

import "time"

type Product struct {
	ID        uint      `bun:"id,pk,autoincrement"`
	BrandID   uint      `bun:"brand_id,notnull"`
	Name      string    `bun:"name,notnull"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`

	Brand         *Brand         `bun:"rel:belongs-to"`
	ProductImages []ProductImage `bun:"rel:has-many,join:id=product_id"`
}
