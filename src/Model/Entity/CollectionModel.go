package Entity

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type CollectionModel struct {
	CollectionName string      `json:"collection_name"`
	Records        []uuid.UUID `json:"records"`
}

func (r *CollectionModel) ToBson() ([]byte, error) {
	return bson.Marshal(r)
}

func (r *CollectionModel) FromBson(data []byte) error {
	return bson.Unmarshal(data, r)
}
