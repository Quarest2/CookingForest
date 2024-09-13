package main

import (
	"CookingForest/parser/archive"
	"CookingForest/parser/parser"
	"CookingForest/parser/request"
	"fmt"
	"io"
)

func main() {
	var err error
	var rBody io.Reader
	var recipes []parser.Recipe

	isEnd := false
	for !isEnd {
		cookingTime := DesiredCookingTime()
		mealTime := MealTime()
		holiday := Holiday()
		n := NumOfRecipes()
		url := NewReqStr(cookingTime, mealTime, holiday)

		fmt.Println(url)
		if rBody, err = request.GetBody(url); err != nil {
			fmt.Println(err)
		}
		if recipes, err = parser.Parse(rBody, n); err != nil {
			fmt.Println(err)
		}
		if err = archive.CreateArchive(recipes); err != nil {
			fmt.Println(err)
		}
		isEnd = EndOfWork()
	}
}
