package main

/*
[Loan Calculator - Stage 2/4: ????](https://hyperskill.org/projects/???/stages/???/implement)
-------------------------------------------------------------------------------
[Introduction to Go](https://hyperskill.org/learn/step/17073)
[Variables and constants](https://hyperskill.org/learn/step/14468)
[Input/Output](https://hyperskill.org/learn/step/14527)
[GoLand](https://hyperskill.org/learn/step/19072)
[Functions](https://hyperskill.org/learn/step/16388)
[Control statements](https://hyperskill.org/learn/step/16235)
*/

import (
	"fmt"
	"math"
)

const (
	enterLoan      = "Enter the loan principal:"
	whatToCalcText = `What do you want to calculate?
type "m" - for number of monthly payments,
type "p" - for the monthly payment:`
	enterMonthlyPayment = "Enter the monthly payment:"
	enterNumberOfMonths = "Enter the number of months:"
)

func main() {
	fmt.Println(enterLoan)

	var loan int
	fmt.Scanln(&loan)

	fmt.Println(whatToCalcText)
	var whatToCalc string
	fmt.Scanln(&whatToCalc)

	switch whatToCalc {
	case "m":
		calcMonths(loan)
	case "p":
		calcPayment(loan)
	}
}

func calcMonths(loan int) {
	fmt.Println(enterMonthlyPayment)
	var payment int
	fmt.Scanln(&payment)

	months := int(math.Ceil(float64(loan) / float64(payment)))

	fmt.Printf("It will take %d ", months)
	if months == 1 {
		fmt.Print("month ")
	} else {
		fmt.Print("months ")
	}
	fmt.Println("to repay the loan")
}

func calcPayment(loan int) {
	fmt.Println(enterNumberOfMonths)
	var months int
	fmt.Scanln(&months)

	payment := int(math.Ceil(float64(loan) / float64(months)))
	if payment != loan/months {
		payment = loan/months + 1
		fmt.Println("Your monthly payment =", payment, "and the last payment =", loan-(payment*(months-1)))
	} else {
		fmt.Println("Your monthly payment =", payment)
	}
}
