<p align="center">
    <img src="logo.png" alt="go-recipe-logo" title="go-recipe" class="img-responsive" />
</p>

<p align="center">
    <a href="https://pkg.go.dev/github.com/kkyr/go-recipe?tab=doc"><img src="https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white" alt="godoc" title="godoc"/></a>
    <a href="https://github.com/kkyr/go-recipe/tags"><img src="https://img.shields.io/github/v/tag/kkyr/go-recipe" alt="semver tag" title="semver tag"/></a>
    <a href="https://goreportcard.com/report/github.com/kkyr/go-recipe"><img src="https://goreportcard.com/badge/github.com/kkyr/go-recipe" alt="go report card" title="go report card"/></a>
    <a href="https://coveralls.io/github/kkyr/go-recipe?branch=main"><img src="https://coveralls.io/repos/github/kkyr/go-recipe/badge.svg?branch=main" alt="coverage status" title="coverage status"/></a>
    <a href="https://github.com/kkyr/go-recipe/blob/main/LICENSE"><img src="https://img.shields.io/github/license/kkyr/go-recipe" alt="license" title="license"/></a>
</p>

# go-recipe

go-recipe is a Go library that scrapes recipes from websites.

## Installation

`$ go get github.com/kkyr/go-recipe@latest`

## Usage

```go
package main

import "github.com/kkyr/go-recipe/pkg/recipe"

func main() {
  url := "https://minimalistbaker.com/quick-pickled-jalapenos/"

  recipe, err := recipe.ScrapeFrom(url)
  if err != nil { 
      // handle err
  }
  
  ingredients, ok := recipe.Ingredients() 
  instructions, ok := recipe.Instructions()
  // ... & more fields available
}
```

## Scraping

The go-recipe default scraper looks for a `ld+json` encoded [Schema Recipe](https://schema.org/Recipe) on the target website and is able to retrieve most fields defined in the schema. However, some websites have incomplete Schema data or simply do not encode their recipe in such a format.

Therefore, custom scrapers exist that are used to scrape specific websites. These scrapers can make use of the default scraper so that custom scraping logic is only defined for fields that the default scraper could not find any data for.

The custom scrapers are registered in [pkg/recipe/scrapers.go](/pkg/recipe/scrapers.go) and are identified by host name, which represents the website that they are used for. When a client provides go-recipe with a link to scrape, the host name is extracted from the link and is used to find the corresponding custom scraper. The default scraper is used if no custom scraper is defined.

## Contributing

Contributions are welcome! You can contribute in a few ways: by adding a custom scraper, patching a bug, implementing a feature based on the roadmap (see further below), or by incorporating any other feature that you'd like to see. For the latter case, please first open a discussion so that we can align on the change before you start coding.

### Custom scrapers

Creating a custom scraper is easy as pie thanks to the code generator that's included in this package. 

The generator requires two arguments: a link to a recipe on a website and the domain of the website that the recipe is hosted on. The domain is used to generate the source code (particularly the file and struct names), and the link is used to scrape recipe data, which is then used to generate a fully functioning unit test. If the generator is unable to scrape recipe data (which can happen if the website does not contain a Schema Recipe), a test will still be generated but test assertions will be made against empty fields.

To use the code generator, run the following command while inside the go-recipe package:

```shell
$ go run cmd/scrapergen/*.go \
  -d CopyKat \
  -u https://copykat.com/dunkin-donuts-caramel-iced-coffee
```

_(replace the domain and link with your own)_

**Important:** the domain should be provided in **PascalCase** so that the generated structs are correctly cased. Otherwise, your PR will not get approved.

The sample command above would generate the following files:

```shell
go-recipe/internal/html/scrape/custom/copykat.go
go-recipe/internal/html/scrape/custom/copykat_test.go
go-recipe/internal/html/scrape/custom/testdata/copykat.com.html
```

The generator can't do everything (at least not yet), so there's some final touches that you must put in:
1. Register your custom scraper in the `hostToScraper` map located in [pkg/recipe/scrapers.go](/pkg/recipe/scrapers.go). Please maintain alphabetical ordering.
2. (Optional) Modify the custom scraper to add your own scraping logic.
3. Verify that the generated test is correct.
4. Verify that `make lint` and `make test` are passing.

You should now be ready to send a PR!

### Roadmap

- [ ] Refactor `scrapergen` and add unit tests.
- [ ] Modify `scrapergen` so that it also adds the new scraper to the `hostToScraper` map.
- [ ] Add option for user to specify http client timeout.
- [ ] Add option for user to specify "strict" mode: in this mode only custom scrapers will be used if defined, otherwise fail.
- [ ] Add more Schema Recipe fields.
- [ ] Add CLI wrapper over the scraper, providing output in JSON.

## Acknowledgements

go-recipe is heavily inspired [recipe-scrapers](https://github.com/hhursev/recipe-scrapers).
