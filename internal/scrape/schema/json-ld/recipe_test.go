package ld_test

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	ld "go-recipe/internal/scrape/schema/json-ld"

	"github.com/PuerkitoBio/goquery"
)

func TestRecipeProcessor_GetRecipeNode(t *testing.T) {
	rp := ld.NewRecipeProcessor()

	for _, tc := range []struct {
		name string
		file string
	}{
		{name: "parses graph", file: "json-ld-schema-graph.html"},
		{name: "parses node", file: "json-ld-schema-node.html"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			b, err := os.ReadFile(fmt.Sprintf("testdata/%s", tc.file))
			if err != nil {
				t.Fatalf("unexpected err while reading file: %v", err)
			}

			doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b))
			if err != nil {
				t.Fatalf("unexpected err while creating doc: %v", err)
			}

			data, err := rp.GetRecipeNode(doc)
			if err != nil {
				t.Fatalf("unexpected err while parsing: %v", err)
			}

			if data["type"] != "Recipe" {
				t.Errorf("want type Recipe, got %v", data["type"])
			}
		})
	}

	t.Run("returns err when no ld+json in doc", func(t *testing.T) {
		const html = `<html><head><script></script></head></html>`
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
		if err != nil {
			t.Fatalf("unexpected err while creating doc: %v", err)
		}

		_, err = rp.GetRecipeNode(doc)
		if err == nil {
			t.Fatalf("expected err, got nil")
		}
	})
}
