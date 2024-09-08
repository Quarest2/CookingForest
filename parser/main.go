package main

import (
	"CookingForest/parser/parser"
	"CookingForest/parser/request"
	"fmt"
	"io"
)

func main() {
	var err error
	var rBody io.Reader

	isEnd := false
	for !isEnd {
		cookingTime := DesiredCookingTime()
		mealTime := MealTime()
		holiday := Holiday()
		n := NumOfRecipes()
		url := NewReqStr(cookingTime, mealTime, holiday)

		if rBody, err = request.GetBody(url); err != nil {
			fmt.Println(err)
		}
		parser.Parse(rBody, n)
		isEnd = EndOfWork()
	}
}
