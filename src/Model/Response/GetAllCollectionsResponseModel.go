package Response

type GetAllCollectionsResponseModel struct {
	Success     bool     `json:"success"`
	Collections []string `json:"collections"`
}

type GetAllCollectionsResponseErrorModel struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
