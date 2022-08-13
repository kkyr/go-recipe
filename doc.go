// Package recipe is a Go library for scraping recipes from website.
//
// The go-recipe default scraper relies on the target website containing a [Schema Recipe] encoded in `ld+json` format and is able to retrieve most fields defined by the schema. However, some websites simply do not contain the recipe encoded this way and some others that do can be incomplete.
//
// Therefore, this package allows custom scrapers to be defined which contain scraping logic customized to specific websites. The custom scrapers can make use of the default scraper so that custom scraping logic only needs to be defined for fields that the default scraper could not extract any information.
//
// The custom scrapers are registered in [github.com/kkyr/go-recipe/recipe/scrapers.go] and are identified by host name, which represents the website that they should be used for. When a client provides this package with a link to scrape, the host name is extracted from the link and is used to find the corresponding custom scraper. If none is found, the default scraper is used. In other words, which scraper will be used is a decision made implicitly based on the target website.
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
