package persist

import (
	"context"

	"github.com/krossroad/imperialfleet/internal/models"
)

type Persist interface {
	List(ctx context.Context, lr models.ListCraftRequest) ([]*models.SpaceCraft, error)
	Create(context.Context, *models.SpaceCraft) error
	Update(context.Context, *models.SpaceCraft) error
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, id int) (*models.SpaceCraft, error)
}
