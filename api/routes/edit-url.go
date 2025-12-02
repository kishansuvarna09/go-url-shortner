package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kishansuvarna09/go-url-shortner/api/database"
	"github.com/kishansuvarna09/go-url-shortner/api/models"
)

func EditUrl(c *gin.Context) {
	shortID := c.Param("shortID")
	var body models.Request

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(400, gin.H{"error": "Cannot Parse JSON"})
		return
	}

	r := database.CreateClient(0)
	defer r.Close()

	val, err := r.Get(database.Ctx, shortID).Result()
	if err != nil || val == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short ID not found"})
		return
	}

	err = r.Set(database.Ctx, shortID, body.URL, body.Expiry*3600*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "URL updated successfully"})
}
