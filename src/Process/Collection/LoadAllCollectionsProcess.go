package Collection

import (
	"CustomNoSQL/src/ConfigProvider"
	"CustomNoSQL/src/Model/Entity"
	"errors"
	"fmt"
	"os"
)

// TODO: Should be loaded as key = collection name, records = collection value (as an array of strings - UUID)
// TODO: Modify service class to use this new structure
// TODO: Even better - load into an array as i am doing, append to the array, and then have a serialize method/bsonify
// method that takes them and and turns them into CollectionName: Records in the json file
func loadAllCollectionsProcess() (Entity.AllCollectionsModel, error) {

	var collections Entity.AllCollectionsModel
	var config ConfigProvider.Config
	config.GetConfig() // load config

	collectionFilePath := config.CollectionFilePath

	// Check if file exists
	if _, err := os.Stat(collectionFilePath); errors.Is(err, os.ErrNotExist) {
		fmt.Println("No collections file found, skipping load.")
		return make(Entity.AllCollectionsModel), nil
	} else if err != nil {
		return collections, fmt.Errorf("failed to check file: %w", err)
	}

	// Read BSON file content
	data, err := os.ReadFile(collectionFilePath)
	if err != nil {
		return collections, fmt.Errorf("failed to read BSON file: %w", err)
	}

	// Convert bson data to AllCollectionsModel
	err = collections.FromBson(data)
	if err != nil {
		return collections, err
	}

	fmt.Println(collections)

	return collections, nil
}
