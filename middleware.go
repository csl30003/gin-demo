package main

import "github.com/gin-gonic/gin"

func MyBenchLogger() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func benchEndpoint(c *gin.Context) {

}

func AuthRequired() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
func Login(c *gin.Context) {

}

func analyticsEndpoint(c *gin.Context) {

}

func main() {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	authorized := r.Group("/")
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", Login)

		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
