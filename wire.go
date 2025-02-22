//go:build wireinject
// +build wireinject

package main

import (
	"github.com/fabriciolfj/loan-service-go/adapters"
	"github.com/fabriciolfj/loan-service-go/configuration"
	"github.com/fabriciolfj/loan-service-go/controller"
	"github.com/fabriciolfj/loan-service-go/listeners"
	"github.com/fabriciolfj/loan-service-go/producers"
	"github.com/fabriciolfj/loan-service-go/repositories"
	"github.com/fabriciolfj/loan-service-go/usecases"
	"github.com/google/wire"
)

func InitControllerLoanWire() (*controller.LoanController, error) {
	wire.Build(
		configuration.ProviderDataBase,
		configuration.ProvideKafkaProperties,
		configuration.ProvideKafkaConfig,
		producers.ProviderLoanProcessProducer,
		repositories.ProviderLoanRepository,
		adapters.ProvideLoanNotifyPendingAdapter,
		adapters.ProvideFindLoanAdapter,
		adapters.ProvideSaveLoanAdapter,
		wire.Bind(new(usecases.SaveLoan), new(*adapters.SaveLoanAdapter)),
		wire.Bind(new(usecases.NotifyLoanPending), new(*adapters.LoanNotifyPendingAdapter)),
		wire.Bind(new(usecases.FindLoan), new(*adapters.FindLoanAdapter)),
		usecases.ProviderProcessLoanUseCase,
		usecases.ProviderGetLoanUseCase,
		controller.ProviderLoanController)

	return nil, nil
}

func InitListenerProcessLoanWire() (*listeners.LoanProcessListener, error) {
	wire.Build(
		configuration.ProvideKafkaProperties,
		configuration.ProvideKafkaConfig,
		listeners.ProviderLoanProcessListener,
	)

	return nil, nil
}
