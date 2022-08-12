<p align="center">
    <img src="logo.png" alt="go-recipe-logo" title="go-recipe" class="img-responsive" />
</p>

<p align="center">
    <a href="https://pkg.go.dev/github.com/kkyr/go-recipe?tab=doc"><img src="https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white" alt="godoc" title="godoc"/></a>
    <a href="https://goreportcard.com/report/github.com/kkyr/go-recipe"><img src="https://goreportcard.com/badge/github.com/kkyr/go-recipe" alt="go report card" title="go report card"/></a>
    <a href="https://coveralls.io/github/kkyr/go-recipe?branch=main"><img src="https://coveralls.io/repos/github/kkyr/go-recipe/badge.svg?branch=main" alt="coverage status" title="coverage status"/></a>
    <a href="https://github.com/kkyr/go-recipe/blob/main/LICENSE"><img src="https://img.shields.io/github/license/kkyr/go-recipe" alt="license" title="license"/></a>
</p>

# go-recipe

go-recipe is a Go library for scraping website recipes.

## Getting started

Pull the package into your module:

`$ go get github.com/kkyr/go-recipe@latest`

And use it in code:

```go
package main

import "go-recipe/pkg/recipe"

func main() {
  url := "https://minimalistbaker.com/quick-pickled-jalapenos/"
	
  scraper, err := recipe.Scraper(url)
  if err != nil {
	panic(err)
  }
	
  ingredients, ok := scraper.Ingredients() 
  // + many more fields available
}
```

## Scraping

The go-recipe default scraper relies on the target website containing a [Schema Recipe](https://schema.org/Recipe) encoded in `ld+json` format which knows how to retrieve most fields defined in the schema. However, some websites do not contain the recipe in Schema format and some others that do can be incomplete.

Therefore, this package allows custom scrapers to be defined which contain customized scraping logic for each website. The custom scrapers can also use the default scraper in a "hybrid mode" so that custom scraping logic only needs to be defined for those fields which the default scraper could not pick anything up.

The custom scrapers are registered in [pkg/recipe/scrapers.go](/pkg/recipe/scrapers.go) and are identified by the host name of the website for which they are defined. go-recipe makes use of this host name mapping when picking a scraper to use for the given url. In other words, which scraper will be used by the package is an implicit decision based on the host of the target website and, of course, based on the custom scrapers that are available.

## Contributing

Contributions are very welcome! You can contribute in a few ways: by adding a custom scraper, patching a bug, implementing a feature based on the roadmap (see below), or simply incorporating a feature that you'd like to see. For the latter case, please open an issue so that we align on the feature and approach before you start coding.

### Custom scrapers

Creating a custom scraper for your website is easy as pie using the code generator that's included in this package. All the generator needs is a domain and a recipe URL of the website that you'd like to create the scraper for. The domain is used to generate the source code files and the url to fetch recipe data which is then used to generate a fully-functioning unit test.

To run the code generator, make sure you're located in the go-recipe package and run the following command:

```shell
$ go run cmd/scrapergen/*.go \
  -d CopyKat \
  -u https://copykat.com/dunkin-donuts-caramel-iced-coffee
```

_(replace the domain and url with your own)_

**Important:** the domain should be provided in **PascalCase** so that the generated scraper types are correctly cased. Otherwise, your PR will not get approved.

The code generator creates the following files:

```shell
go-recipe/internal/scrape/custom/copykat.go
go-recipe/internal/scrape/custom/copykat_test.go
go-recipe/internal/scrape/custom/testdata/copykat.com.html
```

The generator can't do everything (at least not yet), so there's some final touches that you must put in:
1. Register your custom scraper in the `hostToScraper` map located in [pkg/recipe/scrapers.go](/pkg/recipe/scrapers.go). Please maintain alphabetical ordering.
2. (Optional) Modify the custom scraper to add your own scraping logic.
3. Verify that the generated test is correct.
4. Verify that `make lint` and `make test` are passing.

You should now be good to send a PR!

### Roadmap

- [ ] Add `scrapegen` unit tests.
- [ ] Modify `scrapegen` so that it also adds the new scraper to the `hostToScraper` map.
- [ ] Add option for user to specify http client timeout.
- [ ] Add option for user to specify "strict" mode: in this mode only custom scrapers will be used.
- [ ] Add more Schema Recipe fields.
- [ ] Add CLI wrapper over the scraper, providing output in JSON.

## Warning

Given that I created this package very recently, I haven't had the chance to yet use it in a project nor have I received any feedback. Therefore, you should expect breaking changes in the short term based on community feedback and/or personal discoveries that I personally make as a client of the package.

Once things start to stabilize, I'm aiming to publish a 0.1 version which should have a relatively stable public API. This shouldn't take more than a few months.

## Acknowledgements

This project is heavily inspired by the Python [recipe-scrapers](https://github.com/hhursev/recipe-scrapers) package. Big thanks to its creators and maintainers.
