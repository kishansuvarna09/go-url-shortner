package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kishansuvarna09/go-url-shortner/api/database"
)

type AddTagRequest struct {
	ShortID string `json:"short_id" binding:"required"`
	Tag     string `json:"tag" binding:"required"`
}

func AddTag(c *gin.Context) {
	var body AddTagRequest

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(400, gin.H{"error": "Cannot Parse JSON"})
		return
	}

	shortId := body.ShortID
	tag := body.Tag

	r := database.CreateClient(0)
	defer r.Close()

	val, err := r.Get(database.Ctx, shortId).Result()
	if err != nil || val == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short ID not found"})
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(val), &data); err != nil {
		data = make(map[string]interface{})
		data["data"] = val
	}

	var tags []string
	if existingTags, ok := data["tags"].([]interface{}); ok {
		for _, t := range existingTags {
			if strTag, ok := t.(string); ok {
				tags = append(tags, strTag)
			}
		}
	}

	for _, existingTag := range tags {
		if existingTag == tag {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Tag already exists"})
			return
		}
	}

	tags = append(tags, tag)
	data["tags"] = tags

	updatedVal, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update URL"})
		return
	}

	err = r.Set(database.Ctx, shortId, updatedVal, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Tag added successfully"})
}
