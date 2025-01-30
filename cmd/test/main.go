package main

import (
	"log"
	"mamas-kitchen/internal/recipe"
)

func main() {
	ms, err := recipe.NewMongoStore()
	if err != nil {
		log.Fatal("cant make store", err)

	}
	r, err := ms.ListRecipes()
	if err != nil {
		log.Fatal("cant list docs", err)
	}

	for _, recipe := range r {
		log.Println(recipe)
	}
}
