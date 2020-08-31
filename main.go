package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func ParamGroup1Handler(c *gin.Context) {
	param1 := c.Param("param1")
	param2 := c.Param("param2")
	param3 := c.Param("param3")

	if param1 != "" && param2 == "" && param3 == "" { // ("/{param1}")
		c.JSON(http.StatusOK, gin.H{"param1": param1})
	} else if param1 != "" && param2 != "" && param3 == "" { // ("/{param1}/{param2}")
		c.JSON(http.StatusOK, gin.H{"param1": param1, "param2": param2})
	} else if param1 != "" && param2 != "" && param3 != "" { // ("/{param1}/{param2}/{param3}")
		c.JSON(http.StatusOK, gin.H{"param1": param1, "param2": param2, "param3": param3})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Something went wrong with request parameters "})
	}
}

func main() {
	// start application
	log.Println("Starting...")
	r := gin.Default()

	// healthcheck endpoint
	r.GET("/ping", ping)

	// Group endpoints
	ucp := r.Group("/paramgroup1")
	{
		ucp.GET("/:param1", ParamGroup1Handler)
		ucp.GET("/:param1/:param2", ParamGroup1Handler)
		ucp.GET("/:param1/:param2/:param3", ParamGroup1Handler)
	}

	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080
}
