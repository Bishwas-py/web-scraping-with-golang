package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gocolly/colly"
)

type Images struct {
	Image       string `json:"image"`
	Description string `json:"description"`
}

func main() {
	allImages := make([]Images, 0)
	collector := colly.NewCollector(
		colly.AllowedDomains("img.webmatrices.com"),
	)
	collector.OnHTML(".gallery", func(h *colly.HTMLElement) {
		image_element := h.DOM.Find("a > img").Eq(0)

		image, _ := image_element.Attr("src")
		description, _ := image_element.Attr("alt")

		images := Images{
			Image:       image,
			Description: description,
		}
		allImages = append(allImages, images)
	})

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	collector.Visit("https://img.webmatrices.com")

	writeJSON(allImages)
}

func writeJSON(data []Images) {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Println("Unable to create the JSON file.")
	}
	_ = ioutil.WriteFile("images-data.json", file, 0644)
	fmt.Println("Scraping and Writing successful. Go for Good!")
}
