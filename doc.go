// Package recipe is a Go library that scrapes recipes from websites.
//
// The go-recipe default scraper looks for a `ld+json` encoded [Schema Recipe] on the target website and is able to retrieve most fields defined in the schema. However, some websites have incomplete Schema data or simply do not encode their recipe in such a format.
//
// Therefore, custom scrapers exist that are used to scrape specific websites. These scrapers can make use of the default scraper so that custom scraping logic is only defined for fields that the default scraper could not find any data for.
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
//	  recipe, err := recipe.ScrapeFrom(url)
//	  if err != nil {
//		// handle err
//	  }
//
//	  ingredients, ok := recipe.Ingredients()
//	  instructions, ok := recipe.Instructions()
//	  // ... & more fields available
//	}
//
// [Schema Recipe]: https://schema.org/Recipe
// [scrapers.go]: https://github.com/kkyr/go-recipe/blob/main/pkg/recipe/scrapers.go
package recipe
