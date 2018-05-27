package loanpayment

import (
	"errors"
	"math"
	"time"
)

// PaymentPlan repayment plan struct for borrowers
type PaymentPlan struct {
	date                       time.Time
	annuity                    float64
	principal                  float64
	interest                   float64
	initialOustandingPrincipal float64
	remainingPrincipal         float64
}

// GenerateLoanPayment generate loan payment
func GenerateLoanPayment(durationYear int, interestPercent float64, totalLoan float64, startDate time.Time) ([]PaymentPlan, error) {
	var loanPlans []PaymentPlan

	if durationYear < 1 {
		return loanPlans, errors.New("duration cannot be less than 1")
	}

	if totalLoan <= 0 {
		return loanPlans, errors.New("total loan must be more than 0")
	}

	if interestPercent <= 0 {
		return loanPlans, errors.New("interest rate cannot be less than 0")
	}

	// periods = duration * 12 month
	periods := durationYear * 12
	// decimal version of the interest rat
	ratePeriod := (interestPercent / 100) * 30 / 360
	annuity := (ratePeriod * totalLoan) / (1 - math.Pow(1+ratePeriod, float64(periods*-1)))

	outstanding := totalLoan
	for i := 0; i < periods; i++ {
		date := startDate.AddDate(0, i, 0)

		interest := ((interestPercent / 100) * 30 * outstanding) / 360
		principal := annuity - interest

		remaining := outstanding - principal

		payment := 0.0
		if remaining > outstanding {
			payment = outstanding
			remaining = outstanding - payment
		} else {
			payment = principal + interest
		}

		loanPlans = append(loanPlans, PaymentPlan{date, payment, principal, interest, outstanding, remaining})

		outstanding = outstanding - principal
	}

	return loanPlans, nil
}
