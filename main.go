package main

import (
	"fmt"
)

func main() {
	for true {
		input := ReadUserInput("gmath > ")

		if input == "exit" {
			break
		} else if input == "clean" {
			ClearConsole()
			continue
		}

		tokens := TokenizeString(input)

		ast := Parse(tokens)

		fmt.Printf("\t%g\n", ast.Eval())
	}
}
