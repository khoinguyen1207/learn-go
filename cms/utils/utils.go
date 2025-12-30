package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func ReadInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func GetPositiveInt(prompt string) int {
	for {
		input := ReadInput(prompt)

		output, err := strconv.Atoi(input)
		if err == nil && output > 0 {
			return output
		}
		fmt.Println("Input invalid")
	}
}

func GetPositiveFloat(prompt string) float64 {
	for {
		input := ReadInput(prompt)

		output, err := strconv.ParseFloat(input, 64)
		if err == nil && output > 0 {
			return output
		}
		fmt.Println("Input invalid")
	}
}

func GetOptionalPositiveFloat(prompt string, prevData float64) float64 {
	input := ReadInput(prompt)
	output, err := strconv.ParseFloat(input, 64)
	if err != nil && output <= 0 {
		return prevData
	}
	return output
}

func GetNonEmptyString(prompt string) string {
	for {
		input := ReadInput(prompt)
		if input != "" {
			return input
		}
		fmt.Println("Input cannot be empty")
	}
}

func GetOptionalString(prompt string, prevData string) string {
	input := ReadInput(prompt)
	if input == "" {
		return prevData
	}
	return input
}

func ClearScreen() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

type HasId interface {
	GetId() int
}

func IsIdUnique[T HasId](id int, list []T) bool {
	for _, item := range list {
		if item.GetId() == id {
			return false
		}
	}
	return true
}
