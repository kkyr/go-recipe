package custom

import (
	"fmt"
	stringsUtil "go-recipe/internal/html"
	"strings"
	"time"

	recipe "go-recipe"
	"go-recipe/internal/scrape/schema"

	"github.com/PuerkitoBio/goquery"
)

const ForksOverKnivesHost = "forksoverknives.com"

// NewForksOverKnivesScraper returns a ForksOverKnivesScraper.
func NewForksOverKnivesScraper(doc *goquery.Document) (recipe.Scraper, error) {
	s, err := schema.GetRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("unable to create schema scraper: %w", err)
	}

	return &ForksOverKnivesScraper{doc: doc, schema: s}, nil
}

// ForksOverKnivesScraper is a custom recipe scraper for forksoverknives.com.
type ForksOverKnivesScraper struct {
	doc    *goquery.Document
	schema *schema.RecipeScraper
}

func (m *ForksOverKnivesScraper) Author() (string, bool) {
	ss := m.doc.Find("div > .post-info > p > a").Map(func(_ int, sel *goquery.Selection) string {
		return sel.Text()
	})
	if len(ss) > 0 {
		return stringsUtil.CleanString(ss[0]), true
	}

	return m.schema.Author()
}

func (m *ForksOverKnivesScraper) Categories() ([]string, bool) {
	return m.schema.Categories()
}

func (m *ForksOverKnivesScraper) CookTime() (time.Duration, bool) {
	return m.schema.CookTime()
}

func (m *ForksOverKnivesScraper) Cuisine() ([]string, bool) {
	return m.schema.Cuisine()
}

func (m *ForksOverKnivesScraper) Description() (string, bool) {
	return m.schema.Description()
}

func (m *ForksOverKnivesScraper) ImageURL() (string, bool) {
	return m.schema.ImageURL()
}

func (m *ForksOverKnivesScraper) Ingredients() ([]string, bool) {
	return m.schema.Ingredients()
}

func (m *ForksOverKnivesScraper) Instructions() ([]string, bool) {
	return m.schema.Instructions()
}

func (m *ForksOverKnivesScraper) Language() (string, bool) {
	return m.schema.Language()
}

func (m *ForksOverKnivesScraper) Name() (string, bool) {
	return m.schema.Name()
}

func (m *ForksOverKnivesScraper) Nutrition() (recipe.Nutrition, bool) {
	return m.schema.Nutrition()
}

func (m *ForksOverKnivesScraper) PrepTime() (time.Duration, bool) {
	return m.schema.PrepTime()
}

func (m *ForksOverKnivesScraper) SuitableDiets() ([]recipe.Diet, bool) {
	return m.schema.SuitableDiets()
}

func (m *ForksOverKnivesScraper) TotalTime() (time.Duration, bool) {
	return m.schema.TotalTime()
}

func (m *ForksOverKnivesScraper) Yields() (string, bool) {
	ss := m.doc.Find("li > .icon-serving").Map(func(_ int, sel *goquery.Selection) string {
		return sel.Siblings().First().Text()
	})
	if len(ss) > 0 {
		return strings.TrimPrefix(stringsUtil.CleanString(ss[0]), "Makes "), true
	}

	return m.schema.Yields()
}
