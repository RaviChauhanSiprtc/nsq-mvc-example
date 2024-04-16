// controllers/consume_controller.go
package controllers

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
)

func ConsumeMessageHandler(w http.ResponseWriter, r *http.Request) {
	// Connect to NSQ
	consumer, err := nsq.NewConsumer("test-topic", "test-channel", nsq.NewConfig())
	if err != nil {
		http.Error(w, "Failed to connect to NSQ", http.StatusInternalServerError)
		return
	}

	// Handle NSQ messages
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		fmt.Printf("Received message: %s\n", message.Body)
		return nil
	}))

	// Connect consumer to NSQ
	err = consumer.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		http.Error(w, "Failed to connect to NSQD", http.StatusInternalServerError)
		return
	}

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	consumer.Stop()

	w.Write([]byte("Messages consumed from NSQ"))
}
