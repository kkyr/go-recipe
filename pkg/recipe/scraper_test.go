package recipe

import (
	"errors"
	"fmt"
	"testing"

	"go-recipe/internal/scrape/custom"
	"go-recipe/internal/scrape/schema"
	"go-recipe/internal/scrape/test"
)

type mockHTTPClient struct {
	GetFunc func(url string) ([]byte, error)
}

func (m *mockHTTPClient) Get(url string) ([]byte, error) {
	return m.GetFunc(url)
}

const htmlSchemaRecipe = `<html>
    <head>
        <script type="application/ld+json">
            {
                "@context": "https://schema.org",
                "@type": "Recipe",
				"name": "Lemon Cake"
            }
        </script>
    </head>
</html>`

func TestScraper(t *testing.T) {
	client = &mockHTTPClient{
		GetFunc: func(url string) ([]byte, error) {
			return []byte(htmlSchemaRecipe), nil
		},
	}

	t.Run("using custom scraper", func(t *testing.T) {
		scraper, err := Scraper(custom.MinimalistBakerHost)
		if err != nil {
			t.Fatalf("unexpected err when getting scraper: %v", err)
		}

		if _, ok := scraper.(*custom.MinimalistBakerScraper); !ok {
			t.Errorf("want type *custom.MinimalistBakerScraper, got %T", scraper)
		}

		const want = "Lemon Cake"
		got, ok := scraper.Name()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("using default scraper", func(t *testing.T) {
		scraper, err := Scraper("")
		if err != nil {
			t.Fatalf("unexpected err when getting scraper: %v", err)
		}

		if _, ok := scraper.(*schema.RecipeScraper); !ok {
			t.Errorf("want type *schema.RecipeScraper, got %T", scraper)
		}

		const want = "Lemon Cake"
		got, ok := scraper.Name()
		test.Verify(t, true, ok, want, got)
	})
}

func TestScraper_Err(t *testing.T) {
	t.Run("using bad document", func(t *testing.T) {
		client = &mockHTTPClient{
			GetFunc: func(url string) ([]byte, error) {
				return []byte("not an html document"), nil
			},
		}

		_, err := Scraper("")
		if err == nil {
			t.Fatalf("want err, got nil")
		}
	})

	t.Run("bad request", func(t *testing.T) {
		boom := fmt.Errorf("boom")

		client = &mockHTTPClient{
			GetFunc: func(url string) ([]byte, error) {
				return nil, boom
			},
		}

		_, err := Scraper("")
		if !errors.Is(err, boom) {
			t.Fatalf("want %v, got %v", boom, err)
		}
	})
}
