package archive

import (
	"CookingForest/parser/parser"
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"
	"time"
)

func CreateArchive(recipes []parser.Recipe) (err error) {
	var buf bytes.Buffer
	var js []byte
	var f io.Writer
	timeNow := time.Now().String()

	zipW := zip.NewWriter(&buf)
	if f, err = zipW.Create(fmt.Sprintf("recipes_%s.txt", timeNow)); err != nil {
		return
	}
	if js, err = json.Marshal(recipes); err != nil {
		return
	}
	if _, err = f.Write(js); err != nil {
		return
	}

	var file *os.File
	var img image.Image
	for _, recipe := range recipes {
		if file, err = os.Open(fmt.Sprintf("%s", recipe.ImagePath)); err != nil {
			return
		}
		defer file.Close()
		if img, err = jpeg.Decode(file); err != nil {
			return
		}
		if f, err = zipW.Create(fmt.Sprintf("recipesIMG_%s.jpg", recipe.Name)); err != nil {
			return
		}
		if err = jpeg.Encode(f, img, nil); err != nil {
			return
		}
	}

	if err = zipW.Close(); err != nil {
		return
	}

	var zipFile *os.File
	if zipFile, err = os.OpenFile(fmt.Sprintf("returnedArchives/recipes_%s.zip", timeNow), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		return
	}
	if _, err = zipFile.Write(buf.Bytes()); err != nil {
		return
	}

	return nil
}
