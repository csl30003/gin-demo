package main

import "github.com/gin-gonic/gin"

func loginEndpoint(c *gin.Context) {

}

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/login", loginEndpoint)
		//  more method
	}

	v2 := r.Group("/v2")
	{
		v2.GET("login", loginEndpoint)
		//  more method
	}

	err := r.Run(":8080")
	if err != nil {
		return
	}
}