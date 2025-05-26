package Entity

import (
	"go.mongodb.org/mongo-driver/bson"
)

// AllCollectionsModel TODO: Create collections singleton. Should load collections on app startup, and then only return singleton
// - at the moment disk I/O is making the create new collection endpoint take almost 2ms
// - Having the singleton will speed up creating new collection and also fetching collection names
// TODO: Setup some clock to serialize collections every minute or so
// TODO: Can i serialize collections on app end? (defer serialization in public/index.go?)
// TODO: https://refactoring.guru/design-patterns/singleton/go/example

type AllCollectionsModel map[string]CollectionModel

func (c *AllCollectionsModel) FromBson(data []byte) error {
	//var rawMap map[string]bson.M

	err := bson.Unmarshal(data, &c)
	if err != nil {
		return err
	}

	return nil
}

func (c *AllCollectionsModel) ToBson() ([]byte, error) {
	return bson.Marshal(c)
}
