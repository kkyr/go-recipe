// Package recipe is a Go library for scraping website recipes.
//
// The go-recipe default scraper relies on the target website containing a [Schema Recipe] encoded in `ld+json` format which knows how to retrieve most fields defined in the schema. However, some websites do not contain the recipe in Schema format and some others that do can be incomplete.
//
// Therefore, this package allows custom scrapers to be defined which contain customized scraping logic for each website. The custom scrapers can also use the default scraper in a "hybrid mode" so that custom scraping logic only needs to be defined for those fields which the default scraper could not pick anything up.
//
// The custom scrapers are registered in [pkg/recipe/scrapers.go](/pkg/recipe/scrapers.go) and are identified by the host name of the website for which they are defined. go-recipe makes use of this host name mapping when picking a scraper to use for the given url. In other words, which scraper will be used by the package is an implicit decision based on the host of the target website and, of course, based on the custom scrapers that are available.
//
// # Example
//
//	package main
//
//	import "github.com/kkyr/go-recipe/pkg/recipe"
//
//	func main() {
//	  url := "https://minimalistbaker.com/quick-pickled-jalapenos/"
//
//	  scraper, err := recipe.GetScraper(url)
//	  if err != nil {
//		panic(err)
//	  }
//
//	  ingredients, ok := scraper.Ingredients()
//	  // + many more fields available
//	}
//
// [Schema Recipe]: https://schema.org/Recipe
package recipe
