package Controller

import (
	"CustomNoSQL/src/Model/Request"
	"CustomNoSQL/src/Model/Response"
	"CustomNoSQL/src/Service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateNewCollection(c *gin.Context) {
	var req Request.CreateNewCollectionRequestModel

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response.CreateNewCollectionResponseErrorModel{
			Success: false,
			Error:   "Request body is malformed: " + err.Error(),
		})
		return
	}

	err := Service.CreateNewCollection(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response.CreateNewCollectionResponseErrorModel{
			Success: false,
			Error:   "Failed to create new collection: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response.CreateNewCollectionResponseModel{
		Success: true,
		Message: "Successfully created collection: " + req.CollectionName,
	})
}

func GetAllCollections(c *gin.Context) {
	response, err := Service.GetAllCollections()

	if err != nil {
		c.JSON(http.StatusBadRequest, Response.GetAllCollectionsResponseErrorModel{
			Success: false,
			Error:   "Failed to load collection: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
