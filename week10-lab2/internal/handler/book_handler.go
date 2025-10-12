package handler

import "github.com/gin-gonic/gin"

type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// @Summary Get all books
// @Description Get details of a books
// @Tags Books
// @Produce  json
// @Param   id   path      int     true  "Book ID"
// @Success 200  {array}  Book
// @Failure 500  {object}  ErrorResponse
// @Router  /books/ [get]
func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id, "title": "Mastering golang"})
}
