// controllers/produce_controller.go
package controllers

import (
	"net/http"

	"github.com/nsqio/go-nsq"
)

func ProduceMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Connect to NSQ
	producer, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
	if err != nil {
		http.Error(w, "Failed to connect to NSQ", http.StatusInternalServerError)
		return
	}
	defer producer.Stop()

	// Publish message to NSQ
	err = producer.Publish("test-topic", []byte("Hello, NSQ!"))
	if err != nil {
		http.Error(w, "Failed to publish message to NSQ", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Message sent to NSQ"))
}
