package ld_test

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	ld "github.com/kkyr/go-recipe/internal/html/scrape/schema/json-ld"

	"github.com/PuerkitoBio/goquery"
	"github.com/kkyr/assert"
)

func TestRecipeProcessor_GetRecipeNode(t *testing.T) {
	for _, tc := range []struct {
		name string
		file string
	}{
		{name: "parses graph", file: "json-ld-schema-graph.html"},
		{name: "parses graph with no schema", file: "json-ld-schema-graph-no-schema.html"},
		{name: "parses node", file: "json-ld-schema-node.html"},
		{name: "parses graph with array", file: "json-ld-schema-as-array-type-array.html"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			require := assert.New(t).Require()

			rp := ld.NewRecipeProcessor()

			b, err := os.ReadFile(fmt.Sprintf("testdata/%s", tc.file))
			require.Nil(err)

			doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b))
			require.Nil(err)

			data, err := rp.GetRecipeNode(doc)
			require.Nil(err)

			recipeType := "Recipe"

			if typeData, ok := data["type"].(string); ok {
				require.Field("type").Equal(typeData, recipeType)
			} else if typeData, ok := data["type"].([]interface{}); ok {
				require.Field("type[0]").Equal(typeData[0], recipeType)
			} else {
				t.Fatal("type attribute not in expected shape")
			}

			require.Field("name").NotZero(data["name"])
		})
	}

	t.Run("returns err when no ld+json in doc", func(t *testing.T) {
		require := assert.New(t).Require()

		rp := ld.NewRecipeProcessor()

		const html = `<html><head><script></script></head></html>`
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
		require.Nil(err)

		_, err = rp.GetRecipeNode(doc)
		require.NotNil(err)
	})

	t.Run("returns err when syntax error in json", func(t *testing.T) {
		require := assert.New(t).Require()

		rp := ld.NewRecipeProcessor()

		const html = `<html>
			<head>
				<script type="application/ld+json">
					{
						"@type": "Recipe",
						"
					}
				</script>
			</head>
		</html>
		`
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
		require.Nil(err)

		_, err = rp.GetRecipeNode(doc)
		require.NotNil(err)
	})
}
