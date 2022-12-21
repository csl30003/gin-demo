package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	//r := gin.Default()
	//http.ListenAndServe(":8080", r)

	r := gin.Default()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		return
	}
}
