package Service

import (
	"CustomNoSQL/src/ConfigProvider"
	"CustomNoSQL/src/Model/Entity"
	"CustomNoSQL/src/Model/Request"
	"CustomNoSQL/src/Process/Collection"
	"fmt"
	"github.com/google/uuid"
)

// TODO: Any methods that modify collections should use mutex
func CreateNewCollection(request Request.CreateNewCollectionRequestModel) error {
	// Get All Collections File (or create if doesn't exist - process)
	// Verify record name is unique (process)
	// Create new collection model - should be initialised with empty records
	// Create new entry (process) - make sure the name is unique
	// Serialize BSON again
	// Return Success Message
	var config ConfigProvider.Config
	config.GetConfig()

	collectionsPointer, err := Collection.GetAllCollectionsSingletonProcess()
	if err != nil {
		return err
	}

	var collections = *collectionsPointer

	// Check if collection name is already in use
	for collectionName := range collections {
		if collectionName == request.CollectionName {
			return fmt.Errorf("collection name is already in use")
		}
	}

	// Convert request to collection model
	var collection Entity.CollectionModel
	collection = request.ToCollection()

	// Initialise with empty records array
	collection.Records = []uuid.UUID{}

	// Add new collection to collections
	collections[collection.CollectionName] = collection

	return nil
}

func GetAllCollections() ([]string, error) {
	var config ConfigProvider.Config
	config.GetConfig()

	collectionNames := make([]string, 0)

	collections, err := Collection.GetAllCollectionsSingletonProcess()
	if err != nil {
		return collectionNames, err
	}

	for collectionName := range *collections {
		collectionNames = append(collectionNames, collectionName)
	}

	return collectionNames, nil
}
