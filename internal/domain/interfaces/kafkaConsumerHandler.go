package interfaces

import (
	"Ingester/internal/application/dto"
	"Ingester/internal/domain/services"
	"context"
	"encoding/json"
	"log"
)

type KafkaConsumerHandler struct {
	collectService *services.CollectService
}

func NewKafkaConsumerHandler(collectService *services.CollectService) *KafkaConsumerHandler {
	return &KafkaConsumerHandler{collectService: collectService}
}

// HandleMessage decodes Kafka messages and processes them as Measure entities.
func (h *KafkaConsumerHandler) HandleMessage(msg []byte) {
	var data dto.ParameterValuesBatch
	if err := json.Unmarshal(msg, &data); err != nil {
		log.Printf("Failed to unmarshal message: %v", err)
		return
	}

	if err := h.collectService.ProcessParameterData(context.Background(), data); err != nil {
		log.Printf("Failed to process equipment data: %v", err)
	}
}
