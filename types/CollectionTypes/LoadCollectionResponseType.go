package CollectionTypes

type LoadCollectionResponseErrorType struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type LoadCollectionResponseType struct {
	Success    bool           `json:"success"`
	Collection CollectionType `json:"collection"`
}

func (r *LoadCollectionResponseType) FromCollection(collection CollectionType) {
	r.Collection = collection
	r.Success = true
}
