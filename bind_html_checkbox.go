package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type myForm struct {
	Colors []string `form:"colors[]"`
}

func main() {
	r := gin.Default()

	r.POST("/ginDemo", func(c *gin.Context) {
		var fakeForm myForm
		err := c.ShouldBind(&fakeForm)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{"color": fakeForm.Colors})
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}

	//<form action="/" method="POST">
	//<p>Check some colors</p>
	//<label for="red">Red</label>
	//<input type="checkbox" name="colors[]" value="red" id="red">
	//<label for="green">Green</label>
	//<input type="checkbox" name="colors[]" value="green" id="green">
	//<label for="blue">Blue</label>
	//<input type="checkbox" name="colors[]" value="blue" id="blue">
	//<input type="submit">
	//</form>
}
