package recipe

import (
	"bytes"
	"fmt"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/http"
	"github.com/kkyr/go-recipe/internal/scraper/schema"
	urlutil "github.com/kkyr/go-recipe/internal/url"

	"github.com/PuerkitoBio/goquery"
)

type httpClient interface {
	Get(url string) ([]byte, error)
}

var client httpClient = http.NewClient()

// ScrapeFrom retrieves the source at the provided url and returns a
// Scraper that scrapes recipe data from the retrieved HTML.
func ScrapeFrom(url string) (recipe.Scraper, error) {
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

	scraper, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to get new schema scraper: %w", err)
	}

	return scraper, nil
}
