package api

import (
	"log"
	"net/http"

	"github.com/benmorehouse/example-service/models"
	"github.com/gin-gonic/gin"
)

// UserIDParam will be used to set the get req UserID
const UserIDParam = "id"

// GetObject
func (s *Server) GetObject(c *gin.Context) {
	userID := c.Param("id")
	log.Println("user id is", userID)
	object, err := s.ObjectTable.Get(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &object)
}

// PostObject will just return nothing, but it should return how it just put something into the database
func (s *Server) PostObject(c *gin.Context) {
	example := models.Example{}
	c.BindJSON(&example)
	if err := s.ObjectTable.Put(example); err != nil {
		c.JSON(http.StatusOK, struct{}{})
	}
	c.JSON(http.StatusOK, struct{}{})
}
