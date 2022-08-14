package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/http"
	"github.com/kkyr/go-recipe/internal/scraper/schema"
	urlutil "github.com/kkyr/go-recipe/internal/url"

	"github.com/PuerkitoBio/goquery"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] URL\n\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "OPTIONS:")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage

	domain := flag.String("d", "", "The domain of the website in PascalCase.")
	url := flag.String("u", "", "A url to a recipe on the website.")

	flag.Parse()

	if *domain == "" {
		fmt.Fprintln(os.Stderr, "must supply domain using -d")
		os.Exit(1)
	}

	if *url == "" {
		fmt.Fprintln(os.Stderr, "must supply url using -u")
		os.Exit(1)
	}

	run(*domain, *url)
}

func run(domain, url string) {
	host := urlutil.GetHost(url)
	log.Printf("DEBUG: extracted host %q from url", host)

	dir, err := getScrapersDir()
	if err != nil {
		log.Fatalf("ERROR: failed to get scrapers directory: %v\nMake sure you run this cmd from the go-recipe module", err)
	}

	log.Printf("DEBUG: using scrapers directory %q", dir)

	body, err := http.NewClient().Get(url)
	if err != nil {
		log.Fatalf("ERROR: failed to GET url: %v", err)
	}

	data, err := getRecipeData(body)
	if err != nil {
		data = make(map[string]any)

		log.Printf("WARN: could not get recipe data to prefill test: %v", err)
	}

	data["domain"] = domain
	data["host"] = host

	if err := createScraperFile(dir, strings.ToLower(domain), data); err != nil {
		log.Fatalf("ERROR: unable to create go file for scraper: %v", err)
	}

	if err := createScraperTestFile(dir, strings.ToLower(domain), data); err != nil {
		log.Fatalf("ERROR: unable to create go test file for scraper: %v", err)
	}

	if err := createTestDataFile(dir, host, body); err != nil {
		log.Fatalf("ERROR: unable to create test data file: %v", err)
	}
}

func getScrapersDir() (string, error) {
	root, err := getModuleRoot()
	if err != nil {
		return "", fmt.Errorf("could not get module root: %w", err)
	}

	return filepath.Join(root, "internal", "scraper", "custom"), nil
}

func getRecipeData(body []byte) (map[string]any, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("could not create document: %v", err)
	}

	scraper, err := schema.GetRecipeScraper(doc)
	if err != nil {
		return nil, fmt.Errorf("failed to create recipe scraper: %v", err)
	}

	return scraperToMap(scraper), nil
}

func createScraperFile(dir string, name string, data map[string]any) error {
	path := getFilePath(dir, name)
	return createFile(path, scraperTmpl, data)
}

func createScraperTestFile(dir string, name string, data map[string]any) error {
	path := getTestFilePath(dir, name)
	return createFile(path, scraperTestTmpl, data)
}

const newFileFlags = os.O_CREATE | os.O_WRONLY | os.O_EXCL

func createTestDataFile(dir string, name string, body []byte) error {
	path := getTestDataFilePath(dir, name)

	f, err := os.OpenFile(path, newFileFlags, 0644)
	if err != nil {
		return fmt.Errorf("open file failed: %v", err)
	}
	defer f.Close()

	if _, err := f.Write(body); err != nil {
		return fmt.Errorf("write file failed: %v", err)
	}

	log.Printf("INFO: created file %s", path)

	return nil
}

func createFile(path string, tmpl *template.Template, data map[string]any) error {
	f, err := os.OpenFile(path, newFileFlags, 0644)
	if err != nil {
		return fmt.Errorf("open file failed: %v", err)
	}
	defer f.Close()

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("template execute failed: %v", err)
	}

	// run gofmt on source
	b, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("could not format source file: %v", err)
	}

	if _, err := f.Write(b); err != nil {
		return fmt.Errorf("write to file failed: %v", err)
	}

	log.Printf("INFO: created file %s", path)

	return nil
}

func getFilePath(dir string, name string) string {
	file := fmt.Sprintf("%s.go", name)
	return filepath.Join(dir, file)
}

func getTestFilePath(dir string, name string) string {
	file := fmt.Sprintf("%s_test.go", name)
	return filepath.Join(dir, file)
}

func getTestDataFilePath(dir string, name string) string {
	file := fmt.Sprintf("%s.html", name)
	return filepath.Join(dir, "testdata", file)
}

func getModuleRoot() (string, error) {
	cmd := exec.Command("go", "env", "GOMOD")

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("go env failed: %v", err)
	}

	path := strings.TrimSpace(out.String())
	dir := filepath.Dir(path)

	return dir, nil
}

func scraperToMap(scraper *schema.RecipeScraper) map[string]any {
	author, _ := scraper.Author()
	categories, _ := scraper.Categories()
	cookTime, _ := scraper.CookTime()
	cuisine, _ := scraper.Cuisine()
	description, _ := scraper.Description()
	imageURL, _ := scraper.ImageURL()
	ingredients, _ := scraper.Ingredients()
	instructions, _ := scraper.Instructions()
	language, _ := scraper.Language()
	name, _ := scraper.Name()
	nutrition, _ := scraper.Nutrition()
	prepTime, _ := scraper.PrepTime()
	suitableDiets, _ := scraper.SuitableDiets()
	totalTime, _ := scraper.TotalTime()
	yields, _ := scraper.Yields()

	return map[string]any{
		"author":        fmt.Sprintf("%q", author),
		"categories":    sliceGoString(categories),
		"cookTime":      durationGoString(cookTime),
		"cuisine":       sliceGoString(cuisine),
		"description":   fmt.Sprintf("%q", description),
		"imageURL":      fmt.Sprintf("%q", imageURL),
		"ingredients":   sliceGoString(ingredients),
		"instructions":  sliceGoString(instructions),
		"language":      fmt.Sprintf("%q", language),
		"name":          fmt.Sprintf("%q", name),
		"nutrition":     fmt.Sprintf("%#v", nutrition),
		"prepTime":      durationGoString(prepTime),
		"suitableDiets": dietSliceGoString(suitableDiets),
		"totalTime":     durationGoString(totalTime),
		"yields":        fmt.Sprintf("%q", yields),
	}
}

func durationGoString(d time.Duration) string {
	return fmt.Sprintf("%.0f * time.Minute", d.Minutes())
}

func sliceGoString(s []string) string {
	if s == nil {
		return "nil"
	}

	return fmt.Sprintf("%#v", s)
}

func dietSliceGoString(diets []recipe.Diet) string {
	if diets == nil {
		return "nil"
	}

	dietStrs := make([]string, 0, len(diets))

	if len(diets) > 0 {
		dietType := reflect.TypeOf(diets[0]).String()
		dietPkg, _, _ := strings.Cut(dietType, ".")

		for _, diet := range diets {
			dietStrs = append(dietStrs, fmt.Sprintf("%s.%s", dietPkg, diet.String()))
		}
	}

	// e.g. "[]recipe.Diet{VeganDiet, VegetarianDiet}"
	return fmt.Sprintf("%s{%s}", reflect.TypeOf(diets).String(), strings.Join(dietStrs, ", "))
}
