package product

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Service struct {
	DB *gorm.DB
}

func NewProductService(db *gorm.DB) *Service {
	return &Service{DB: db}
}

func (s *Service) CreateProduct(ctx context.Context, product *Product) error {
	return s.DB.WithContext(ctx).Create(product).Error
}

func (s *Service) GetAllProducts(ctx context.Context) ([]Product, error) {
	var products []Product
	result := s.DB.WithContext(ctx).Find(&products)
	return products, result.Error
}

func (s *Service) GetProductById(ctx context.Context, id string) (*Product, error) {
	var product Product
	result := s.DB.WithContext(ctx).First(&product, id)
	return &product, result.Error
}

func (s *Service) UpdateProduct(ctx context.Context, id string, updatedData *Product) error {
	result := s.DB.WithContext(ctx).Model(&Product{}).Where("id = ?", id).Updates(updatedData)
	return result.Error
}

func (s *Service) DeleteProduct(ctx context.Context, id string) (*Product, error) {
	var product Product
	result := s.DB.WithContext(ctx).Clauses(clause.Returning{}).Where("id = ?", id).Delete(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &product, nil
}
