package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	g := gin.Default()
	g.StaticFS("/", http.Dir("./"))
	g.Run(":8088")
}
