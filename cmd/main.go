package main

import (
	"Ingester/internal/domain/interfaces"
	"Ingester/internal/domain/services"
	"Ingester/internal/infrastructure/influxdb"
	"Ingester/internal/infrastructure/kafka"
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// InfluxDB setup
	influxDBClient := influxdb2.NewClient(os.Getenv("INFLUXDB_URL"), os.Getenv("INFLUXDB_TOKEN"))
	measureRepo := influxdb.NewMeasureRepositoryImpl(influxDBClient, "SparkSentry", "collect-raw-data")

	// Kafka setup
	brokerAddress := os.Getenv("KAFKA_BROKER_ADDRESS")
	topic := os.Getenv("KAFKA_TOPIC")
	groupID := os.Getenv("KAFKA_GROUP_ID")

	consumer, err := kafka.NewConsumer(brokerAddress, topic, groupID)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %s", err)
	}

	collectService := services.NewCollectService(measureRepo)
	handler := interfaces.NewKafkaConsumerHandler(collectService)

	// Start consuming
	fmt.Println("Starting the consumer...")
	err = consumer.Consume(context.Background(), func(msg []byte) {
		handler.HandleMessage(msg)
	})
	if err != nil {
		log.Fatalf("Failed to start consumer: %s", err)
	}
}
