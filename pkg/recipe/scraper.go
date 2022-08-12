package recipe

import (
	"bytes"
	"fmt"

	recipe "go-recipe"
	"go-recipe/internal/http"
	"go-recipe/internal/scrape/schema"
	urlutil "go-recipe/internal/url"

	"github.com/PuerkitoBio/goquery"
)

type httpClient interface {
	Get(url string) ([]byte, error)
}

var client httpClient = http.NewClient()

// Scraper returns a recipe Scraper that scrapes data from the website at the given url.
func Scraper(url string) (recipe.Scraper, error) {
	body, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("unable to GET url: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("unable to parse HTML document: %w", err)
	}

	host := urlutil.GetHost(url)
	if scraper, ok := hostToScraper[host]; ok {
		return scraper(doc)
	}

	return schema.GetRecipeScraper(doc)
}
