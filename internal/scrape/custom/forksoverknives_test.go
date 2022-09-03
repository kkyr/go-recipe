package custom_test

import (
	"testing"
	"time"

	"github.com/kkyr/go-recipe"
	"github.com/kkyr/go-recipe/internal/scrape/custom"
	"github.com/kkyr/go-recipe/internal/scrape/test"

	"github.com/kkyr/assert"
)

func TestNewForksOverKnivesScraper(t *testing.T) {
	doc := test.ReadHTMLFileOrFail(t, custom.ForksOverKnivesHost)

	scraper, err := custom.NewForksOverKnivesScraper(doc)
	assert.New(t).Require().Nil(err)

	scraperTest := test.Scraper{
		Author:        "Shelli McConnell",
		Categories:    []string{"<a href=\"https://www.forksoverknives.com/recipes/vegan-baked-stuffed/\">Baked & Stuffed</a>"},
		CookTime:      0,
		Cuisine:       nil,
		Description:   "Baba ghanoush, made with tahini, blended eggplant, and mashed chickpeas, can serve as a delicious, creamy sauce for flatbreads. Quick-pickled veggies add color, crunch, and tang. Be sure to drain off any excess liquid before adding the pickles.",
		ImageURL:      "https://www.forksoverknives.com/wp-content/uploads/Baba-Ghanoush-Flatbread-WORDPRESS.jpg",
		Ingredients:   []string{"Cornmeal", "1 recipe <a href=\"https://www.forksoverknives.com/recipes/oil-free-vegan-pizza-dough\">Homemade Oil-Free Pizza Dough</a>", "½ cup matchstick-cut cucumber", "½ cup matchstick-cut carrot", "¼ cup matchstick-cut radishes", "2 tablespoons white wine vinegar", "½ teaspoon pure cane sugar", "⅛ teaspoon sea salt", "1 small eggplant (12 oz.)", "1 tablespoon tahini", "½ teaspoon lemon zest", "1 tablespoon lemon juice", "1 cup chopped fresh parsley", "1 small clove garlic, minced", "1 cup no-salt-added canned chickpeas, rinsed and drained", "Crushed red pepper flakes"},
		Instructions:  []string{"Preheat oven to 400°F. Lightly sprinkle a large baking sheet with cornmeal.", "Divide dough into four portions. On a lightly floured surface, roll portions into 7- to 8-inch circles or 10×5-inch ovals. Transfer flatbreads to prepared pan. Bake 10 to 13 minutes or until lightly browned and set (flatbreads may puff). Let cool.", "Preheat oven to 425°F. For quick pickles, in a small bowl combine the next six ingredients (through salt). Cover and chill until ready to use.", "Meanwhile, line a large baking sheet with foil. Prick eggplant with a fork. Place on prepared pan. Roast 30 minutes or until very soft and skin is charred. Cool until easy to handle.", "Cut eggplant in half; scoop out flesh and place in a food processor. Discard skin. Add the next five ingredients (through garlic); cover and pulse until nearly smooth. In a small bowl coarsely mash the chickpeas. Stir in eggplant mixture. Season with additional salt to taste.", "Spread eggplant mixture on flatbreads. Top with quick pickles. Sprinkle with crushed red pepper and additional parsley."},
		Language:      "",
		Name:          "Baba Ghanoush Flatbreads",
		Nutrition:     recipe.Nutrition{},
		PrepTime:      30 * time.Minute,
		SuitableDiets: nil,
		TotalTime:     1 * time.Hour,
		Yields:        "4 flatbreads",
	}

	scraperTest.Run(t, scraper)
}
