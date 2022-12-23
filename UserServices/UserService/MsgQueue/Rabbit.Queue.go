package MsgQueue

import (
	"UserServices/UserService/Model"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func SendMSG(raw Model.Verification) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Load .env File")
	}
	body := Model.VerificationMSG{
		Email: raw.Email,
		Token: raw.Token,
	}
	RABBIT_URL := os.Getenv("RABBIT_URI")
	RABBIT_QUEUE := os.Getenv("RABBIT_QUEUE")
	conn, err := amqp.Dial(RABBIT_URL)

	ch, err := conn.Channel()
	//fmt.Println("Conn Success")
	q, err := ch.QueueDeclare(
		RABBIT_QUEUE, // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)

	b, err := json.Marshal(body)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(b),
		})
	defer conn.Close()
	defer ch.Close()
}
