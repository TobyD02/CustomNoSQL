package Route

import (
	"CustomNoSQL/src/Controller"
	"github.com/gin-gonic/gin"
)

func RegisterCollectionRoutes(r *gin.Engine) {
	r.PUT("/collection", Controller.CreateNewCollection)
	r.POST("/collection", Controller.LoadCollection)
}
