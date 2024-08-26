package inventory

import (
	"context"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func NewInventoryService(db *gorm.DB) *Service {
	return &Service{DB: db}
}

func (s *Service) CreateInventory(ctx context.Context, inventory *Inventory) error {
	return s.DB.WithContext(ctx).Create(inventory).Error
}

func (s *Service) GetAllInventories(ctx context.Context) ([]Inventory, error) {
	var inventories []Inventory
	result := s.DB.WithContext(ctx).Find(&inventories)
	return inventories, result.Error
}
