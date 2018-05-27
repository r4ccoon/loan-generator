package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/r4ccoon/loan-generator/loanpayment"
)

// LoanRequest payload
type LoanRequest struct {
	LoanAmount  float64
	NominalRate float64
	Duration    int
	StartDate   string
}

func generatePaymentPlan(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		panic(err)
	}

	var loanRequest LoanRequest
	err = json.Unmarshal(body, &loanRequest)
	if err != nil {
		panic(err)
	}

	layout := "2006-01-02T15:04:05Z"
	startDate, err := time.Parse(layout, loanRequest.StartDate)
	payments, _ := loanpayment.GenerateLoanPayment(loanRequest.Duration/12, loanRequest.NominalRate, loanRequest.LoanAmount, startDate)

	result := make([]map[string]string, 0)

	for i := 0; i < len(payments); i++ {
		monthPay := payments[i]

		paymentFormatted := make(map[string]string)
		paymentFormatted["date"] = monthPay.Date.Format("02.01.2006")
		paymentFormatted["borrowerPaymentAmount"] = fmt.Sprintf("%.2f", monthPay.Annuity)
		paymentFormatted["principal"] = fmt.Sprintf("%.2f", monthPay.Principal)
		paymentFormatted["interest"] = fmt.Sprintf("%.2f", monthPay.Interest)
		paymentFormatted["initialOutstandingPrincipal"] = fmt.Sprintf("%.2f", monthPay.InitialOustandingPrincipal)
		paymentFormatted["remainingOutstandingPrincipal"] = fmt.Sprintf("%.2f", monthPay.RemainingPrincipal)
		result = append(result, paymentFormatted)
	}

	jsonMessage, _ := json.Marshal(result)

	fmt.Println("200")
	w.Write(jsonMessage)
}

func main() {
	http.HandleFunc("/generate-plan", generatePaymentPlan)
	fmt.Println("server started at port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
