package model

import "time"

type Brand struct {
	ID        uint     `bun:"id,pk,autoincrement"`
	Name      string    `bun:"name,notnull,unique"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`

	Products []Product `bun:"rel:has-many,join:id=brand_id"`
}
