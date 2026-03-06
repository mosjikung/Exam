package repository

import (
	"errors"
	"fmt"
	"strings"

	"product-api/internal/domain/product"

	"gorm.io/gorm"
)

var ErrNotFound  = errors.New("record not found")
var ErrDuplicate = errors.New("duplicate product code")

type ProductRepository interface {
	FindAll() ([]product.Product, error)
	FindByID(id uint) (*product.Product, error)
	Create(p *product.Product) error
	Delete(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) FindAll() ([]product.Product, error) {
	var products []product.Product
	if err := r.db.Order("id asc").Find(&products).Error; err != nil {
		return nil, fmt.Errorf("FindAll: %w", err)
	}
	return products, nil
}

func (r *productRepository) FindByID(id uint) (*product.Product, error) {
	var p product.Product
	err := r.db.First(&p, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("FindByID: %w", err)
	}
	return &p, nil
}

func (r *productRepository) Create(p *product.Product) error {
	if err := r.db.Create(p).Error; err != nil {
		if strings.Contains(err.Error(), "23505") || strings.Contains(err.Error(), "unique") {
			return ErrDuplicate
		}
		return fmt.Errorf("Create: %w", err)
	}
	return nil
}

func (r *productRepository) Delete(id uint) error {
	result := r.db.Delete(&product.Product{}, id)
	if result.Error != nil {
		return fmt.Errorf("Delete: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}
