package sql

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/krossroad/imperialfleet/internal/models"
	"github.com/krossroad/imperialfleet/internal/persist"
)

var _ persist.Persist = (*SQL)(nil)

type SQL struct {
	db *gorm.DB
}

func (s *SQL) List(ctx context.Context, lr models.ListCraftRequest) ([]*models.SpaceCraft, error) {
	craft := []*models.SpaceCraft{}

	tx := s.db.WithContext(ctx).Order("id DESC").Preload("Armaments")

	if lr.Name != "" {
		tx.Where("name like ?", lr.Name+"%")
	} else if lr.Class != "" {
		tx.Where("class = ?", lr.Class)
	} else if lr.Status != "" {
		tx.Where("status = ?", lr.Status)
	}

	if err := tx.Find(&craft).Error; err != nil {
		return nil, fmt.Errorf("failed to list items: %w", err)
	}

	return craft, nil
}

func (s *SQL) Create(ctx context.Context, mi *models.SpaceCraft) error {
	tx := s.db.WithContext(ctx).Create(mi)

	if err := tx.Error; err != nil {
		return fmt.Errorf("failed to create item: %w", err)
	} else if tx.RowsAffected == 0 {
		return fmt.Errorf("craft already exists")
	}

	return nil
}

func (s *SQL) Update(ctx context.Context, mi *models.SpaceCraft) error {
	old := &models.SpaceCraft{ID: mi.ID}
	tx := s.db.WithContext(ctx).Find(&old)
	if tx.Error != nil {
		return fmt.Errorf("failed to find item: %w", tx.Error)
	} else if tx.RowsAffected == 0 {
		return fmt.Errorf("craft not found")
	}
	mi.UpdatedAt = time.Now()
	mi.CreatedAt = old.CreatedAt

	tx = s.db.WithContext(ctx).Save(mi)

	if err := tx.Error; err != nil {
		return fmt.Errorf("failed to update item: %w", err)
	} else if tx.RowsAffected == 0 {
		return fmt.Errorf("craft not found")
	}

	return nil
}

func New(db *gorm.DB) *SQL {
	return &SQL{
		db: db,
	}
}

func (s *SQL) Delete(ctx context.Context, id int) error {
	tx := s.db.WithContext(ctx).Delete(&models.SpaceCraft{}, id)
	if err := tx.Error; err != nil {
		return fmt.Errorf("failed to delete item: %w", err)
	} else if tx.RowsAffected == 0 {
		return fmt.Errorf("craft not found")
	}

	return nil
}

func (s *SQL) Get(ctx context.Context, id int) (*models.SpaceCraft, error) {
	craft := &models.SpaceCraft{}
	tx := s.db.WithContext(ctx).Preload("Armaments").Where("id =?", id).First(craft)
	if err := tx.Error; err != nil {
		return nil, fmt.Errorf("failed to get item: %w", err)
	} else if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("craft not found")
	}

	return craft, nil
}
