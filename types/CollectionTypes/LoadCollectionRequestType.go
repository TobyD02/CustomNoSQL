package CollectionTypes

type LoadCollectionRequestType struct {
	CollectionName string `json:"collection_name"`
}

func (r *LoadCollectionRequestType) FromCollection(collection CollectionType) {
	r.CollectionName = collection.CollectionName
}
