package dto

// MessageData represents the structure of decoded Kafka messages.
type MessageData struct {
	EquipmentID string `json:"equipmentId"`
	Timestamp   string `json:"timestamp"`
	Parameter   struct {
		ParameterID string  `json:"parameterId"`
		Value       float64 `json:"value"`
	} `json:"parameter"`
}
