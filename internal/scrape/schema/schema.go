package schema

import (
	"fmt"
	"strings"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/html"
	ld "github.com/kkyr/go-recipe/internal/scrape/schema/json-ld"

	"github.com/PuerkitoBio/goquery"
	"github.com/senseyeio/duration"
)

var ldProcessor = ld.NewRecipeProcessor()

// GetRecipeScraper returns a RecipeScraper that scrapes recipe data from an HTML
// document containing a Schema.org Recipe.
func GetRecipeScraper(doc *goquery.Document) (*RecipeScraper, error) {
	node, err := ldProcessor.GetRecipeNode(doc)
	if err != nil {
		return nil, fmt.Errorf("could not get recipe root from ld+json document: %w", err)
	}

	return newRecipeScraper(node), nil
}

func newRecipeScraper(data map[string]any) *RecipeScraper {
	return &RecipeScraper{root: data}
}

var _ recipe.Scraper = (*RecipeScraper)(nil)

// RecipeScraper is a recipe scraper.
type RecipeScraper struct {
	root map[string]any
}

// Author is the author of the recipe.
func (r *RecipeScraper) Author() (string, bool) {
	if node, ok := r.root["author"].(map[string]any); ok {
		return getStringValue(node, "name")
	}

	return "", false
}

// Categories are the categories of the recipe, e.g. appetizer, entrée, etc.
func (r *RecipeScraper) Categories() ([]string, bool) {
	if s, ok := r.root["recipeCategory"].(string); ok {
		ss := strings.Split(s, ",")
		for i := range ss {
			ss[i] = html.CleanString(ss[i])
		}

		return ss, true
	}

	return getSliceValue(r.root, "recipeCategory")
}

// CookTime is the time it takes to actually cook the dish.
func (r *RecipeScraper) CookTime() (time.Duration, bool) {
	return getDurationValue(r.root, "cookTime")
}

// Cuisine is the cuisine of the recipe, e.g. mexican-inspired, french, etc.
func (r *RecipeScraper) Cuisine() ([]string, bool) {
	return getSliceValue(r.root, "recipeCuisine")
}

// Description is the description of the recipe.
func (r *RecipeScraper) Description() (string, bool) {
	return getStringValue(r.root, "description")
}

// ImageURL is a URL to an image of the dish.
func (r *RecipeScraper) ImageURL() (string, bool) {
	node := r.root["image"]

	if s, ok := node.(string); ok {
		return html.CleanString(s), true
	}

	if m, ok := node.(map[string]any); ok {
		return getStringValue(m, "url")
	}

	if ss, ok := node.([]any); ok {
		if len(ss) > 0 {
			if s, ok := ss[0].(string); ok {
				return html.CleanString(s), true
			}
		}
	}

	return "", false
}

// Ingredients are all the ingredients used in the recipe.
func (r *RecipeScraper) Ingredients() ([]string, bool) {
	if ss, ok := getSliceValue(r.root, "recipeIngredient"); ok {
		return ss, true
	}

	return getSliceValue(r.root, "ingredients")
}

// Instructions are all the steps in making the recipe.
func (r *RecipeScraper) Instructions() ([]string, bool) {
	if nodes, ok := r.root["recipeInstructions"].([]any); ok {
		instructions := getInstructions(nodes)
		if len(instructions) > 0 {
			return instructions, true
		}
	}

	return nil, false
}

// Language is the language used in the recipe expressed in IETF BCP 47 standard.
func (r *RecipeScraper) Language() (string, bool) {
	for _, key := range []string{"inLanguage", "language"} {
		if s, ok := getStringValue(r.root, key); ok {
			return s, true
		}
	}

	return "", false
}

// Name is the name of the recipe.
func (r *RecipeScraper) Name() (string, bool) {
	return getStringValue(r.root, "name")
}

// Nutrition is nutritional information about the dish.
func (r *RecipeScraper) Nutrition() (recipe.Nutrition, bool) {
	if m, ok := r.root["nutrition"].(map[string]any); ok {
		return ParseNutritionalInformation(m), true
	}

	return recipe.Nutrition{}, false
}

// PrepTime is the length of time it takes to prepare the items to be used in the instructions.
func (r *RecipeScraper) PrepTime() (time.Duration, bool) {
	return getDurationValue(r.root, "prepTime")
}

// SuitableDiets indicates dietary restrictions or guidelines for which the recipe is suitable.
func (r *RecipeScraper) SuitableDiets() ([]recipe.Diet, bool) {
	if v, ok := getStringValue(r.root, "suitableForDiet"); ok {
		ss := strings.Split(v, ",")
		ret := make([]recipe.Diet, 0, len(ss))

		for _, s := range ss {
			if diet := ParseDiet(s); diet != recipe.UnknownDiet {
				ret = append(ret, diet)
			}
		}

		return ret, true
	}

	return nil, false
}

// TotalTime is the total time required to perform the instructions (including the prep time).
func (r *RecipeScraper) TotalTime() (time.Duration, bool) {
	return getDurationValue(r.root, "totalTime")
}

// Yields is the quantity that results from the recipe.
func (r *RecipeScraper) Yields() (string, bool) {
	yieldData := r.root["recipeYield"]
	if yieldData == nil {
		yieldData = r.root["yield"]
	}

	if s, ok := yieldData.(string); ok {
		return html.CleanString(s), true
	}

	if ss, ok := yieldData.([]any); ok {
		if len(ss) > 0 {
			// the last element seems to be the most descriptive
			if s, ok := ss[len(ss)-1].(string); ok {
				return html.CleanString(s), true
			}
		}
	}

	return "", false
}

func getInstructions(nodes []any) []string {
	var ret []string

	for _, instr := range nodes {
		if s, ok := instr.(string); ok {
			ret = append(ret, html.CleanString(s))
		}

		if m, ok := instr.(map[string]any); ok {
			if m["type"] == "HowToStep" {
				if step, ok := m["text"].(string); ok {
					ret = append(ret, html.CleanString(step))
				}
			} else if items, ok := m["itemListElement"].([]any); ok {
				ret = append(ret, getInstructions(items)...)
			}
		}
	}

	return ret
}

func getStringValue(node map[string]any, key string) (string, bool) {
	s, ok := node[key].(string)
	if !ok {
		return "", false
	}

	return html.CleanString(s), true
}

func getSliceValue(node map[string]any, key string) ([]string, bool) {
	if s, ok := getStringValue(node, key); ok {
		return []string{s}, true
	}

	if ss, ok := node[key].([]any); ok {
		ret := make([]string, 0, len(ss))

		for _, v := range ss {
			if s, ok := v.(string); ok {
				ret = append(ret, html.CleanString(s))
			}
		}

		return ret, true
	}

	return nil, false
}

func getDurationValue(node map[string]any, key string) (time.Duration, bool) {
	v, ok := node[key].(string)
	if !ok {
		return 0, false
	}

	dur, err := duration.ParseISO8601(v)
	if err != nil {
		return 0, false
	}

	var td time.Duration
	td += time.Duration(dur.TH) * time.Hour
	td += time.Duration(dur.TM) * time.Minute
	td += time.Duration(dur.TS) * time.Second

	return td, true
}
