package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"fesnuk-api/models"
)

// CreateLike handles POST /likes
func CreateLike(c *gin.Context) {
	var newLike models.Like
	if err := c.ShouldBindJSON(&newLike); err != nil {
		sendResponse(c, http.StatusBadRequest, "Invalid request body", nil, err.Error())
		return
	}

	userFound := false
	for _, user := range models.Users {
		if user.ID == newLike.UserID {
			userFound = true
			break
		}
	}
	if !userFound {
		sendResponse(c, http.StatusBadRequest, "Invalid UserID for like", nil, "Validation error")
		return
	}

	postFound := false
	for _, post := range models.Posts {
		if post.ID == newLike.PostID {
			postFound = true
			break
		}
	}
	if !postFound {
		sendResponse(c, http.StatusBadRequest, "Invalid PostID for like", nil, "Validation error")
		return
	}

	for _, like := range models.Likes {
		if like.UserID == newLike.UserID && like.PostID == newLike.PostID {
			sendResponse(c, http.StatusConflict, "User has already liked this post", nil, "Validation error")
			return
		}
	}

	newLike.ID = uuid.New().String()
	models.Likes = append(models.Likes, newLike)
	sendResponse(c, http.StatusCreated, "Like created successfully", newLike, nil)
}

// GetLikesByPostID handles GET /posts/:id/likes
func GetLikesByPostID(c *gin.Context) {
	postID := c.Param("id")
	var postLikes []models.Like
	for _, like := range models.Likes {
		if like.PostID == postID {
			postLikes = append(postLikes, like)
		}
	}
	
	postFound := false
	for _, post := range models.Posts {
		if post.ID == postID {
			postFound = true
			break
		}
	}
	if !postFound {
		sendResponse(c, http.StatusNotFound, "Post not found", nil, "Not Found")
		return
	}

	sendResponse(c, http.StatusOK, "Likes for post retrieved successfully", postLikes, nil)
}

// GetLikesByUserID handles GET /users/:id/likes
func GetLikesByUserID(c *gin.Context) {
	userID := c.Param("id")
	var userLikes []models.Like
	for _, like := range models.Likes {
		if like.UserID == userID {
			userLikes = append(userLikes, like)
		}
	}
	
	userFound := false
	for _, user := range models.Users {
		if user.ID == userID {
			userFound = true
			break
		}
	}
	if !userFound {
		sendResponse(c, http.StatusNotFound, "User not found", nil, "Not Found")
		return
	}

	sendResponse(c, http.StatusOK, "Likes by user retrieved successfully", userLikes, nil)
}