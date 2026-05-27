package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// I already tried gin on another repo of mine, see:
// https://github.com/AlexandreSJ/puddle

func main() {
	r := SetupRouter()
	if err := r.Run(":3000"); err != nil {
		println("failed to run server: %v", err)
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", Ping)
	return r
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
