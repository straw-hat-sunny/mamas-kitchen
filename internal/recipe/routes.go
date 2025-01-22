package recipe

import "github.com/gorilla/mux"

func RegisterRoutes(router *mux.Router) error {

	rh, err := NewRecipeHandler()
	if err != nil {
		return err
	}

	router.HandleFunc("/", rh.HandleListRecipes).Methods("GET")
	router.HandleFunc("/{id}", rh.HandleGetRecipe).Methods("GET")
	return nil
}
