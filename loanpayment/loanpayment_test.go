package loanpayment

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateLoanPayment(t *testing.T) {
	pPlan, _ := GenerateLoanPayment(2, 5.00, 5000, time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC))

	if len(pPlan) != 24 {
		t.Error("payment plan should be 24 times")
	}

	if pPlan[0].principal != 198.52361533700798 {
		t.Error("first principle should be 198.52361533700798")
	}

	zeroVar := fmt.Sprintf("%.0f", pPlan[23].remainingPrincipal)
	if zeroVar != "0" {
		t.Error(zeroVar)
		t.Error("last remaining must be 0")
	}
}
