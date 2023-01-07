package recipe

import (
	"bytes"
	"fmt"
	"io"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"
	"github.com/kkyr/go-recipe/internal/http"
	"github.com/kkyr/go-recipe/internal/url"

	"github.com/PuerkitoBio/goquery"
)

type httpClient interface {
	Get(url string) ([]byte, error)
}

var client httpClient = http.NewClient()

// ScrapeURL retrieves the source at the provided url and returns a
// Scraper that scrapes recipe data from the retrieved HTML.
func ScrapeURL(urlStr string) (recipe.Scraper, error) {
	body, err := client.Get(urlStr)
	if err != nil {
		return nil, fmt.Errorf("unable to GET url: %w", err)
	}

	return ScrapeHTML(urlStr, bytes.NewReader(body))
}

// ScrapeHTML returns a Scraper that scrapes recipe data from the provided HTML.
// the urlStr is used to determine if a specific scraper should be used, otherwise
// a generic scraper is used.
func ScrapeHTML(urlStr string, body io.Reader) (recipe.Scraper, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, fmt.Errorf("unable to parse HTML document: %w", err)
	}

	host := url.GetHost(urlStr)
	if scraper, ok := hostToScraper[host]; ok {
		return scraper(doc)
	}

	scraper, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to get new schema scraper: %w", err)
	}

	return scraper, nil
}
