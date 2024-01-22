package env

import (
	"fmt"
	"os"
)

var (
	PORT             string
	AmqpServerURL    string
	AmqpQueueName    string
	AWSRegion        string
	PSID             string
	PSIDTS           string
	TelegramApiToken string
)

func Start() error {
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "5001"
	}

	AmqpServerURL = os.Getenv("AMQP_SERVER_URL")
	if AmqpServerURL == "" {
		return fmt.Errorf("Error to load AMQP_SERVER_URL")
	}

	AmqpQueueName = os.Getenv("AMQP_QUEUE_NAME")
	if AmqpQueueName == "" {
		return fmt.Errorf("Error to load AMQP_QUEUE_NAME")
	}

	AWSRegion = os.Getenv("AWS_REGION")
	if AWSRegion == "" {
		return fmt.Errorf("Error to load AWS_REGION")
	}

	TelegramApiToken = os.Getenv("TELEGRAM_API_TOKEN")
	if TelegramApiToken == "" {
		return fmt.Errorf("Error to load TELEGRAM_API_TOKEN")
	}

	PSID = os.Getenv("BARD_PSID")
	if PSID == "" {
		return fmt.Errorf("Error to load BARD_PSID")
	}

	PSIDTS = os.Getenv("BARD_PSIDTS")
	if PSIDTS == "" {
		return fmt.Errorf("Error to load BARD_PSIDTS")
	}

	return nil
}
