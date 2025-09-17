package main

import (
	"github.com/gin-gonic/gin" //non origin library
)

type User struct {
	ID   string `json:"id"` //force to small
	Name string `json:"name"`
}

func main() {
	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		user := []User{{ID: "1", Name: "Nichaphat"}}
		c.JSON(200, user)
	})

	r.Run(":8080") //create web server port 8080

}
