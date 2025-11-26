package main

import (
	"fmt"
)

func main() {
	var number1, number2 float64
	var operator string

	fmt.Println("Welcome To CLI calculator")
	fmt.Println("Input Only Numbers")

	fmt.Print("Enter the first number : ")
	_, err1 := fmt.Scanln(&number1)
	if err1 != nil {
		fmt.Println(" Error: First input is NOT a number.")
		return
	}

	fmt.Print("Enter the second number : ")
	_, err2 := fmt.Scanln(&number2)
	if err2 != nil {
		fmt.Println("Error: Second input is NOT a number.")
		return
	}

	fmt.Print("Enter the operator ( + - * / ) : ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%v %s %v = %v\n", number1, operator, number2, number1+number2)

	case "-":
		fmt.Printf("%v %s %v = %v\n", number1, operator, number2, number1-number2)

	case "*":
		fmt.Printf("%v %s %v = %v\n", number1, operator, number2, number1*number2)

	case "/":
		if number2 == 0 {
			fmt.Println("Error: Division by zero.")
		} else {
			fmt.Printf("%v %s %v = %v\n", number1, operator, number2, number1/number2)
		}

	default:
		fmt.Println("Invalid operator")
	}
	fmt.Println()
}
