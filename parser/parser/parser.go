package parser

import (
	"CookingForest/parser/request"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"regexp"
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
					match, err = regexp.MatchString("/retsepty/[1-9]{5}", a.Val)
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
	text, _ := ioutil.ReadAll(reader)
	fmt.Println(string(text))
	return Recipe{}, nil
}
