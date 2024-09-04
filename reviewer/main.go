package main

import "fmt"

func main() {
	isEnd := false
	for !isEnd {
		command := GetCommand()

		if command != 0 {
			switch command {
			case 1:
				if r, err := UploadRecipeClient(); err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("Recipe %s successfully uploaded\n", r.Name)
				}
			case 2:

			case 3:

			}
		}

		command = 0
		isEnd = EndOfWork()
	}
}
