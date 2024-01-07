package main

/*
[Loan Calculator - Stage 1/4: Beginning](https://hyperskill.org/projects/???/stages/???/implement)
-------------------------------------------------------------------------------
[Introduction to Go](https://hyperskill.org/learn/step/17073)
[Variables and constants](https://hyperskill.org/learn/step/14468)
[Input/Output](https://hyperskill.org/learn/step/14527)
[GoLand](https://hyperskill.org/learn/step/19072)
*/

import "fmt"

func main() {
	var (
		loanPrincipal = "Loan principal: 1000"
		finalOutput   = "The loan has been repaid!"
		firstMonth    = "Month 1: repaid 250"
		secondMonth   = "Month 2: repaid 250"
		thirdMonth    = "Month 3: repaid 500"
	)

	fmt.Println(loanPrincipal)
	fmt.Println(firstMonth)
	fmt.Println(secondMonth)
	fmt.Println(thirdMonth)
	fmt.Println(finalOutput)
}
