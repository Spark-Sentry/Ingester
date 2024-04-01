package entities

import (
	"time"
)

// Measure represents the data structure for a measurement from a device.
type Measure struct {
	MeasurementName string
	Timestamp       time.Time // Timestamp when the measurement was taken
	ParameterID     string    // Identifier of the measured parameter
	Value           float64   // Value of the measurement
}
