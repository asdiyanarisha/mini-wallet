package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"mini-wallet/internal/http"
)

func main() {
	g := gin.New()
	http.NewHttp(g)

	if err := g.Run(":8080"); err != nil {
		log.Fatal("Can't start server.")
	}
}
