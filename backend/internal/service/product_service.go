package service

import (
	"errors"
	"fmt"

	"product-api/internal/domain/product"
	"product-api/internal/repository"
	"product-api/pkg/validator"
)


type ErrValidation struct{ Message string }

func (e *ErrValidation) Error() string { return e.Message }


type ProductService interface {
	ListProducts() ([]product.Product, error)
	AddProduct(code string) (*product.Product, error)
	RemoveProduct(id uint) error
}

type productService struct {
	repo repository.ProductRepository
}


func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) ListProducts() ([]product.Product, error) {
	return s.repo.FindAll()
}

func (s *productService) AddProduct(code string) (*product.Product, error) {
	if err := validator.ProductCode(code); err != nil {
		return nil, &ErrValidation{Message: err.Error()}
	}

	p := &product.Product{ProductCode: code}
	if err := s.repo.Create(p); err != nil {
		if errors.Is(err, repository.ErrDuplicate) {
			return nil, &ErrValidation{Message: "รหัสสินค้านี้มีอยู่แล้วในระบบ"}
		}
		return nil, fmt.Errorf("AddProduct: %w", err)
	}
	return p, nil
}

func (s *productService) RemoveProduct(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return &ErrValidation{Message: "ไม่พบรหัสสินค้า"}
		}
		return fmt.Errorf("RemoveProduct: %w", err)
	}
	return nil
}
