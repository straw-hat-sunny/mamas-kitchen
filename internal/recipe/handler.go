package recipe

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RecipeHandler struct {
	svc Service
}

func NewRecipeHandler() (*RecipeHandler, error) {
	svc, err := NewService()
	if err != nil {
		return nil, err
	}

	return &RecipeHandler{
		svc: svc,
	}, nil
}

// PartialRecipe struct to represent a partial recipe namely the id, name and description
type PartialRecipe struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

type Service interface {
	ListRecipes(ctx context.Context) ([]PartialRecipe, error)
	GetRecipe(ctx context.Context, id string) (*Recipe, error)
}

// ListRecipeResponse struct to represent the response of a list recipe request
type ListRecipeResponse struct {
	Recipes []PartialRecipe `json:"recipes"`
}

func (rh RecipeHandler) HandleListRecipes(w http.ResponseWriter, r *http.Request) {
	resp := &ListRecipeResponse{}
	recipes, err := rh.svc.ListRecipes(r.Context())
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	resp.Recipes = recipes

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (rh RecipeHandler) HandleGetRecipe(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /api/v1/recipes/{id}")
	// get the id from the URL path
	vars := mux.Vars(r)
	id := vars["id"]
	log.Println(id)

	recipeId := id

	// get the recipe from the database
	recipe, err := rh.svc.GetRecipe(r.Context(), recipeId)
	if err != nil {
		log.Println("there was an error")
		// if there is an error, return a 500 internal server error
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// if the recipe is nil, return a 404 not found
	if recipe == nil {
		log.Println("404")
		http.Error(w, "recipe not found", http.StatusNotFound)
		return
	}

	// encode the recipe as JSON and write it to the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipe)
}
