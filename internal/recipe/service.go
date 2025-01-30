package recipe

import (
	"context"
	"log"
)

type Store interface {
	ListRecipes() ([]Recipe, error)
	GetRecipe(id string) (*Recipe, error)
}

type service struct {
	// Add any fields needed for the implementation
	store Store
}

func NewService() (Service, error) {
	//store, err := NewFileStore("data")
	store, err := NewMongoStore()
	if err != nil {
		return nil, err
	}
	return &service{store: store}, nil
}

func (s *service) ListRecipes(ctx context.Context) ([]PartialRecipe, error) {
	// Implement the ListRecipes method
	// return []PartialRecipe{
	// 	{Id: 1, Title: "Spaghetti Bolognese", Type: "main course"},
	// 	{Id: 2, Title: "Chicken Curry", Type: "main course"},
	// 	{Id: 3, Title: "Beef Stew", Type: "main course"},
	// 	{Id: 4, Title: "Apple Pie", Type: "dessert"},
	// 	{Id: 5, Title: "Rum and Coke", Type: "drink"},
	// 	{Id: 6, Title: "Vanilla Ice Cream", Type: "dessert"},
	// 	{Id: 7, Title: "Salad", Type: "appetizer"},
	// }, nil
	recipes, err := s.store.ListRecipes()
	if err != nil {
		return nil, err
	}

	partialRecipes := make([]PartialRecipe, len(recipes))
	for i, recipe := range recipes {
		log.Println(recipe)
		partialRecipes[i] = PartialRecipe{
			Id:    recipe.Id,
			Title: recipe.Title,
			Type:  recipe.Type,
		}
	}

	return partialRecipes, nil
}

func (s *service) GetRecipe(ctx context.Context, id string) (*Recipe, error) {

	recipe, err := s.store.GetRecipe(id)
	if err != nil {
		return nil, err
	}

	return recipe, nil

	// return &Recipe{
	// 	Id:    id,
	// 	Title: "Spaghetti Bolognese",
	// 	Type:  "main course",
	// 	Ingredients: []Ingredient{
	// 		{Item: "spaghetti", Quantity: 200, Unit: "g"},
	// 		{Item: "minced beef", Quantity: 500, Unit: "g"},
	// 		{Item: "onion", Quantity: 1, Unit: ""},
	// 		{Item: "garlic", Quantity: 2, Unit: "cloves"},
	// 		{Item: "tomato sauce", Quantity: 500, Unit: "ml"},
	// 		{Item: "olive oil", Quantity: 2, Unit: "tbsp"},
	// 		{Item: "salt", Quantity: 1, Unit: "tsp"},
	// 		{Item: "pepper", Quantity: 1, Unit: "tsp"},
	// 	},
	// 	Instructions: []string{
	// 		"Chop the onion and garlic",
	// 		"Heat the olive oil in a pan",
	// 		"Add the onion and garlic and cook until soft",
	// 		"Add the minced beef and cook until browned",
	// 		"Add the tomato sauce, salt, and pepper",
	// 		"Simmer for 20 minutes",
	// 		"Cook the spaghetti according to the package instructions",
	// 		"Serve the spaghetti with the bolognese sauce on top",
	// 	},
	// }, nil
}
