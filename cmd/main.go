package main

import "github.com/gin-gonic/gin"

func main() {

	server := gin.Default()

	server.GET("/teste", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"message": "Successful call",
		})
	})

	server.Run(":8000")

}
