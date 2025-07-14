package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"fesnuk-api/models"
)

// CreateComment handles POST /comments
func CreateComment(c *gin.Context) {
	var newComment models.Comment
	if err := c.ShouldBindJSON(&newComment); err != nil {
		sendResponse(c, http.StatusBadRequest, "Invalid request body", nil, err.Error())
		return
	}

	if newComment.Content == "" {
		sendResponse(c, http.StatusBadRequest, "Comment content cannot be empty", nil, "Validation error")
		return
	}

	userFound := false
	for _, user := range models.Users {
		if user.ID == newComment.UserID {
			userFound = true
			break
		}
	}
	if !userFound {
		sendResponse(c, http.StatusBadRequest, "Invalid UserID for comment", nil, "Validation error")
		return
	}

	postFound := false
	for _, post := range models.Posts {
		if post.ID == newComment.PostID {
			postFound = true
			break
		}
	}
	if !postFound {
		sendResponse(c, http.StatusBadRequest, "Invalid PostID for comment", nil, "Validation error")
		return
	}

	newComment.ID = uuid.New().String()
	newComment.CreatedAt = time.Now().Format(time.RFC3339)
	models.Comments = append(models.Comments, newComment)
	sendResponse(c, http.StatusCreated, "Comment created successfully", newComment, nil)
}

// GetCommentsByPostID handles GET /posts/:id/comments
func GetCommentsByPostID(c *gin.Context) {
	postID := c.Param("id")
	var postComments []models.Comment

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

	for _, comment := range models.Comments {
		if comment.PostID == postID {
			postComments = append(postComments, comment)
		}
	}
	sendResponse(c, http.StatusOK, "Comments for post retrieved successfully", postComments, nil)
}