package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/basebandit/rpncalculator"
)

func main() {
	fmt.Fprintln(os.Stdout, "Reverse Polish Calculator")
	fmt.Fprintln(os.Stdout, "Currently only supporting +-/* operations. Example: 3 4 + *")
	fmt.Fprintln(os.Stdout, "Type 'exit' at any time to quit")
	var input string

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Fprintln(os.Stdout, "Input a RPN expression.")
		fmt.Fprint(os.Stdout, "-> ")

		scanner.Scan()
		input = scanner.Text()
		if input == "exit" {
			break
		}
		if len(input) == 0 {
			continue
		}

		result, err := rpncalculator.Evaluate(input)
		if err != nil {
			fmt.Fprintf(os.Stdout, "oops, %v. Please Try again\n", err)
		}
		fmt.Fprintln(os.Stdout, result)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
		os.Exit(1)
	}
}
