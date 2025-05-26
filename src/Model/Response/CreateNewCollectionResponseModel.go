package Response

type CreateNewCollectionResponseModel struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type CreateNewCollectionResponseErrorModel struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
