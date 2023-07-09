package controllers

import (
	"net/http"

	"shop_khordad/models/repositories"

	"github.com/gin-gonic/gin"
)

type LoggingController struct {
	LoggingRepo *repositories.LoggingRepository
}

func (lc *LoggingController) LogActivity(c *gin.Context) {
	userID, err := lc.LoggingRepo.GetUserIDFromSession(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
