package listeners

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/fabriciolfj/loan-service-go/configuration"
	"github.com/fabriciolfj/loan-service-go/entities"
	"github.com/magiconair/properties"
	"log"
)

type LoanProcessListener struct {
	consumer sarama.ConsumerGroup
	topic    string
}

func ProviderLoanProcessListener(cfg *configuration.KafkaConfig) (*LoanProcessListener, error) {
	prop, err := properties.LoadFile("config.properties", properties.UTF8)

	if err != nil {
		panic(err)
	}

	listener := &LoanProcessListener{
		consumer: cfg.Consumer,
		topic:    prop.GetString("topic.request.process.loan", ""),
	}

	return listener, nil
}

func (c *LoanProcessListener) Start() error {
	ctx := context.Background()

	for {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		err := c.consumer.Consume(ctx, []string{c.topic}, c)
		if err != nil {
			return fmt.Errorf("erro ao consumir mensagem: %w", err)
		}
	}
}

func (c *LoanProcessListener) Setup(sarama.ConsumerGroupSession) error {
	log.Println("init consumer group...")
	return nil
}

func (c *LoanProcessListener) Cleanup(sarama.ConsumerGroupSession) error {
	log.Println("finish consumer group...")
	return nil
}

func (c *LoanProcessListener) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		if err := c.handleMessage(message); err != nil {
			log.Printf("Erro ao processar mensagem: %v", err)
		}

		session.MarkMessage(message, "")
	}
	return nil
}

func (c *LoanProcessListener) Close() error {
	if err := c.consumer.Close(); err != nil {
		return fmt.Errorf("erro close consumer group: %w", err)
	}
	log.Println("close consumer group...")
	return nil
}

func (c *LoanProcessListener) handleMessage(message *sarama.ConsumerMessage) error {
	log.Printf("message received - Tópico: %s, Partição: %d, Offset: %d",
		message.Topic, message.Partition, message.Offset)

	var loan entities.Loan
	if err := json.Unmarshal(message.Value, &loan); err != nil {
		return fmt.Errorf("error deserializer message: %w", err)
	}

	log.Printf("message process success - ID: %s", loan.Code)
	return nil
}
