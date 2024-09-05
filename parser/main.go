package main

import "fmt"

func main() {
	isEnd := false
	for !isEnd {
		cookingTime := DesiredCookingTime()
		mealTime := MealTime()
		holiday := Holiday()
		req := NewReqStr(cookingTime, mealTime, holiday)
		fmt.Println(req)
		isEnd = EndOfWork()
	}
}
