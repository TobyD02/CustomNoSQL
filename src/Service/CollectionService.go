package Service

import (
	"CustomNoSQL/types/CollectionTypes"
	"os"
	"path/filepath"
)

func CreateNewCollection(request CollectionTypes.CreateNewCollectionRequestType) error {
	filename := "./data/collections/" + request.CollectionName + ".bson"
	dir := filepath.Dir(filename)

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	collection := request.ToCollection()

	data, err := collection.ToBson()

	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func LoadCollection(request CollectionTypes.LoadCollectionRequestType) (CollectionTypes.LoadCollectionResponseType, error) {
	var response CollectionTypes.LoadCollectionResponseType
	var collection CollectionTypes.CollectionType

	data, err := os.ReadFile("./data/collections/" + request.CollectionName + ".bson")
	if err != nil {
		return response, err
	}

	err = collection.FromBson(data)
	if err != nil {
		return response, err
	}

	response.FromCollection(collection)
	return response, nil
}
