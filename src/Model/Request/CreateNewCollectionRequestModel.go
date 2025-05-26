package Request

import "CustomNoSQL/src/Model/Entity"

type CreateNewCollectionRequestModel struct {
	CollectionName string `json:"collection_name"`
}

func (r *CreateNewCollectionRequestModel) ToCollection() Entity.CollectionModel {
	return Entity.CollectionModel{
		CollectionName: r.CollectionName,
	}
}
