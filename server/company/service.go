package company

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Service struct {
	DB *gorm.DB
}

func NewCompanyService(db *gorm.DB) *Service {
	return &Service{DB: db}
}

func (s *Service) CreateCompany(ctx context.Context, company *Company) error {
	return s.DB.WithContext(ctx).Create(company).Error
}

func (s *Service) GetAllCompanies(ctx context.Context) ([]Company, error) {
	var companies []Company
	result := s.DB.WithContext(ctx).Preload("Products.Category").Find(&companies)
	return companies, result.Error
}

func (s *Service) GetCompanyById(ctx context.Context, id string) (*Company, error) {
	var company Company
	result := s.DB.WithContext(ctx).First(&company, id)
	return &company, result.Error
}

func (s *Service) UpdateCompany(ctx context.Context, id string, updatedData *Company) error {
	result := s.DB.WithContext(ctx).Model(&Company{}).Where("id = ?", id).Updates(updatedData)
	return result.Error
}

func (s *Service) DeleteCompany(ctx context.Context, id string) (*Company, error) {
	var company Company
	result := s.DB.WithContext(ctx).Clauses(clause.Returning{}).Where("id = ?", id).Delete(&company)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &company, nil
}
