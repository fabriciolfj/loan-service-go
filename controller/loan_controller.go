package controller

import (
	"encoding/json"
	dtorequestloan "github.com/fabriciolfj/loan-service-go/model_loan_process"
	"github.com/fabriciolfj/loan-service-go/usecases"
	"net/http"
)

type LoanController struct {
	processUseCase  *usecases.ProcessLoanUseCase
	findLoanUseCase *usecases.GetLoanUseCase
}

func ProviderLoanController(processUseCase *usecases.ProcessLoanUseCase, findLoanUseCase *usecases.GetLoanUseCase) (*LoanController, error) {
	return &LoanController{
		processUseCase:  processUseCase,
		findLoanUseCase: findLoanUseCase,
	}, nil
}

func (controller *LoanController) find(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	result, err := controller.findLoanUseCase.Execute(code)
	if err != nil {
		badRequest(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(result); err != nil {
		badRequest(w, err)
		return
	}
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

func (pc *LoanController) HandlerLoan(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		pc.create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
