package Controller

import (
	"CustomNoSQL/src/Service"
	"CustomNoSQL/types/CollectionTypes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateNewCollection(c *gin.Context) {
	var req CollectionTypes.CreateNewCollectionRequestType

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, CollectionTypes.CreateNewCollectionResponseErrorType{
			Success: false,
			Error:   "Request body is malformed: " + err.Error(),
		})
		return
	}

	err := Service.CreateNewCollection(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, CollectionTypes.CreateNewCollectionResponseErrorType{
			Success: false,
			Error:   "Failed to create new collection: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, CollectionTypes.CreateNewCollectionResponseType{
		Success: true,
		Message: "Successfully created collection: " + req.CollectionName,
	})
}

func LoadCollection(c *gin.Context) {
	var req CollectionTypes.LoadCollectionRequestType

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, CollectionTypes.LoadCollectionResponseErrorType{
			Success: false,
			Error:   "Request body is malformed: " + err.Error(),
		})
		return
	}

	response, err := Service.LoadCollection(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, CollectionTypes.LoadCollectionResponseErrorType{
			Success: false,
			Error:   "Failed to load collection: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
