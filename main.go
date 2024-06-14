package main

import {
	"github.com/gin-gonic/gin"
}


func main() {
	// Create a new Gin router
	r := gin.Default();


	// Define a GET route
	r.GET("/",func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message":"PING";
		})
	})

	r.Run(":8080");
}

// Summary - r.GET("/ping", ...) defines a route that will respond to GET requests at the /ping endpoint.
func(c *gin.Context) { ... } is an anonymous function that takes a gin.Context object (c) as a parameter. This function is called a handler and is executed every time a request matches the route.
c.JSON(200, gin.H{ "message": "pong" }) is used to respond to the request. It sends a JSON response with HTTP status code 200 (OK). The gin.H is a shortcut for a map (map[string]interface{}) to easily create JSON objects.


// Command to run -> go run main.go
// Output -> {"message":"pong"}