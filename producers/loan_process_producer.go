package producers

import (
	"github.com/IBM/sarama"
	"github.com/fabriciolfj/loan-service-go/configuration"
	"github.com/magiconair/properties"
)

type LoanProcessProducer struct {
	producer sarama.AsyncProducer
	topic    string
}

func ProviderLoanProcessProducer(cfg *configuration.KafkaConfig) (*LoanProcessProducer, error) {
	prod, _ := properties.LoadFile("config.properties", properties.UTF8)

	producer := &LoanProcessProducer{
		producer: cfg.Producer,
		topic:    prod.GetString("topic.request.process.loan", ""),
	}

	return producer, nil
}

func (p *LoanProcessProducer) Close() error {
	if p.producer != nil {
		return p.producer.Close()
	}
	return nil
}

func (p *LoanProcessProducer) SendMessage(message string, topic string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	p.producer.Input() <- msg

	select {
	case err := <-p.producer.Errors():
		return err.Err
	case <-p.producer.Successes():
		return nil
	}
}
