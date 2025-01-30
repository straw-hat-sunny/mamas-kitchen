package recipe

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	mongo_options "go.mongodb.org/mongo-driver/mongo/options"
)

const dbConnectionString = "mongodb://root:example@mongo:27017/"

type fileStore struct {
	data     map[string]*Recipe
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

func loadData(data_dir string) (map[string]*Recipe, error) {
	// iterate over the files in the data directory
	// for each file, read the contents and unmarshal it into a Recipe struct
	// add the Recipe struct to the map with the id as the key
	// return the map

	files, err := os.ReadDir(data_dir)
	if err != nil {
		return nil, err
	}

	recipes := make(map[string]*Recipe)
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

func (fs *fileStore) GetRecipe(id string) (*Recipe, error) {
	recipe, ok := fs.data[id]
	if !ok {
		return nil, ErrRecipeNotFound
	}
	return recipe, nil
}

type MongoStore struct {
	cache      map[string]*Recipe
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoStore() (*MongoStore, error) {
	opts := mongo_options.Client().ApplyURI(dbConnectionString).SetDirect(true)
	c, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}
	db := c.Database("local")

	return &MongoStore{
		client:     c,
		collection: db.Collection("recipes"),
		cache:      make(map[string]*Recipe),
	}, nil
}

func (ms *MongoStore) ListRecipes() ([]Recipe, error) {
	ms.collection.Find(context.Background(), bson.D{})
	collectionSize, err := ms.collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	recipes := make([]Recipe, 0, collectionSize)

	rs, err := ms.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var results []bson.M
	err = rs.All(context.Background(), &results)
	if err != nil {
		return nil, err
	}

	for _, result := range results {
		id, _ := result["_id"].(string)
		title, _ := result["dish_name"].(string)
		dish_type, _ := result["dish_type"].(string)

		recipe := &Recipe{
			Id:    id,
			Title: title,
			Type:  dish_type,
		}
		ms.cache[id] = recipe

		recipes = append(recipes, *recipe)

	}

	return recipes, nil
}

func (ms *MongoStore) GetRecipe(id string) (*Recipe, error) {
	recipe, ok := ms.cache[id]
	if !ok {
		return nil, ErrRecipeNotFound
	}
	return recipe, nil
}
