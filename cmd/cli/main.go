//nolint
package main

import (
	"fmt"
	"go-recipe/pkg/recipe"
)

const (
	url1 = "https://www.bbcgoodfood.com/recipes/mini-dark-chocolate-blackberry-bay-pavlovas"
	url2 = "https://minimalistbaker.com/smoky-1-pot-refried-lentils"
	url3 = "https://leitesculinaria.com/7759/recipes-pasteis-de-natas.html"
	url4 = "https://www.forksoverknives.com/recipes/vegan-baked-stuffed/baba-ghanoush-flatbreads/blabla"
	url5 = "https://www.forksoverknives.com/recipes/vegan-baked-stuffed/baba-ghanoush-flatbreads"
	url6 = "https://www.bbcgoodfood.com/recipes/noodle-salad-sesame-dressing"
	url7 = "https://cookpad.com/us/recipes/16427532-california-farm-summer-clam-chowder-spicy"
	url8 = "https://copykat.com/dunkin-donuts-caramel-iced-coffee"
)

func main() {
	//res, err := http.Get(url6)
	//if err != nil {
	//	panic(err)
	//}
	//defer res.Body.Close()
	//
	//dst, err := os.Create("forksoverknives.com.html")
	//if err != nil {
	//	panic(err)
	//}
	//
	//io.Copy(dst, res.Body)

	for i, u := range []string{
		url1, url2, url3, url4, url5, url6, url7, url8,
	} {
		fmt.Println(fmt.Sprintf("url%d %s", i+1, u))
		_, err := recipe.Scraper(u)
		fmt.Println()
		if err != nil {
			fmt.Println("err", err)
		}
	}
	//recipe, err := recipe.GetScraper(url8)
	//if err != nil {
	//	panic(err)
	//}

	//ingr, ok := recipe.Ingredients()
	//if !ok {
	//	panic("not ok")
	//}

	//fmt.Printf("%#v\n", ingr)

	//fmt.Println(recipe.Author())
	//fmt.Println(recipe.Categories())
	//fmt.Println(recipe.CookTime())
	//fmt.Println(recipe.Cuisine())
	//fmt.Println(recipe.Description())
	//fmt.Println(recipe.ImageURL())
	//fmt.Println(recipe.Ingredients())
	//fmt.Println(recipe.Instructions())
	//fmt.Println(recipe.Language())
	//fmt.Println(recipe.Name())
	//fmt.Println(recipe.Nutrition())
	//fmt.Println(recipe.PrepTime())
	//fmt.Println(recipe.SuitableDiets())
	//fmt.Println(recipe.TotalTime())
	//fmt.Println(recipe.Yields())
}
