package custom

import (
	"fmt"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"

	"github.com/PuerkitoBio/goquery"
)

const AllRecipesHost = "allrecipes.com"

// NewAllRecipesScraper returns a new instance of AllRecipesScraper.
func NewAllRecipesScraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &AllRecipesScraper{schema: s}, nil
}

// AllRecipesScraper is a custom recipe scraper for allrecipes.com.
type AllRecipesScraper struct {
	schema *schema.RecipeScraper
}

func (m *AllRecipesScraper) Author() (string, bool) {
	return m.schema.Author()
}

func (m *AllRecipesScraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *AllRecipesScraper) CookTime() (time.Duration, bool) {
	return m.schema.CookTime()
}

func (m *AllRecipesScraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *AllRecipesScraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *AllRecipesScraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *AllRecipesScraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

func (m *AllRecipesScraper) Instructions() ([]string, bool) {
	return m.schema.Instructions()
}

func (m *AllRecipesScraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *AllRecipesScraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *AllRecipesScraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *AllRecipesScraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *AllRecipesScraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *AllRecipesScraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *AllRecipesScraper) Yields() (string, bool) {
	return m.schema.Yields()
}
