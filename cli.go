
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("error: division by zero is not allowed")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("error: invalid operator '%s'. Use +, -, *, or /", operator)
	}
}


func readInput(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}


func runCalculator() {
	fmt.Println("Welcome to the CLI Calculator")
	fmt.Println("Enter 'q' or 'quit' to exit at any time.")

	for {
		
		input1, err := readInput("Enter the first number: ")
		if err != nil || input1 == "q" || input1 == "quit" {
			break
		}
		num1, err := strconv.ParseFloat(input1, 64)
		if err != nil {
			fmt.Println("Error: Invalid number input for the first value. Please try again.")
			continue
		}

		
		operator, err := readInput("Enter the operator (+ - * /): ")
		if err != nil || operator == "q" || operator == "quit" {
			break
		}

		
		input2, err := readInput("Enter the second number: ")
		if err != nil || input2 == "q" || input2 == "quit" {
			break
		}
		num2, err := strconv.ParseFloat(input2, 64)
		if err != nil {
			fmt.Println("Error: Invalid number input for the second value. Please try again.")
			continue
		}

	
		result, err := calculate(num1, num2, operator)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Result: %v %s %v = %v\n", num1, operator, num2, result)
		}
		fmt.Println("-" + strings.Repeat("-", 30)) 
	}

	fmt.Println("Calculator session ended. Goodbye!")
}

func main() {
	runCalculator()
}
