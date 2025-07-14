package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"fesnuk-api/models"
)

// CreateUser handles POST /users
func CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		sendResponse(c, http.StatusBadRequest, "Invalid request body", nil, err.Error())
		return
	}

	if newUser.Username == "" || newUser.Email == "" {
		sendResponse(c, http.StatusBadRequest, "Username and email cannot be empty", nil, "Validation error")
		return
	}

	for _, user := range models.Users {
		if user.Username == newUser.Username {
			sendResponse(c, http.StatusBadRequest, "Username already exists", nil, "Validation error")
			return
		}
		if user.Email == newUser.Email {
			sendResponse(c, http.StatusBadRequest, "Email already exists", nil, "Validation error")
			return
		}
	}

	newUser.ID = uuid.New().String()
	models.Users = append(models.Users, newUser)
	sendResponse(c, http.StatusCreated, "User created successfully", newUser, nil)
}

// GetAllUsers handles GET /users
func GetAllUsers(c *gin.Context) {
	sendResponse(c, http.StatusOK, "Users retrieved successfully", models.Users, nil)
}

// GetUserByID handles GET /users/:id
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	for _, user := range models.Users {
		if user.ID == id {
			sendResponse(c, http.StatusOK, "User retrieved successfully", user, nil)
			return
		}
	}
	sendResponse(c, http.StatusNotFound, "User not found", nil, "Not Found")
}

// UpdateUser handles PUT /users/:id
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		sendResponse(c, http.StatusBadRequest, "Invalid request body", nil, err.Error())
		return
	}

	if updatedUser.Username == "" || updatedUser.Email == "" {
		sendResponse(c, http.StatusBadRequest, "Username and email cannot be empty", nil, "Validation error")
		return
	}

	for i, user := range models.Users {
		if user.ID == id {
			for j, existingUser := range models.Users {
				if i != j && (existingUser.Username == updatedUser.Username || existingUser.Email == updatedUser.Email) {
					sendResponse(c, http.StatusBadRequest, "Username or email already exists", nil, "Validation error")
					return
				}
			}

			models.Users[i].Username = updatedUser.Username
			models.Users[i].Email = updatedUser.Email
			models.Users[i].Bio = updatedUser.Bio
			sendResponse(c, http.StatusOK, "User updated successfully", models.Users[i], nil)
			return
		}
	}
	sendResponse(c, http.StatusNotFound, "User not found", nil, "Not Found")
}

// DeleteUser handles DELETE /users/:id
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	for i, user := range models.Users {
		if user.ID == id {
			models.Users = append(models.Users[:i], models.Users[i+1:]...)
			// delete post, like, dan komentar yang terkait dengan user
			var remainingPosts []models.Post
			for _, post := range models.Posts {
				if post.UserID != id {
					remainingPosts = append(remainingPosts, post)
				}
			}
			models.Posts = remainingPosts

			var remainingLikes []models.Like
			for _, like := range models.Likes {
				if like.UserID != id {
					remainingLikes = append(remainingLikes, like)
				}
			}
			models.Likes = remainingLikes

			var remainingComments []models.Comment
			for _, comment := range models.Comments {
				if comment.UserID != id {
					remainingComments = append(remainingComments, comment)
				}
			}
			models.Comments = remainingComments

			var remainingFollowers []models.Follower
			for _, f := range models.Followers {
				if f.FollowerID != id && f.FollowingID != id {
					remainingFollowers = append(remainingFollowers, f)
				}
			}
			models.Followers = remainingFollowers


			sendResponse(c, http.StatusOK, "User deleted successfully", nil, nil)
			return
		}
	}
	sendResponse(c, http.StatusNotFound, "User not found", nil, "Not Found")
}