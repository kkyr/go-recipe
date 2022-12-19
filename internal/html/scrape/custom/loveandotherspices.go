package custom

import (
	"fmt"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"

	"github.com/PuerkitoBio/goquery"
)

const LoveAndOtherSpicesHost = "loveandotherspices.com"

// NewLoveAndOtherSpicesScraper returns a new instance of LoveAndOtherSpicesScraper.
func NewLoveAndOtherSpicesScraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &LoveAndOtherSpicesScraper{schema: s}, nil
}

// LoveAndOtherSpicesScraper is a custom recipe scraper for loveandotherspices.com.
type LoveAndOtherSpicesScraper struct {
	schema *schema.RecipeScraper
}

func (m *LoveAndOtherSpicesScraper) Author() (string, bool) {
	return m.schema.Author()
}

func (m *LoveAndOtherSpicesScraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *LoveAndOtherSpicesScraper) CookTime() (time.Duration, bool) {
	return m.schema.CookTime()
}

func (m *LoveAndOtherSpicesScraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *LoveAndOtherSpicesScraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *LoveAndOtherSpicesScraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *LoveAndOtherSpicesScraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

func (m *LoveAndOtherSpicesScraper) Instructions() ([]string, bool) {
	return m.schema.Instructions()
}

func (m *LoveAndOtherSpicesScraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *LoveAndOtherSpicesScraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *LoveAndOtherSpicesScraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *LoveAndOtherSpicesScraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *LoveAndOtherSpicesScraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *LoveAndOtherSpicesScraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *LoveAndOtherSpicesScraper) Yields() (string, bool) {
	return m.schema.Yields()
}
