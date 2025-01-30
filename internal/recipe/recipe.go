package recipe

// Ingredient struct to represent an ingredient in a recipe
type Ingredient struct {
	Item     string `json:"item"`
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

// Recipe struct to represent a recipe
type Recipe struct {
	Id           string       `json:"id"`
	Title        string       `json:"title"`
	Type         string       `json:"type"`
	Ingredients  []Ingredient `json:"ingredients"`
	Instructions []string     `json:"instructions"`
}
