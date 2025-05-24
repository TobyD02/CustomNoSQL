package CollectionTypes

type CreateNewCollectionResponseType struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type CreateNewCollectionResponseErrorType struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
