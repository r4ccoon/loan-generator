package main

import (
	"fmt"
	"time"

	"./loanpayment"
)

func main() {
	payments, _ := loanpayment.GenerateLoanPayment(2, 5, 5000, time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC))

	for i := 0; i < len(payments); i++ {
		monthPay := payments[i]
		fmt.Printf("%s,%.2f,%.2f,%.2f,%.2f,%.2f \n",
			monthPay.Date.Format("02.01.2006"),
			monthPay.Annuity,
			monthPay.Principal,
			monthPay.Interest,
			monthPay.InitialOustandingPrincipal,
			monthPay.RemainingPrincipal)
	}
}
