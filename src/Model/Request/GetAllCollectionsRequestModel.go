package Request

import "CustomNoSQL/src/Model/Entity"

type GetAllCollectionsRequestModel struct {
	CollectionName string `json:"collection_name"`
}

func (r *GetAllCollectionsRequestModel) ToCollection() Entity.CollectionModel {
	return Entity.CollectionModel{
		CollectionName: r.CollectionName,
	}
}
