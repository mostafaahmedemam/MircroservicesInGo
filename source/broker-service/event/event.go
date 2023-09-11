package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"log_topic", // name
		"topic",     //type
		true,        //durable?
		false,       //auto-deleted
		false,       //internal
		false,       //no-wait
		nil,         //arguments
	)

}
func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    //name
		false, //durable
		false, //delete when unused
		false, // exclusive --> delete when connection close, and can connect with queue with that name
		false, //no-wait
		nil,   //arguments
	)

}
