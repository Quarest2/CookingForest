package main

import (
	"CookingForest/reviewer/storage"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetCommand() int {
	fmt.Println("Choose a command (integer):")
	fmt.Println("1. Add a new recipe")
	fmt.Println("2. Find recipes")
	fmt.Println("3. Export recipes")

	var command string
	_, _ = fmt.Fscan(os.Stdin, &command)

	var comInt int
	var err error
	if comInt, err = strconv.Atoi(command); err != nil || comInt < 1 || comInt > 3 {
		fmt.Println("Unknown command.")
	}
	fmt.Println()

	return comInt
}

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

func UploadRecipeClient() (storage.Recipe, error) {
	var name string
	var imagePath string
	var hasImage bool
	var persons int
	var time string
	var steps []string
	var err error
	var text string

	fmt.Println("Write a name of a dish:")
	if _, err = fmt.Fscan(os.Stdin, &name); err != nil {
		fmt.Printf("Error reading name: %v\n", err)
		return storage.Recipe{}, err
	}

	fmt.Println("Write how many people the dish is for:")
	if _, err = fmt.Fscan(os.Stdin, &persons); err != nil {
		fmt.Printf("Error reading persons: %v\n", err)
		return storage.Recipe{}, err
	}

	fmt.Println("Write how long it takes to prepare the dish:")
	if time, err = bufio.NewReader(os.Stdin).ReadString('\n'); err != nil {
		fmt.Printf("Error reading time: %v\n", err)
		return storage.Recipe{}, err
	}

	var n int
	fmt.Println("Write how many steps there are in the recipe:")
	if _, err = fmt.Fscan(os.Stdin, &n); err != nil {
		fmt.Printf("Error reading number of steps: %v\n", err)
		return storage.Recipe{}, err
	}
	for i := 1; i <= n; i++ {
		fmt.Printf("Write step number %d:\n", i)
		if text, err = bufio.NewReader(os.Stdin).ReadString('\n'); err != nil {
			fmt.Printf("Error reading input in steps: %v\n", err)
			return storage.Recipe{}, err
		}
		steps = append(steps, text)
	}

	fmt.Println("Do you want to add an image? (y/n)")
	if _, _ = fmt.Fscan(os.Stdin, &text); text == "y" {
		fmt.Println("Write an absolute path to the image")
		if _, err = fmt.Fscan(os.Stdin, &imagePath); err != nil {
			fmt.Printf("Error reading image path: %v\n", err)
			return storage.Recipe{}, err
		}
		if err = storage.UploadImage(imagePath, name); err != nil {
			fmt.Println("Error with uploading image", err)
			return storage.Recipe{}, err
		}
		hasImage = true
	}

	r := storage.Recipe{Name: name, HasImage: hasImage, Time: time, Persons: persons, Steps: steps}
	ok := storage.UploadRecipe(r)
	if ok {
		fmt.Println("Recipe successfully uploaded")
		return r, nil
	} else {
		fmt.Println("Error with uploading recipe")
		return storage.Recipe{}, errors.New("unknown error")
	}
}
