package recipe

import (
	"fmt"
	"testing"

	"github.com/kkyr/go-recipe/internal/scrape/custom"
	"github.com/kkyr/go-recipe/internal/scrape/schema"
	"github.com/kkyr/go-recipe/internal/scrape/test"

	"github.com/kkyr/assert"
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

func TestScrapeFrom(t *testing.T) {
	assert := assert.New(t)

	client = &mockHTTPClient{
		GetFunc: func(url string) ([]byte, error) {
			return []byte(htmlSchemaRecipe), nil
		},
	}

	t.Run("using custom scraper", func(t *testing.T) {
		scraper, err := ScrapeFrom(custom.MinimalistBakerHost)
		assert.Require().Nil(err)

		if _, ok := scraper.(*custom.MinimalistBakerScraper); !ok {
			t.Errorf("want type *custom.MinimalistBakerScraper, got %T", scraper)
		}

		const want = "Lemon Cake"
		got, ok := scraper.Name()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("using default scraper", func(t *testing.T) {
		scraper, err := ScrapeFrom("")
		assert.Require().Nil(err)

		if _, ok := scraper.(*schema.RecipeScraper); !ok {
			t.Errorf("want type *schema.RecipeScraper, got %T", scraper)
		}

		const want = "Lemon Cake"
		got, ok := scraper.Name()
		test.Verify(t, true, ok, want, got)
	})
}

func TestScrapeFrom_Err(t *testing.T) {
	assert := assert.New(t)

	t.Run("using bad document", func(t *testing.T) {
		client = &mockHTTPClient{
			GetFunc: func(url string) ([]byte, error) {
				return []byte("not an html document"), nil
			},
		}

		_, err := ScrapeFrom("")
		assert.NotNil(err)
	})

	t.Run("bad request", func(t *testing.T) {
		boom := fmt.Errorf("boom")

		client = &mockHTTPClient{
			GetFunc: func(url string) ([]byte, error) {
				return nil, boom
			},
		}

		_, err := ScrapeFrom("")
		assert.ErrorIs(err, boom)
	})
}
