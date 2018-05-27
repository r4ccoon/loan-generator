package main

import (
	"testing"
	"time"
)

func TestGenerateLoanPayment(t *testing.T) {
	pPlan, _ := GenerateLoanPayment(2, 5.00, 5000, time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC))

	if len(pPlan) != 24 {
		t.Error("payment plan should be 24 times")
	}

}
