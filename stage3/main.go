package main

/*
[Loan Calculator - Stage 3/4: ????](https://hyperskill.org/projects/???/stages/???/implement)
-------------------------------------------------------------------------------
[Introduction to Go](https://hyperskill.org/learn/step/17073)
[Variables and constants](https://hyperskill.org/learn/step/14468)
[Input/Output](https://hyperskill.org/learn/step/14527)
[GoLand](https://hyperskill.org/learn/step/19072)
[Functions](https://hyperskill.org/learn/step/16388)
[Control statements](https://hyperskill.org/learn/step/16235)
*/

import (
	"flag"
	"fmt"
	"math"
	"os"
)

var (
	payment   = flag.Float64("payment", 0, "payment amount")
	principal = flag.Int("principal", 0, "principal amount")
	periods   = flag.Int("periods", 0, "number of months needed to repay the loan")
	interest  = flag.Float64("interest", 0, "interest rate, required parameter")
)

func main() {
	flag.CommandLine.SetOutput(os.Stdout)
	flag.Parse()

	switch {
	case *periods == 0:
		monthlyPayment(*principal, *payment, *interest)
	case *payment == 0:
		annuityPayment(*principal, *periods, *interest)
	case *principal == 0:
		loanPrincipal(*payment, *periods, *interest)
	}
}

func monthlyPayment(principal int, annuityPayment, interestRate float64) {
	monthlyInterest := interestRate / 100 / 12

	months := int(math.Ceil(
		math.Log(annuityPayment/
			(annuityPayment-monthlyInterest*float64(principal)),
		) / math.Log(1+monthlyInterest),
	))

	fmt.Print("It will take ")
	years := months / 12
	if years > 1 {
		fmt.Printf("%d years ", years)
	} else if years == 1 {
		fmt.Print("1 year ")
	}
	months = months % 12

	if months > 1 {
		fmt.Printf("%d months ", months)
	} else if months == 1 {
		fmt.Print("1 month ")
	}

	fmt.Println("to repay this loan!")
}

func annuityPayment(principal, numberOfMonths int, interest float64) {
	monthlyInterestRate := interest / 100 / 12
	paymentPerMonth := int(math.Ceil(
		float64(principal) *
			(monthlyInterestRate * math.Pow(1+monthlyInterestRate, float64(numberOfMonths))) /
			(math.Pow(1+monthlyInterestRate, float64(numberOfMonths)) - 1),
	))

	fmt.Println("Your annuity payment = ", paymentPerMonth)
}

func loanPrincipal(annuityPayment float64, numberOfMonths int, interest float64) {
	monthlyInterestRate := interest / 100 / 12
	principal := int(math.Floor(
		annuityPayment /
			((monthlyInterestRate * math.Pow(1+monthlyInterestRate, float64(numberOfMonths))) /
				(math.Pow(1+monthlyInterestRate, float64(numberOfMonths)) - 1)),
	))

	fmt.Println("Your loan principal = ", principal)
}
