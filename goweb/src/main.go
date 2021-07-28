package main

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type DataModel struct {
	Name   string `json:"Name"`
	Amount string `json:"Amount"`
}

type DataResponse struct {
	Message string `json:"Message"`
}

func customHandler(c *gin.Context) {
	dataModel := DataModel{}
	response := DataResponse{}

	if err := c.ShouldBindJSON(&dataModel); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	app := "echo"

	arg0 := "-e"
	arg1 := "Hello world"
	arg2 := "\n\tfrom"
	arg3 := "golang"

	cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	// Print the output
	fmt.Println(string(stdout))
	response.Message = string(stdout)
	c.JSON(http.StatusOK, response)
}

func setupRouting() *gin.Engine {
	r := gin.Default()
	r.POST("/rosered/", customHandler)

	return r
}

func main() {
	// r := gin.Default()

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r := setupRouting()

	r.Run(":8081")
}
