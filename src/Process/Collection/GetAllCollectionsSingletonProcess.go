package Collection

import (
	"CustomNoSQL/src/Model/Entity"
	"sync"
)

var lock = &sync.Mutex{}
var allCollectionsSingleton *Entity.AllCollectionsModel

func GetAllCollectionsSingletonProcess() (*Entity.AllCollectionsModel, error) {
	if allCollectionsSingleton == nil {
		lock.Lock()
		defer lock.Unlock()
		if allCollectionsSingleton == nil {
			model, err := loadAllCollectionsProcess()
			if err != nil {
				return allCollectionsSingleton, nil
			}

			allCollectionsSingleton = &model
		}
	}

	return allCollectionsSingleton, nil
}
