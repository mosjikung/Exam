package product

import "time"

// Product is the core domain entity.
// No Fiber, no GORM tags that leak business rules — only what the domain needs.
type Product struct {
	ID          uint      `json:"id"           gorm:"primaryKey;autoIncrement"`
	ProductCode string    `json:"product_code" gorm:"type:varchar(19);uniqueIndex;not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TableName overrides GORM default ("products" would be inferred anyway,
// but explicit is better than implicit).
func (Product) TableName() string { return "products" }
