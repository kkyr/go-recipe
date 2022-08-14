package custom

import (
	"fmt"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/scraper/schema"

	"github.com/PuerkitoBio/goquery"
)

const MinimalistBakerHost = "minimalistbaker.com"

// NewMinimalistBakerScraper returns a MinimalistBakerScraper.
func NewMinimalistBakerScraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.GetRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &MinimalistBakerScraper{schema: s}, nil
}

// MinimalistBakerScraper is a custom recipe scraper for minimalistbaker.com.
type MinimalistBakerScraper struct {
	schema *schema.RecipeScraper
}

func (m *MinimalistBakerScraper) Author() (string, bool) {
	return m.schema.Author()
}

func (m *MinimalistBakerScraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *MinimalistBakerScraper) CookTime() (time.Duration, bool) {
	return m.schema.CookTime()
}

func (m *MinimalistBakerScraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *MinimalistBakerScraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *MinimalistBakerScraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *MinimalistBakerScraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

func (m *MinimalistBakerScraper) Instructions() ([]string, bool) {
	return m.schema.Instructions()
}

func (m *MinimalistBakerScraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *MinimalistBakerScraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *MinimalistBakerScraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *MinimalistBakerScraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *MinimalistBakerScraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *MinimalistBakerScraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *MinimalistBakerScraper) Yields() (string, bool) {
	return m.schema.Yields()
}
