package CollectionTypes

import "go.mongodb.org/mongo-driver/bson"

type CollectionType struct {
	CollectionName string `json:"collection_name"`
}

func (r *CollectionType) ToBson() ([]byte, error) {
	return bson.Marshal(r)
}

func (r *CollectionType) FromBson(data []byte) error {
	return bson.Unmarshal(data, r)
}
