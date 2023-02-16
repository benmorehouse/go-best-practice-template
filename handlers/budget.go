package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserIDParam will be used to set the get req UserID
const UserIDParam = "id"

// GetBudget will return the
func (s *Server) GetBudget(c *gin.Context) {
	userID := c.Param("id")
	log.Println("user id is", userID)
	budget, err := s.BudgetTable.Get(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	basicResponse := struct {
		Food           int `json:"food"`
		Shopping       int `json:"shopping"`
		Transportation int `json:"travel"`
		Health         int `json:"health"`
		Other          int `json:"other"`
	}{
		budget.Food,
		budget.Shopping,
		budget.Travel,
		budget.Health,
		budget.Other,
	}

	c.JSON(http.StatusOK, &basicResponse)
}

// PostBudget will just return a general set of budgets for now
func (s *Server) PostBudget(c *gin.Context) {
	basicResponse := struct {
		Food     int `json:"food"`
		Shopping int `json:"shopping"`
		Travel   int `json:"travel"`
		Health   int `json:"health"`
		Other    int `json:"other"`
	}{
		100,
		200,
		300,
		400,
		500,
	}

	c.JSON(http.StatusOK, &basicResponse)
}
