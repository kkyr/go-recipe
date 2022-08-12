package recipe

import "time"

// Scraper is a type that is returns recipe information from an underlying data source.
type Scraper interface {
	// Author is the author of the recipe.
	Author() (string, bool)
	// Categories are the categories of the recipe, e.g. appetizer, entrée, etc.
	Categories() ([]string, bool)
	// CookTime is the time it takes to actually cook the dish.
	CookTime() (time.Duration, bool)
	// Cuisine is the cuisine of the recipe, e.g. mexican-inspired, french, etc.
	Cuisine() ([]string, bool)
	// Description is the description of the recipe.
	Description() (string, bool)
	// ImageURL is a URL to an image of the dish.
	ImageURL() (string, bool)
	// Ingredients are all the ingredients used in the recipe.
	Ingredients() ([]string, bool)
	// Instructions are all the steps in making the recipe.
	Instructions() ([]string, bool)
	// Language is the language used in the recipe expressed in IETF BCP 47 standard.
	Language() (string, bool)
	// Name is the name of the recipe.
	Name() (string, bool)
	// Nutrition is nutritional information about the dish.
	Nutrition() (Nutrition, bool)
	// PrepTime is the length of time it takes to prepare the items to be used in the instructions.
	PrepTime() (time.Duration, bool)
	// SuitableDiets indicates dietary restrictions or guidelines for which the recipe is suitable.
	SuitableDiets() ([]Diet, bool)
	// TotalTime is the total time required to perform the instructions (including the prep time).
	TotalTime() (time.Duration, bool)
	// Yields is the quantity that results from the recipe.
	Yields() (string, bool)
}
