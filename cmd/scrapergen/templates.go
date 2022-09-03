package main

import "text/template"

var scraperTmpl = template.Must(template.New("").Parse(`package custom

import (
	"fmt"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/schema"

	"github.com/PuerkitoBio/goquery"
)

const {{.domain}}Host = "{{.host}}"

// New{{.domain}}Scraper returns a new instance of {{.domain}}Scraper.
func New{{.domain}}Scraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.NewRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &{{.domain}}Scraper{schema: s}, nil
}

// {{.domain}}Scraper is a custom recipe scraper for {{.host}}.
type {{.domain}}Scraper struct {
	schema *schema.RecipeScraper
}

func (m *{{.domain}}Scraper) Author() (string, bool) {
	return m.schema.Author()
}

func (m *{{.domain}}Scraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *{{.domain}}Scraper) CookTime() (time.Duration, bool) {
	return m.schema.CookTime()
}

func (m *{{.domain}}Scraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *{{.domain}}Scraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *{{.domain}}Scraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *{{.domain}}Scraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

func (m *{{.domain}}Scraper) Instructions() ([]string, bool) {
	return m.schema.Instructions()
}

func (m *{{.domain}}Scraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *{{.domain}}Scraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *{{.domain}}Scraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *{{.domain}}Scraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *{{.domain}}Scraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *{{.domain}}Scraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *{{.domain}}Scraper) Yields() (string, bool) {
	return m.schema.Yields()
}
`))

var scraperTestTmpl = template.Must(template.New("").Parse(`package custom_test

import (
	"testing"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/custom"
	"github.com/kkyr/go-recipe/internal/html/scrape/test"

	"github.com/kkyr/assert"
)

func TestNew{{.domain}}Scraper(t *testing.T) {
	doc := test.ReadHTMLFileOrFail(t, custom.{{.domain}}Host)

	scraper, err := custom.New{{.domain}}Scraper(doc)
	assert.New(t).Require().Nil(err)

	scraperTest := test.Scraper{
		Author:        {{.author}},
		Categories:    {{.categories}},
		CookTime:      {{.cookTime}},
		Cuisine:       {{.cuisine}},
		Description:   {{.description}},
		ImageURL:      {{.imageURL}},
		Ingredients:   {{.ingredients}},
		Instructions:  {{.instructions}},
		Language:      {{.language}},
		Name:          {{.name}},
		Nutrition:     {{.nutrition}},
		PrepTime:      {{.prepTime}},
		SuitableDiets: {{.suitableDiets}},
		TotalTime:     {{.totalTime}},
		Yields:        {{.yields}},
	}

	scraperTest.Run(t, scraper)
}`))
