package mq

import (
	"encoding/json"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {

	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}

}

// Emit a client request message to the RabbitMQ AMQP port.
func Emit(requestMessage []byte) ([]byte, error) {

	rabbitmqUsername := os.Getenv("RABBITMQ_USERNAME")
	rabbitmqPassword := os.Getenv("RABBITMQ_PASSWORD")
	rabbitmqServiceName := os.Getenv("RABBITMQ_SERVICE_NAME")
	rabbitmqAmqpPort := os.Getenv("RABBITMQ_AMQP_PORT")

	conn, err := amqp.Dial("amqp://" + rabbitmqUsername + ":" + rabbitmqPassword + "@" + rabbitmqServiceName + ":" + rabbitmqAmqpPort)
	failOnError(err, "Failed to connect to RabbitMQ AMQP")

	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"request_exchange", // name
		"topic",            // type
		true,               // durable
		false,              // auto-deleted
		false,              // internal
		false,              // no wait
		nil,                // argument
	)
	failOnError(err, "Failed to declare an exchange")

	err = ch.Publish(
		"request_exchange",    // exchange
		"request_routing_key", // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        requestMessage,
		},
	)
	failOnError(err, "Failed to published message")

	// DEBUG LOG
	log.Printf(" [x] Sent %s", requestMessage)

	// TEST SERVER RESPONSE
	resp := struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data"`
	}{
		Status: 200,
		Data: struct {
			Firstname string `json:"firstname"`
			Lastname  string `json:"lastname"`
		}{
			Firstname: "Thanakrit",
			Lastname:  "Lee",
		},
	}

	respByte, err := json.Marshal(resp)
	failOnError(err, "Failed to marshal response")

	return respByte, nil

}
