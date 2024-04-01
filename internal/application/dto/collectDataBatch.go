package dto

// ParameterValuesBatch represents a batch of values for a specific parameter.
type ParameterValuesBatch struct {
	ParameterID string  `json:"parameterId"` // Unique identifier for the parameter.
	Timestamp   string  `json:"timestamp"`   // Timestamp of the data collection.
	Value       float64 `json:"value"`       // Collected value for the parameter.
}
