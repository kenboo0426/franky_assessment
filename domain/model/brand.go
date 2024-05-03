package model

type Brand struct {
	BaseModel
	Name string `bun:"name,notnull,unique"`

	Products []Product `bun:"rel:has-many,join:id=brand_id"`
}
