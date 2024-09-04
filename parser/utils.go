package main

import (
	"fmt"
	"os"
	"strconv"
)

func EndOfWork() bool {
	fmt.Println("If you want to do so, enter 'exit' to exit program")

	var command string
	if _, _ = fmt.Fscan(os.Stdin, &command); command == "exit" {
		fmt.Println()
		return true
	}
	fmt.Println()
	return false
}

func DesiredCookingTime() int {
	for {
		fmt.Println("Select desired cooking time: (integer)")
		fmt.Println("1. No more than 15 minutes")
		fmt.Println("2. No more than 30 minutes")
		fmt.Println("3. No more than 45 minutes")
		fmt.Println("4. No more than 60 minutes")
		fmt.Println("5. Doesn't matter")

		var command string
		_, _ = fmt.Fscan(os.Stdin, &command)
		if commandInt, err := strconv.Atoi(command); err == nil && commandInt >= 1 && commandInt <= 5 {
			fmt.Println()
			return commandInt
		}

		fmt.Println("Unknown command")
	}
}

func MealTime() int {
	for {
		fmt.Println("Select desired meal time: (integer)")
		fmt.Println("1. Breakfast")
		fmt.Println("2. Second breakfast")
		fmt.Println("3. With you")
		fmt.Println("4. Afternoon tea")
		fmt.Println("5. Dinner")
		fmt.Println("6. Doesn't matter")

		var command string
		_, _ = fmt.Fscan(os.Stdin, &command)
		if commandInt, err := strconv.Atoi(command); err == nil && commandInt >= 1 && commandInt <= 6 {
			fmt.Println()
			return commandInt
		}

		fmt.Println("Unknown command")
	}
}

func Holiday() int {
	for {
		fmt.Println("Select desired holiday: (integer)")
		fmt.Println("1. Shrovetide")
		fmt.Println("2. New year")
		fmt.Println("3. Easter")
		fmt.Println("4. Lent")
		fmt.Println("5. Doesn't matter")

		var command string
		_, _ = fmt.Fscan(os.Stdin, &command)
		if commandInt, err := strconv.Atoi(command); err == nil && commandInt >= 1 && commandInt <= 5 {
			fmt.Println()
			return commandInt
		}

		fmt.Println("Unknown command")
	}
}
