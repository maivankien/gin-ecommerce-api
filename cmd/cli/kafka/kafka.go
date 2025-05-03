package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL   = "localhost:9092"
	kafkaTopic = "user_topic_vip"
)

// For consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		StartOffset:    kafka.LastOffset, // kafka.FirstOffset, kafka.LastOffset
	})
}

// For producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func newStockInfo(message string, typeStock string) *StockInfo {
	s := StockInfo{
		Message: message,
		Type:    typeStock,
	}
	return &s
}

func actionStockInfo(c *gin.Context) {
	s := newStockInfo(c.Query("message"), c.Query("type"))

	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = s

	jsonData, _ := json.Marshal(body)

	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonData),
	}

	err := kafkaProducer.WriteMessages(c.Request.Context(), msg)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func registerConsumerATC(id int) {
	kafkaGroupId := "consumer-group-"
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)

	defer reader.Close()

	fmt.Println("Consumer registered: ", id)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Error reading message: ", err)
			continue
		}

		fmt.Printf(
			"Consumer(%d), listen topic:%v, partition:%v, offset:%v, time:%d %s = %s\n",
			id, m.Topic, m.Partition, m.Offset, m.Time.Unix(), string(m.Key), string(m.Value),
		)
	}
}

func main() {
	r := gin.Default()

	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close()

	r.POST("/action", actionStockInfo)

	go registerConsumerATC(1)
	go registerConsumerATC(2)
	go registerConsumerATC(3)

	r.Run(":8080")
}
