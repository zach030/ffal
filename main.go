package main

import (
	"ffal/ffal"
	"net/http"
)

func main() {
	r := ffal.New()
	r1 := r.Group("/api")
	{
		r1.GET("/", func(c *ffal.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		r1.GET("/hello", func(c *ffal.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	r2 := r.Group("/open")
	{
		r2.GET("/hello/:name", func(c *ffal.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})

		r2.GET("/assets/*filepath", func(c *ffal.Context) {
			c.JSON(http.StatusOK, ffal.H{"filepath": c.Param("filepath")})
		})
	}

	r.Run(":9999")
}
