//go:build wireinject
// +build wireinject

package main

import (
	"github.com/fabriciolfj/loan-service-go/adapters"
	"github.com/fabriciolfj/loan-service-go/configuration"
	"github.com/fabriciolfj/loan-service-go/controller"
	"github.com/fabriciolfj/loan-service-go/listeners"
	"github.com/fabriciolfj/loan-service-go/producers"
	"github.com/fabriciolfj/loan-service-go/usecases"
	"github.com/google/wire"
)

func InitControllerLoan() (*controller.LoanController, error) {
	wire.Build(
		configuration.KafkaProperties,
		configuration.KafkaConfig,
		producers.ProviderLoanProcessProducer,
		adapters.LoanNotifyPendingAdapter,
		usecases.ProviderProcessLoanUseCase,
		controller.LoanController)

	return &controller.LoanController{}, nil
}

func InitListenerProcessLoan() (*listeners.LoanProcessListener, error) {
	wire.Build(
		configuration.ProvideKafkaProperties,
		configuration.ProvideKafkaConfig,
		listeners.ProviderLoanProcessListener,
	)

	return &listeners.LoanProcessListener{}, nil
}
