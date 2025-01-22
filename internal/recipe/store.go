package recipe

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type fileStore struct {
	data     map[int]*Recipe
	data_dir string
}

var ErrRecipeNotFound = errors.New("recipe not found")

func NewFileStore(data_dir string) (*fileStore, error) {
	data, err := loadData(data_dir)
	if err != nil {
		return nil, err
	}

	return &fileStore{
		data:     data,
		data_dir: data_dir,
	}, nil
}

func loadData(data_dir string) (map[int]*Recipe, error) {
	// iterate over the files in the data directory
	// for each file, read the contents and unmarshal it into a Recipe struct
	// add the Recipe struct to the map with the id as the key
	// return the map

	files, err := os.ReadDir(data_dir)
	if err != nil {
		return nil, err
	}

	recipes := make(map[int]*Recipe)
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := filepath.Join(data_dir, file.Name())
		data, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		var recipe Recipe
		err = json.Unmarshal(data, &recipe)
		if err != nil {
			return nil, err
		}

		recipes[recipe.Id] = &recipe
	}

	return recipes, nil

}

func (fs *fileStore) ListRecipes() ([]Recipe, error) {
	recipes := make([]Recipe, 0, len(fs.data))
	for _, recipe := range fs.data {
		recipes = append(recipes, *recipe)
	}
	return recipes, nil
}

func (fs *fileStore) GetRecipe(id int) (*Recipe, error) {
	recipe, ok := fs.data[id]
	if !ok {
		return nil, ErrRecipeNotFound
	}
	return recipe, nil
}
