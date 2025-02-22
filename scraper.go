package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type Recipe struct {
	Name        string
	Ingredients []string
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please provide a URL to scrape.")
		return
	}
	url := args[1]

	collector := colly.NewCollector()

	recipe := Recipe{}

	// Scrape recipe title
	collector.OnHTML("h1", func(e *colly.HTMLElement) {
		recipe.Name = e.Text
	})

	// Scrape ingredients
	collector.OnHTML(".ingredients-list__item", func(e *colly.HTMLElement) {
		recipe.Ingredients = append(recipe.Ingredients, e.Text)
	})

	// Print the collected data when scraping is done
	collector.OnScraped(func(_ *colly.Response) {
		fmt.Println("Recipe Name:", recipe.Name)
		fmt.Println("Ingredients:")
		for _, ingredient := range recipe.Ingredients {
			fmt.Println("-", ingredient)
		}
	})

	collector.Visit(url)
}
