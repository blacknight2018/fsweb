package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"net/http"
)

var path *string

func main() {
	path = flag.String("path", "./", "path to the directory")
	flag.Parse()
	g := gin.Default()
	g.StaticFS("/", http.Dir(*path))
	g.Run(":8088")
}
