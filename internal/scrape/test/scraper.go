package test

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"

	recipe "go-recipe"

	"github.com/google/go-cmp/cmp"
)

// Scraper is a struct that helps test packages to test values returned by a recipe.Scraper.
// Clients should fill in the values of the struct that they expect to be returned by the
// recipe.Scraper and then call Run(). Defined values will be validated for an exact match
// and zero values will be validated that nothing is returned.
type Scraper struct {
	Author        string
	Categories    []string
	CookTime      time.Duration
	Cuisine       []string
	Description   string
	ImageURL      string
	Ingredients   []string
	Instructions  []string
	Language      string
	Name          string
	Nutrition     recipe.Nutrition
	PrepTime      time.Duration
	SuitableDiets []recipe.Diet
	TotalTime     time.Duration
	Yields        string
}

// Run runs the test and compares struct values against those returned by the scraper.
// Any failures are reported using t.Error().
func (s *Scraper) Run(t *testing.T, scraper recipe.Scraper) {
	t.Helper()

	t.Run("author", func(t *testing.T) {
		wantVal, wantOK := s.Author, !reflect.ValueOf(s.Author).IsZero()
		got, ok := scraper.Author()
		Verify(t, wantOK, ok, wantVal, got)
	})

	t.Run("categories", func(t *testing.T) {
		wantVal, wantOK := s.Categories, !reflect.ValueOf(s.Categories).IsZero()
		got, ok := scraper.Categories()
		Verify(t, wantOK, ok, wantVal, got)
	})

	t.Run("cookTime", func(t *testing.T) {
		wantVal, wantOK := s.CookTime, !reflect.ValueOf(s.CookTime).IsZero()
		got, ok := scraper.CookTime()
		Verify(t, wantOK, ok, wantVal, got)
	})

	t.Run("cuisine", func(t *testing.T) {
		wantVal, wantOK := s.Cuisine, !reflect.ValueOf(s.Cuisine).IsZero()
		got, ok := scraper.Cuisine()
		Verify(t, wantOK, ok, wantVal, got)
	})

	t.Run("description", func(t *testing.T) {
		wantVal, wantOK := s.Description, !reflect.ValueOf(s.Description).IsZero()
		got, ok := scraper.Description()
		Verify(t, wantOK, ok, wantVal, got)
	})

	t.Run("imageURL", func(t *testing.T) {
		wantVal, wantOK := s.ImageURL, !reflect.ValueOf(s.ImageURL).IsZero()
		got, ok := scraper.ImageURL()
		Verify(t, wantOK, ok, wantVal, got)
	})

	t.Run("ingredients", func(t *testing.T) {
		wantVal, wantOK := s.Ingredients, !reflect.ValueOf(s.Ingredients).IsZero()
		got, ok := scraper.Ingredients()
		Verify(t, wantOK, ok, wantVal, got)
	})

	t.Run("instructions", func(t *testing.T) {
		wantVal, wantOK := s.Instructions, !reflect.ValueOf(s.Instructions).IsZero()
		got, ok := scraper.Instructions()
		Verify(t, wantOK, ok, wantVal, got)
	})

	t.Run("language", func(t *testing.T) {
		wantVal, wantOK := s.Language, !reflect.ValueOf(s.Language).IsZero()
		got, ok := scraper.Language()
		Verify(t, wantOK, ok, wantVal, got)
	})

	t.Run("name", func(t *testing.T) {
		wantVal, wantOK := s.Name, !reflect.ValueOf(s.Name).IsZero()
		got, ok := scraper.Name()
		Verify(t, wantOK, ok, wantVal, got)
	})

	t.Run("nutrition", func(t *testing.T) {
		wantVal, wantOK := s.Nutrition, !reflect.ValueOf(s.Nutrition).IsZero()
		got, ok := scraper.Nutrition()
		Verify(t, wantOK, ok, wantVal, got)
	})

	t.Run("prepTime", func(t *testing.T) {
		wantVal, wantOK := s.PrepTime, !reflect.ValueOf(s.PrepTime).IsZero()
		got, ok := scraper.PrepTime()
		Verify(t, wantOK, ok, wantVal, got)
	})

	t.Run("suitableDiets", func(t *testing.T) {
		wantVal, wantOK := s.SuitableDiets, !reflect.ValueOf(s.SuitableDiets).IsZero()
		got, ok := scraper.SuitableDiets()
		Verify(t, wantOK, ok, wantVal, got)
	})

	t.Run("totalTime", func(t *testing.T) {
		wantVal, wantOK := s.TotalTime, !reflect.ValueOf(s.TotalTime).IsZero()
		got, ok := scraper.TotalTime()
		Verify(t, wantOK, ok, wantVal, got)
	})

	t.Run("yields", func(t *testing.T) {
		wantVal, wantOK := s.Yields, !reflect.ValueOf(s.Yields).IsZero()
		got, ok := scraper.Yields()
		Verify(t, wantOK, ok, wantVal, got)
	})
}

// Verify checks whether wantOK matches with gotOK and reports failures with t.Errorf().
// If got OK is set to true then it also checks that wantValue matches with gotValue,
// and similarly reports failures with t.Errorf().
func Verify(t *testing.T, wantOK bool, gotOK bool, wantVal any, gotVal any) {
	t.Helper()

	if wantOK != gotOK {
		t.Errorf("want ok = %t, got %t", wantOK, gotOK)
	}

	if gotOK {
		if diff := cmp.Diff(wantVal, gotVal); diff != "" {
			t.Errorf("(-want +got):\n%v", diff)
		}
	}
}

// ReadHTMLFileOrFail attempts to read an HTML file from the testdata directory and
// then load it into a Document. The filename is formed by adding a .html suffix
// name and then appending it to the testdata directory to form the file path.
// Any failures are reported with t.Fatal().
func ReadHTMLFileOrFail(t *testing.T, host string) *goquery.Document {
	t.Helper()

	file, err := os.Open(fmt.Sprintf("testdata/%s.html", host))
	if err != nil {
		t.Fatalf("unexpected err while opening file: %v", err)
	}
	defer file.Close()

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		t.Fatalf("unexpected err while reading file: %v", err)
	}

	return doc
}
