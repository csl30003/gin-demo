package main

import "github.com/gin-gonic/gin"

type person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	r := gin.Default()

	r.GET("/:name/:id", func(c *gin.Context) {
		var p person

		err := c.ShouldBindUri(&p)
		if err != nil {
			c.JSON(400, gin.H{"msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"name": p.Name, "uuid": p.ID})
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
