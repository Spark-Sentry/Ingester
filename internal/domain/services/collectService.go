package services

import (
	"Ingester/internal/application/dto"
	"Ingester/internal/domain/entities"
	"Ingester/internal/domain/repositories"
	"context"
	"fmt"
	"log"
	"time"
)

type CollectService struct {
	measureRepo repositories.MeasureRepository
}

func NewCollectService(measureRepo repositories.MeasureRepository) *CollectService {
	return &CollectService{measureRepo: measureRepo}
}

// ProcessParameterData processes incoming data for a specific parameter and saves it.
func (s *CollectService) ProcessParameterData(ctx context.Context, parameterData dto.ParameterValuesBatch) error {
	const layout = "2006-01-02T15:04:05"
	timestamp, err := time.Parse(layout, parameterData.Timestamp)
	if err != nil {
		log.Printf("Error parsing timestamp: %v", err)
		return err
	}

	measure := entities.Measure{
		ParameterID: parameterData.ParameterID,
		Timestamp:   timestamp,
		Value:       parameterData.Value,
	}
	fmt.Println("PRocesss")
	fmt.Println(measure)
	if err := s.measureRepo.Save(ctx, measure); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
