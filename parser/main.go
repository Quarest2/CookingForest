package main

func main() {
	isEnd := false
	for !isEnd {
		cookingTime := DesiredCookingTime()
		mealTime := MealTime()
		holiday := Holiday()
		isEnd = EndOfWork()
	}
}
