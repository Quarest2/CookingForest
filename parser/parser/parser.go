package parser

import (
	"CookingForest/parser/request"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Parse(bodyStr io.Reader, counter int) ([]Recipe, error) {
	var doc *html.Node
	var err error
	var recipes = make([]string, 0)

	if doc, err = html.Parse(bodyStr); err != nil {
		return nil, err
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					var match bool
					match, err = regexp.MatchString("/retsepty/[0-9]{5}", a.Val)
					if err == nil && match {
						match, err = regexp.MatchString("#comments_anchor", a.Val)
						if err == nil && !match {
							recipes = append(recipes, a.Val)
						}
					}
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	recipes = DedupeStrings(recipes)
	var reader io.Reader
	var recipe Recipe
	var result []Recipe
	if len(recipes) < counter {
		for _, r := range recipes {
			if reader, err = request.GetBody("https://www.edimdoma.ru" + r); err == nil {
				if recipe, err = GetOneRecipe(reader); err == nil {
					result = append(result, recipe)
				}
			}
		}
	} else {
		for i := 0; i < counter; i++ {
			if reader, err = request.GetBody("https://www.edimdoma.ru" + recipes[i]); err == nil {
				if recipe, err = GetOneRecipe(reader); err == nil {
					result = append(result, recipe)
				}
			}
		}
	}
	return result, nil
}

func GetOneRecipe(reader io.Reader) (Recipe, error) {
	var doc *html.Node
	var err error
	var recipe Recipe

	if doc, err = html.Parse(reader); err != nil {
		return Recipe{}, err
	}

	var strTimeAndPersons []string
	var recipeName string
	var strUrlImg string
	var steps = make([]string, 0)
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" {
			// Ищем время приготовления и кол-во персон
			if len(n.Attr) > 0 && n.Attr[0].Val == "entry-stats__value" {
				secondChild := n.FirstChild.FirstChild
				strTimeAndPersons = append(strTimeAndPersons, secondChild.Data)
			}
			// Ищем шаги приготовления
			if len(n.Attr) > 0 && n.Attr[0].Val == "step_hint" {
				if n.FirstChild != nil {
					steps = append(steps, n.FirstChild.Data)
				}
			}
		}
		// Ищем название рецепта
		if n.Type == html.ElementNode && n.Data == "h1" {
			if len(n.Attr) > 0 && n.Attr[0].Val == "recipe-header__name" {
				recipeName = strings.TrimSpace(n.FirstChild.Data)
			}
		}
		// Ищем изображение
		if n.Type == html.ElementNode && n.Data == "img" {
			var match bool
			if len(n.Attr) > 0 {
				if match, err = regexp.MatchString("/data/recipes/[0-9]{4}/[0-9]{4}", n.Attr[0].Val); err == nil {
					if match {
						strUrlImg = "https://www.edimdoma.ru" + n.Attr[0].Val
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	recipe.Name = recipeName

	var persons int
	if len(strTimeAndPersons) == 2 {
		if persons, err = strconv.Atoi(strTimeAndPersons[1]); err != nil {
			return Recipe{}, err
		}
		recipe.Persons = persons
		recipe.Time = strTimeAndPersons[0] + " minutes"
	}
	if len(strTimeAndPersons) == 3 {
		if persons, err = strconv.Atoi(strTimeAndPersons[2]); err != nil {
			return Recipe{}, err
		}
		recipe.Persons = persons
		recipe.Time = strTimeAndPersons[1] + " minutes"
	}

	var response *http.Response
	var file *os.File
	if response, err = http.Get(strUrlImg); err == nil {
		defer response.Body.Close()
		if file, err = os.Create(fmt.Sprintf("images/%s.jpg", recipeName)); err == nil {
			defer file.Close()
			_, _ = io.Copy(file, response.Body)
			recipe.ImagePath = fmt.Sprintf("images/%s.jpg", recipeName)
		}
	}

	recipe.Steps = steps

	return recipe, nil
}
