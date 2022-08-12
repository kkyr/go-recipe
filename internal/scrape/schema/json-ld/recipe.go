package ld

import (
	"encoding/json"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/piprate/json-gold/ld"
)

const (
	graphKey       = "@graph"
	typeKey        = "@type"
	recipeType     = "Recipe"
	jsonLdSelector = `script[type="application/ld+json"]`
)

// NewRecipeProcessor returns a RecipeProcessor with default settings.
func NewRecipeProcessor() *RecipeProcessor {
	return &RecipeProcessor{
		proc: ld.NewJsonLdProcessor(),
		opts: ld.NewJsonLdOptions(""),
		ctx: map[string]any{
			"@context": "http://schema.org/",
			"@type":    "Recipe",
		},
	}
}

// RecipeProcessor is a json-ld Schema Recipe processor.
type RecipeProcessor struct {
	proc *ld.JsonLdProcessor
	opts *ld.JsonLdOptions
	ctx  map[string]any
}

// GetRecipeNode searches doc to find a Schema.org Recipe node encoded in ld+json format.
// If found, the Recipe is serialized into a map. Individual recipe fields can be accessed
// in the map using the field names defined in https://schema.org/Recipe.
func (rp *RecipeProcessor) GetRecipeNode(doc *goquery.Document) (map[string]any, error) {
	jsonLdDocs := doc.Find(jsonLdSelector).Map(func(_ int, sel *goquery.Selection) string {
		return sel.Text()
	})
	if len(jsonLdDocs) == 0 {
		return nil, fmt.Errorf("no ld+json document found")
	}

	var (
		node map[string]any
		err  error
	)

	for _, doc := range jsonLdDocs {
		if node, err = rp.parseJSON(doc); err == nil {
			return node, nil
		}
	}

	return nil, err
}

func (rp *RecipeProcessor) parseJSON(data string) (map[string]any, error) {
	var nodeMap map[string]any
	if err := json.Unmarshal([]byte(data), &nodeMap); err != nil {
		return nil, fmt.Errorf("unmarshal data failed: %v", err)
	}

	var nodes []any
	if isGraphNode(nodeMap) {
		nodes = ld.Arrayify(nodeMap[graphKey])
	} else {
		nodes = ld.Arrayify(nodeMap)
	}

	recipeNode, ok := findRecipeNode(nodes)
	if !ok {
		return nil, fmt.Errorf("could not find Recipe node")
	}

	recipeNode, err := rp.proc.Compact(recipeNode, rp.ctx, rp.opts)
	if err != nil {
		return nil, fmt.Errorf("could not compact Recipe node: %v", err)
	}

	return recipeNode, nil
}

func isGraphNode(v any) bool {
	vMap, isMap := v.(map[string]any)
	_, containsGraph := vMap[graphKey]

	return isMap && containsGraph
}

func findRecipeNode(nodes []any) (map[string]any, bool) {
	for _, node := range nodes {
		if m, ok := node.(map[string]any); ok {
			if m[typeKey] == recipeType {
				return m, true
			}
		}
	}

	return nil, false
}
