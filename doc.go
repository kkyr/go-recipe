// Package recipe is a Go library that is able to scrape recipes from websites.
//
// The go-recipe default scraper looks for a [Schema Recipe] on the target website encoded in `ld+json` format and is able to retrieve most fields defined in the schema. However, not all websites contain the recipe encoded as such, whereas some others that do can contain incomplete data.
//
// Therefore, this package allows custom scrapers to be defined which contain scraping logic specific to a certain website. The custom scrapers can make use of the default scraper so that custom scraping logic only needs to be defined for fields that the default scraper could not extract any data.
//
// The custom scrapers are registered in [scrapers.go] and are identified by host name, which represents the website that they are used for. When a client provides go-recipe with a link to scrape, the host name is extracted from the link and is used to find the corresponding custom scraper. The default scraper is used if no custom scraper is defined.
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
//	  instructions, ok := scraper.Instructions()
//	  // ... & more fields available
//	}
//
// [Schema Recipe]: https://schema.org/Recipe
// [scrapers.go]: https://github.com/kkyr/go-recipe/blob/main/pkg/recipe/scrapers.go
package recipe
