package main

import (
	"ffal/ffal"
	"net/http"
)

func main() {
	engine := ffal.New()
	engine.GET("/index", func(c *ffal.Context) {
		c.HTML(http.StatusOK,"<h1>Index Page</h1>")
	})
	engine.Run(":8009")
}
