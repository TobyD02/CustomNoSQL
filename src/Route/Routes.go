package Route

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	RegisterCollectionRoutes(r)
}
