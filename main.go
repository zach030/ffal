package main

import (
	"ffal/ffal"
	"net/http"
)

func main() {
	r := ffal.New()
	r.GET("/", func(c *ffal.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *ffal.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *ffal.Context) {
		c.JSON(http.StatusOK, ffal.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}