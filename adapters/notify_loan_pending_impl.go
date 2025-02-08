package adapters

import (
	"encoding/json"
	"github.com/fabriciolfj/loan-service-go/entities"
	"github.com/fabriciolfj/loan-service-go/producers"
	"github.com/magiconair/properties"
	"log"
)

type LoanNotifyPendingAdapter struct {
	producer *producers.LoanProcessProducer
}

func ProvideLoanNotifyPendingAdapter(producer *producers.LoanProcessProducer) *LoanNotifyPendingAdapter {
	return &LoanNotifyPendingAdapter{
		producer: producer,
	}
}

func (adapter *LoanNotifyPendingAdapter) NotifyLoanPending(entity entities.Loan) error {
	prop, err := properties.LoadFile("config.properties", properties.UTF8)
	if err != nil {
		log.Fatal("error loading properties topic nofify loan pending:", err)
		return err
	}

	topic := prop.GetString("topic.request.process.loan", "")

	message, err := json.Marshal(entity)
	if err != nil {
		log.Fatal("failed to marshal loan: %w", err)
		return err
	}

	return adapter.producer.SendMessage(string(message), topic)
}
