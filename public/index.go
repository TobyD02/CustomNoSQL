package main

import (
	"CustomNoSQL/src/Route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	Route.RegisterRoutes(r)
	r.Run()
}
