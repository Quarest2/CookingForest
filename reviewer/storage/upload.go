package storage

import (
	"fmt"
	"io"
	"os"
)

func UploadImage(imagePath string, name string) error {
	var file *os.File
	var err error
	if file, err = os.Open(imagePath); err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	imageBytes := make([]byte, 64)

	for {
		if _, err = file.Read(imageBytes); err == io.EOF {
			break
		}
	}

	var newFile *os.File
	if newFile, err = os.Create(fmt.Sprintf("/Users/Askar/GolandProjects/CookingForest/images/%s.png", name)); err != nil {
		fmt.Println("Unable to create file:", err)
		return err
	}
	defer newFile.Close()
	if _, err = newFile.Write(imageBytes); err != nil {
		fmt.Println("Unable to write image:", err)
		return err
	}

	fmt.Println("Image uploaded successfully.")
	return nil
}

func UploadRecipe(recipe Recipe) bool {
	alreadyExists := false
	for _, r := range recipes {
		if r.Name == recipe.Name && r.HasImage == recipe.HasImage && r.Time == recipe.Time && r.Persons == recipe.Persons {
			alreadyExists = true
			for i, step := range recipe.Steps {
				if r.Steps[i] != step {
					alreadyExists = false
					break
				}
			}
			if alreadyExists {
				fmt.Println("This recipe already exists in the storage")
				return false
			}
		}
	}

	recipes = append(recipes, Recipe{Name: recipe.Name, HasImage: recipe.HasImage, Time: recipe.Time, Persons: recipe.Persons, Steps: recipe.Steps})
	return true
}
