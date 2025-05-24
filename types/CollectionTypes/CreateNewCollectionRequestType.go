package CollectionTypes

type CreateNewCollectionRequestType struct {
	CollectionName string `json:"collection_name"`
}

func (r *CreateNewCollectionRequestType) ToCollection() CollectionType {
	return CollectionType{
		CollectionName: r.CollectionName,
	}
}
