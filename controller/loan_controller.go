package controller

import (
	"encoding/json"
	dtorequestloan "github.com/fabriciolfj/loan-service-go/model_loan_process"
	"github.com/fabriciolfj/loan-service-go/usecases"
	"net/http"
)

type LoanController struct {
	processUseCase *usecases.ProcessLoanUseCase
}

func ProviderLoanController(processUseCase *usecases.ProcessLoanUseCase) (*LoanController, error) {
	return &LoanController{
		processUseCase: processUseCase,
	}, nil
}

func (controller *LoanController) create(w http.ResponseWriter, r *http.Request) {
	var request dtorequestloan.LoanRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		badRequest(w, err)
	}

	entity := request.ToEntity()
	err = controller.processUseCase.Execute(entity)
	if err != nil {
		badRequest(w, err)
	}
}

func badRequest(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
}

func (pc *LoanController) HandlerProduct(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		pc.create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
