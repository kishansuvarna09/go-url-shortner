package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kishansuvarna09/go-url-shortner/api/database"
)

func DeleteUrl(c *gin.Context) {
	shortID := c.Param("shortID")

	r := database.CreateClient(0)
	defer r.Close()

	err := r.Del(database.Ctx, shortID).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "URL deleted successfully"})
}
