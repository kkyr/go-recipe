package schema

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html/scrape/test"
)

func TestRecipeScraper_Author(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		const want = "krishnamurti"
		scraper := newRecipeScraper(map[string]any{
			"author": map[string]any{
				"name": want,
			},
		})

		got, ok := scraper.Author()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("empty", func(t *testing.T) {
		scraper := newRecipeScraper(map[string]any{})

		got, ok := scraper.Author()
		test.Verify(t, false, ok, "", got)
	})
}

func TestRecipeScraper_Categories(t *testing.T) {
	t.Run("string-slice", func(t *testing.T) {
		want := []string{"Appetizer", "Dessert"}
		scraper := newRecipeScraper(map[string]any{
			"recipeCategory": strings.Join(want, ","),
		})

		got, ok := scraper.Categories()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("slice", func(t *testing.T) {
		want := []string{"Dessert", "Appetizer"}
		scraper := newRecipeScraper(map[string]any{
			"recipeCategory": []any{
				want[0], want[1],
			},
		})

		got, ok := scraper.Categories()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("string", func(t *testing.T) {
		want := []string{"aperitif"}
		scraper := newRecipeScraper(map[string]any{
			"recipeCategory": want[0],
		})

		got, ok := scraper.Categories()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("empty", func(t *testing.T) {
		scraper := newRecipeScraper(map[string]any{})

		got, ok := scraper.Categories()
		test.Verify(t, false, ok, "", got)
	})
}

func TestRecipeScraper_CookTime(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		const want = 5 * time.Minute
		scraper := newRecipeScraper(map[string]any{
			"cookTime": "PT5M",
		})

		got, ok := scraper.CookTime()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("empty", func(t *testing.T) {
		scraper := newRecipeScraper(map[string]any{})

		got, ok := scraper.CookTime()
		test.Verify(t, false, ok, "", got)
	})
}

func TestRecipeScraper_Cuisine(t *testing.T) {
	t.Run("slice", func(t *testing.T) {
		want := []string{"French", "Italian", "Indian"}
		scraper := newRecipeScraper(map[string]any{
			"recipeCuisine": []any{
				want[0], want[1], want[2],
			},
		})

		got, ok := scraper.Cuisine()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("empty", func(t *testing.T) {
		scraper := newRecipeScraper(map[string]any{})

		got, ok := scraper.Cuisine()
		test.Verify(t, false, ok, "", got)
	})
}

func TestRecipeScraper_Description(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		const want = "best you've ever had"
		scraper := newRecipeScraper(map[string]any{
			"description": want,
		})

		got, ok := scraper.Description()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("empty", func(t *testing.T) {
		scraper := newRecipeScraper(map[string]any{})

		got, ok := scraper.Description()
		test.Verify(t, false, ok, "", got)
	})
}

func TestRecipeScraper_ImageURL(t *testing.T) {
	t.Run("inner-string", func(t *testing.T) {
		const want = "https://www.example.com/image.jpg"
		scraper := newRecipeScraper(map[string]any{
			"image": map[string]any{
				"url": want,
			},
		})

		got, ok := scraper.ImageURL()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("string", func(t *testing.T) {
		const want = "https://www.example.com/image.jpg"
		scraper := newRecipeScraper(map[string]any{
			"image": want,
		})

		got, ok := scraper.ImageURL()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("slice", func(t *testing.T) {
		const want = "https://www.example.com/image.jpg"
		scraper := newRecipeScraper(map[string]any{
			"image": []any{
				want,
			},
		})

		got, ok := scraper.ImageURL()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("empty", func(t *testing.T) {
		scraper := newRecipeScraper(map[string]any{})

		got, ok := scraper.ImageURL()
		test.Verify(t, false, ok, "", got)
	})
}

func TestRecipeScraper_Ingredients(t *testing.T) {
	t.Run("string-recipeIngredient", func(t *testing.T) {
		want := []string{"ingredient 1"}

		scraper := newRecipeScraper(map[string]any{
			"recipeIngredient": want[0],
		})

		got, ok := scraper.Ingredients()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("slice-recipeIngredient", func(t *testing.T) {
		want := []string{
			"ingredient 1",
			"ingredient 2",
			"ingredient 3",
		}
		scraper := newRecipeScraper(map[string]any{
			"recipeIngredient": []any{
				want[0],
				want[1],
				want[2],
			},
		})

		got, ok := scraper.Ingredients()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("string-ingredients", func(t *testing.T) {
		want := []string{"ingredient 1"}

		scraper := newRecipeScraper(map[string]any{
			"ingredients": want[0],
		})

		got, ok := scraper.Ingredients()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("slice-ingredients", func(t *testing.T) {
		want := []string{
			"ingredient 1",
			"ingredient 2",
			"ingredient 3",
		}
		scraper := newRecipeScraper(map[string]any{
			"ingredients": []any{
				want[0],
				want[1],
				want[2],
			},
		})

		got, ok := scraper.Ingredients()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("empty", func(t *testing.T) {
		scraper := newRecipeScraper(map[string]any{})

		got, ok := scraper.Ingredients()
		test.Verify(t, false, ok, "", got)
	})
}

func TestRecipeScraper_Instructions(t *testing.T) {
	t.Run("slice", func(t *testing.T) {
		want := []string{
			"instruction 1",
			"instruction 2",
			"instruction 3",
		}
		scraper := newRecipeScraper(map[string]any{
			"recipeInstructions": []any{
				want[0],
				map[string]any{
					"type": "HowToStep",
					"text": want[1],
				},
				map[string]any{
					"itemListElement": []any{
						map[string]any{
							"type": "HowToStep",
							"text": want[2],
						},
					},
				},
			},
		})

		got, ok := scraper.Instructions()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("empty", func(t *testing.T) {
		scraper := newRecipeScraper(map[string]any{})

		got, ok := scraper.Instructions()
		test.Verify(t, false, ok, "", got)
	})
}

func TestRecipeScraper_Language(t *testing.T) {
	t.Run("string-inLanguage", func(t *testing.T) {
		const want = "en-US"
		scraper := newRecipeScraper(map[string]any{
			"inLanguage": want,
		})

		got, ok := scraper.Language()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("string-language", func(t *testing.T) {
		const want = "en-US"
		scraper := newRecipeScraper(map[string]any{
			"language": want,
		})

		got, ok := scraper.Language()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("empty", func(t *testing.T) {
		scraper := newRecipeScraper(map[string]any{})

		got, ok := scraper.Language()
		test.Verify(t, false, ok, "", got)
	})
}

func TestRecipeScraper_Name(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		const want = "fried milk"
		scraper := newRecipeScraper(map[string]any{
			"name": want,
		})

		got, ok := scraper.Name()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("empty", func(t *testing.T) {
		scraper := newRecipeScraper(map[string]any{})

		got, ok := scraper.Name()
		test.Verify(t, false, ok, "", got)
	})
}

func TestRecipeScraper_Nutrition(t *testing.T) {
	t.Run("map", func(t *testing.T) {
		want := recipe.Nutrition{Calories: 320}
		scraper := newRecipeScraper(map[string]any{
			"nutrition": map[string]any{
				"calories": fmt.Sprintf("%f kcal", want.Calories),
			},
		})

		got, ok := scraper.Nutrition()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("empty", func(t *testing.T) {
		scraper := newRecipeScraper(map[string]any{})

		got, ok := scraper.Nutrition()
		test.Verify(t, false, ok, "", got)
	})
}

func TestRecipeScraper_PrepTime(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		const want = 1 * time.Hour
		scraper := newRecipeScraper(map[string]any{
			"prepTime": "PT1H",
		})

		got, ok := scraper.PrepTime()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("empty", func(t *testing.T) {
		scraper := newRecipeScraper(map[string]any{})

		got, ok := scraper.PrepTime()
		test.Verify(t, false, ok, "", got)
	})
}

func TestRecipeScraper_SuitableDiets(t *testing.T) {
	t.Run("map", func(t *testing.T) {
		want := []recipe.Diet{recipe.VeganDiet, recipe.VegetarianDiet}
		scraper := newRecipeScraper(map[string]any{
			"suitableForDiet": fmt.Sprintf("%s, %s", want[0], want[1]),
		})

		got, ok := scraper.SuitableDiets()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("empty", func(t *testing.T) {
		scraper := newRecipeScraper(map[string]any{})

		got, ok := scraper.SuitableDiets()
		test.Verify(t, false, ok, "", got)
	})
}

func TestRecipeScraper_TotalTime(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		const want = 62 * time.Minute
		scraper := newRecipeScraper(map[string]any{
			"totalTime": "PT1H2M",
		})

		got, ok := scraper.TotalTime()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("empty", func(t *testing.T) {
		scraper := newRecipeScraper(map[string]any{})

		got, ok := scraper.TotalTime()
		test.Verify(t, false, ok, "", got)
	})
}

func TestRecipeScraper_Yields(t *testing.T) {
	t.Run("string-yield", func(t *testing.T) {
		const want = "5 tons"
		scraper := newRecipeScraper(map[string]any{
			"yield": want,
		})

		got, ok := scraper.Yields()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("slice-yield", func(t *testing.T) {
		const want = "5 tons"
		scraper := newRecipeScraper(map[string]any{
			"yield": []any{
				want,
			},
		})

		got, ok := scraper.Yields()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("string-recipeYield", func(t *testing.T) {
		const want = "5 tons"
		scraper := newRecipeScraper(map[string]any{
			"recipeYield": want,
		})

		got, ok := scraper.Yields()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("slice-recipeYield", func(t *testing.T) {
		const want = "5 tons"
		scraper := newRecipeScraper(map[string]any{
			"recipeYield": []any{
				want,
			},
		})

		got, ok := scraper.Yields()
		test.Verify(t, true, ok, want, got)
	})

	t.Run("empty", func(t *testing.T) {
		scraper := newRecipeScraper(map[string]any{})

		got, ok := scraper.Yields()
		test.Verify(t, false, ok, "", got)
	})
}
