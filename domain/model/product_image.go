package model

import "time"

type ProductImage struct {
	ID        uint      `bun:"id,pk,autoincrement"`
	ProductID uint      `bun:"product_id,notnull"`
	URL       string    `bun:"url,notnull"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`

	Product *Product `rel:"belongs-to"`
}
