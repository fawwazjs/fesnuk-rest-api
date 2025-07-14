package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"fesnuk-api/models"
)

// CreatePost handles POST /posts
func CreatePost(c *gin.Context) {
	var newPost models.Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		sendResponse(c, http.StatusBadRequest, "Invalid request body", nil, err.Error())
		return
	}

	if newPost.Content == "" {
		sendResponse(c, http.StatusBadRequest, "Post content cannot be empty", nil, "Validation error")
		return
	}

	userFound := false
	for _, user := range models.Users {
		if user.ID == newPost.UserID {
			userFound = true
			break
		}
	}
	if !userFound {
		sendResponse(c, http.StatusBadRequest, "Invalid UserID for post", nil, "Validation error")
		return
	}

	newPost.ID = uuid.New().String()
	newPost.CreatedAt = time.Now().Format(time.RFC3339)
	models.Posts = append(models.Posts, newPost)
	sendResponse(c, http.StatusCreated, "Post created successfully", newPost, nil)
}

// GetAllPosts handles GET /posts with optional keyword filtering
func GetAllPosts(c *gin.Context) {
	keyword := c.Query("keyword") // Mengambil query parameter 'keyword'

	var filteredPosts []models.Post
	if keyword == "" {
		filteredPosts = models.Posts // Jika g ada keyword, kembalikan semua post
	} else {
		lowerKeyword := strings.ToLower(keyword)
		for _, post := range models.Posts {
			if strings.Contains(strings.ToLower(post.Content), lowerKeyword) {
				filteredPosts = append(filteredPosts, post)
			}
		}
	}
	sendResponse(c, http.StatusOK, "Posts retrieved successfully", filteredPosts, nil)
}

// GetPostByID handles GET /posts/:id
func GetPostByID(c *gin.Context) {
	id := c.Param("id")
	for _, post := range models.Posts {
		if post.ID == id {
			sendResponse(c, http.StatusOK, "Post retrieved successfully", post, nil)
			return
		}
	}
	sendResponse(c, http.StatusNotFound, "Post not found", nil, "Not Found")
}

// GetPostsByUserID handles GET /users/:id/posts
func GetPostsByUserID(c *gin.Context) {
	userID := c.Param("id")
	var userPosts []models.Post
	for _, post := range models.Posts {
		if post.UserID == userID {
			userPosts = append(userPosts, post)
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

	sendResponse(c, http.StatusOK, "User posts retrieved successfully", userPosts, nil)
}

// DeletePost handles DELETE /posts/:id
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	for i, post := range models.Posts {
		if post.ID == id {
			models.Posts = append(models.Posts[:i], models.Posts[i+1:]...)
			// Delete like & komentar yg terkait dgn post
			var remainingLikes []models.Like
			for _, like := range models.Likes {
				if like.PostID != id {
					remainingLikes = append(remainingLikes, like)
				}
			}
			models.Likes = remainingLikes

			var remainingComments []models.Comment
			for _, comment := range models.Comments {
				if comment.PostID != id {
					remainingComments = append(remainingComments, comment)
				}
			}
			models.Comments = remainingComments

			sendResponse(c, http.StatusOK, "Post deleted successfully", nil, nil)
			return
		}
	}
	sendResponse(c, http.StatusNotFound, "Post not found", nil, "Not Found")
}