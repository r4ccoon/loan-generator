package main

import (
	"fmt"
	"time"

	"github.com/r4ccoon/loan-generator/loanpayment"
)

func main() {
	payments, _ := loanpayment.GenerateLoanPayment(2, 5, 5000, time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC))

	fmt.Println("No. | Annuity | Principal | Interest | Initial | Remaining")

	for i := 0; i < len(payments); i++ {
		monthPay := payments[i]
		fmt.Printf("payment %d  | %s | %.2f | %.2f | %.2f | %.2f | %.2f \n",
			i+1,
			monthPay.date.Format("02.01.2006"),
			monthPay.annuity,
			monthPay.principal,
			monthPay.interest,
			monthPay.initialOustandingPrincipal,
			monthPay.remainingPrincipal)
	}
}
