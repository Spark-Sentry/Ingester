package repositories

import (
	"Ingester/internal/domain/entities"
	"context"
)

type MeasureRepository interface {
	Save(ctx context.Context, measure entities.Measure) error
}
