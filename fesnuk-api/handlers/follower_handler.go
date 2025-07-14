package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"fesnuk-api/models"
)

// FollowUser handles POST /users/:id/follow - user melakukan follow ke user lain
// :id ID user yang akan difollow
func FollowUser(c *gin.Context) {
	followingID := c.Param("id") // User yang akan difollow

	var requestBody struct {
		FollowerID string `json:"follower_id"` // User yang melakukan follow
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		sendResponse(c, http.StatusBadRequest, "Invalid request body", nil, err.Error())
		return
	}

	followerID := requestBody.FollowerID

	if followerID == followingID {
		sendResponse(c, http.StatusBadRequest, "User cannot follow themselves", nil, "Validation error")
		return
	}

	followerFound := false
	for _, user := range models.Users {
		if user.ID == followerID {
			followerFound = true
			break
		}
	}
	if !followerFound {
		sendResponse(c, http.StatusBadRequest, "Invalid FollowerID", nil, "Validation error")
		return
	}

	followingFound := false
	for _, user := range models.Users {
		if user.ID == followingID {
			followingFound = true
			break
		}
	}
	if !followingFound {
		sendResponse(c, http.StatusBadRequest, "Invalid FollowingID", nil, "Validation error")
		return
	}

	for _, f := range models.Followers {
		if f.FollowerID == followerID && f.FollowingID == followingID {
			sendResponse(c, http.StatusConflict, "User already follows this user", nil, "Validation error")
			return
		}
	}

	newFollow := models.Follower{
		ID:          uuid.New().String(),
		FollowerID:  followerID,
		FollowingID: followingID,
	}
	models.Followers = append(models.Followers, newFollow)
	sendResponse(c, http.StatusCreated, "User followed successfully", newFollow, nil)
}

// UnfollowUser handles DELETE /users/:id/unfollow - user berhenti follow user lain
// :id ID dari user yang tidak lagi difollow
func UnfollowUser(c *gin.Context) {
	unfollowingID := c.Param("id") // User yang tidak lagi difollow

	var requestBody struct {
		FollowerID string `json:"follower_id"` // User yang berhenti follow
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		sendResponse(c, http.StatusBadRequest, "Invalid request body", nil, err.Error())
		return
	}

	followerID := requestBody.FollowerID

	if followerID == unfollowingID {
		sendResponse(c, http.StatusBadRequest, "User cannot unfollow themselves", nil, "Validation error")
		return
	}

	found := false
	for i, f := range models.Followers {
		if f.FollowerID == followerID && f.FollowingID == unfollowingID {
			models.Followers = append(models.Followers[:i], models.Followers[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		sendResponse(c, http.StatusNotFound, "Follow relationship not found", nil, "Not Found")
		return
	}
	sendResponse(c, http.StatusOK, "User unfollowed successfully", nil, nil)
}

// GetFollowers handles GET /users/:id/followers - melihat daftar follower dari user tertentu
// :id ID user yang ingin dilihat followernya
func GetFollowers(c *gin.Context) {
	userID := c.Param("id")
	var followers []models.User

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

	for _, f := range models.Followers {
		if f.FollowingID == userID {
			for _, user := range models.Users {
				if user.ID == f.FollowerID {
					followers = append(followers, user)
					break
				}
			}
		}
	}
	sendResponse(c, http.StatusOK, "Followers retrieved successfully", followers, nil)
}

// GetFollowing handles GET /users/:id/following - melihat daftar user yang diikuti oleh user tertentu
// :id ID user yang ingin dilihat siapa yang dia ikuti
func GetFollowing(c *gin.Context) {
	userID := c.Param("id")
	var following []models.User

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

	for _, f := range models.Followers {
		if f.FollowerID == userID {
			for _, user := range models.Users {
				if user.ID == f.FollowingID {
					following = append(following, user)
					break
				}
			}
		}
	}
	sendResponse(c, http.StatusOK, "Following retrieved successfully", following, nil)
}