package main

import (
	"fmt"
	"net/url"
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

func DesiredCookingTime() string {
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
			switch commandInt {
			case 1:
				return "durations" + url.QueryEscape("[]") + "=15"
			case 2:
				return "durations" + url.QueryEscape("[]") + "=30"
			case 3:
				return "durations" + url.QueryEscape("[]") + "=45"
			case 4:
				return "durations" + url.QueryEscape("[]") + "=60"
			case 5:
				return ""
			}
		}

		fmt.Println("Unknown command")
	}
}

func MealTime() string {
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
			switch commandInt {
			case 1:
				return url.QueryEscape("tags[recipe_mealtime][]") + "=" + url.QueryEscape("завтрак")
			case 2:
				return url.QueryEscape("tags[recipe_mealtime][]") + "=" + url.QueryEscape("второй+завтрак")
			case 3:
				return url.QueryEscape("tags[recipe_mealtime][]") + "=" + url.QueryEscape("ссобойки")
			case 4:
				return url.QueryEscape("tags[recipe_mealtime][]") + "=" + url.QueryEscape("полдник")
			case 5:
				return url.QueryEscape("tags[recipe_mealtime][]") + "=" + url.QueryEscape("ужин")
			case 6:
				return ""
			}
		}

		fmt.Println("Unknown command")
	}
}

func Holiday() string {
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
			switch commandInt {
			case 1:
				return url.QueryEscape("tags[recipe_holiday][]") + "=" + url.QueryEscape("масленица")
			case 2:
				return url.QueryEscape("tags[recipe_holiday][]") + "=" + url.QueryEscape("новый+год")
			case 3:
				return url.QueryEscape("tags[recipe_holiday][]") + "=" + url.QueryEscape("пасха")
			case 4:
				return url.QueryEscape("tags[recipe_holiday][]") + "=" + url.QueryEscape("пост")
			case 5:
				return ""
			}
		}

		fmt.Println("Unknown command")
	}
}

func NumOfRecipes() int {
	for {
		fmt.Println("Enter the number of recipes to download: (integer)")

		var command string
		_, _ = fmt.Fscan(os.Stdin, &command)
		if commandInt, err := strconv.Atoi(command); err == nil && commandInt >= 1 {
			fmt.Println()
			return commandInt
		}

		fmt.Println("Unknown command")
	}
}

func NewReqStr(time string, mealTime string, holiday string) string {
	return "https://www.edimdoma.ru/retsepty?" + time + "&" + mealTime + "&" + holiday
}
