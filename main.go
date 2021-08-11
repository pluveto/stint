package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	gin.SetMode(gin.DebugMode)
	engine.Run("127.0.0.1:8989")
}
