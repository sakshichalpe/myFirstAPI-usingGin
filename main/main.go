package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	//setting up a simple route
	r.GET("/Welcome", firstpg)
	r.POST("/getString", getString)
	r.Run("localhost:9090")
}
func firstpg(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello, welcome to the ODP Corporation!"})
}

type Value struct {
	Str string `json:"str"`
}

func getString(c *gin.Context) {
	var val Value
	err := c.BindJSON(&val)
	if err != nil {
		fmt.Errorf("Error in Binding %s", err.Error())
	}
	fmt.Println("str before", val.Str)
	reversed := reverseString(val.Str)
	fmt.Println("str After", reversed)

	c.JSON(http.StatusOK, reversed)

}
func reverseString(str string) string {
	fmt.Println(str)
	ans := ""
	for i := len(str) - 1; i >= 0; i-- {
		ans = ans + string(str[i])
	}
	return ans
}

//
