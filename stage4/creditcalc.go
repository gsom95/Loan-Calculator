package main

/*
[Loan Calculator - Stage 4/4: ????](https://hyperskill.org/projects/???/stages/???/implement)
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
	paymentType = flag.String("type", "", "type of payment: 'annuity' or 'diff'")
	payment     = flag.Float64("payment", 0, "payment amount")
	principal   = flag.Int("principal", 0, "principal amount")
	periods     = flag.Int("periods", 0, "number of months needed to repay the loan")
	interest    = flag.Float64("interest", 0, "interest rate, required parameter")
)

func main() {
	flag.CommandLine.SetOutput(os.Stdout)
	flag.Parse()

	switch {
	case
		len(os.Args) != 5,
		*paymentType != "annuity" && *paymentType != "diff",
		*paymentType == "diff" && *payment != 0,
		*interest == 0,
		*payment < 0 || *principal < 0 || *periods < 0 || *interest < 0:
		fmt.Println("Incorrect parameters")
		return
	}

	if *interest <= 0 {
		fmt.Printf("Wrong interest value ('%f'). Possible parameters:\n", *interest)
		flag.PrintDefaults()
		return
	}

	if *paymentType == "diff" {
		diffPayment(*principal, *periods, *interest)
		return
	}

	switch {
	case *periods == 0:
		monthlyPayment(*principal, *payment, *interest)
	case *payment == 0:
		annuityPayment(*principal, *periods, *interest)
	case *principal == 0:
		loanPrincipal(*payment, *periods, *interest)
	}
}

func diffPayment(principal, periods int, interest float64) {
	monthlyInterestRate := interest / 100 / 12
	total := 0.0
	for i := 1; i <= periods; i++ {
		payment := math.Ceil(
			(float64(principal) / float64(periods)) +
				(monthlyInterestRate * (float64(principal) - (float64(principal)*(float64(i)-1))/float64(periods))),
		)
		total += payment
		fmt.Printf("Month %d: payment is %d\n", i, int(payment))
	}
	fmt.Println()
	fmt.Println("Overpayment = ", int(total)-principal)
}

func monthlyPayment(principal int, annuityPayment, interestRate float64) {
	monthlyInterest := interestRate / 100 / 12

	months := int(math.Ceil(
		math.Log(annuityPayment/
			(annuityPayment-monthlyInterest*float64(principal)),
		) / math.Log(1+monthlyInterest),
	))
	overPayment := annuityPayment*float64(months) - float64(principal)

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
	fmt.Println("Overpayment = ", overPayment)
}

func annuityPayment(principal, numberOfMonths int, interest float64) {
	monthlyInterestRate := interest / 100 / 12
	paymentPerMonth := int(math.Ceil(
		float64(principal) *
			(monthlyInterestRate * math.Pow(1+monthlyInterestRate, float64(numberOfMonths))) /
			(math.Pow(1+monthlyInterestRate, float64(numberOfMonths)) - 1),
	))

	fmt.Println("Your annuity payment = ", paymentPerMonth)
	fmt.Println("Overpayment = ", paymentPerMonth*numberOfMonths-principal)
}

func loanPrincipal(annuityPayment float64, numberOfMonths int, interest float64) {
	monthlyInterestRate := interest / 100 / 12
	principal := int(math.Floor(
		annuityPayment /
			((monthlyInterestRate * math.Pow(1+monthlyInterestRate, float64(numberOfMonths))) /
				(math.Pow(1+monthlyInterestRate, float64(numberOfMonths)) - 1)),
	))

	fmt.Println("Your loan principal = ", principal)
	fmt.Println("Overpayment = ", annuityPayment*float64(numberOfMonths)-float64(principal))
}
