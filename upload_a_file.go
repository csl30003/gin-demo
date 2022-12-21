package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	//  设置较低的内存限制（默认是32MB）
	r.MaxMultipartMemory = 8 << 20 //  8MB
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		dst := "./" + file.Filename
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
