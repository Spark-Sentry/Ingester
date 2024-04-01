package influxdb

import (
	"Ingester/internal/domain/entities"
	"Ingester/internal/domain/repositories"
	"context"
	"fmt"
	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"log"
)

// MeasureRepositoryImpl implements the MeasureRepository interface using InfluxDB.
type MeasureRepositoryImpl struct {
	client influxdb2.Client
	org    string
	bucket string
}

// NewMeasureRepositoryImpl creates a new instance of MeasureRepositoryImpl.
func NewMeasureRepositoryImpl(client influxdb2.Client, org, bucket string) repositories.MeasureRepository {
	return &MeasureRepositoryImpl{
		client: client,
		org:    org,
		bucket: bucket,
	}
}

// Save writes a measure to InfluxDB.
func (r *MeasureRepositoryImpl) Save(ctx context.Context, measure entities.Measure) error {
	writeAPI := r.client.WriteAPIBlocking(r.org, r.bucket)
	u := uuid.New().String()

	point := write.NewPointWithMeasurement(measure.ParameterID).
		AddTag("parameterId", measure.ParameterID).
		AddField("value", measure.Value).
		AddTag("uniqueId", u).
		SetTime(measure.Timestamp)

	if err := writeAPI.WritePoint(ctx, point); err != nil {
		log.Printf("Failed to write measure to InfluxDB: %v", err)
		return err
	}

	fmt.Printf("Measure written to InfluxDB: %v\n", measure)
	return nil
}
