package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearConsole() {
	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
	}
}

func ReadUserInput(prompt string) string {
	fmt.Print(prompt)

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		return scanner.Text()
	}

	return ""
}
