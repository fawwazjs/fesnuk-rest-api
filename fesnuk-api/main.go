package main

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"fesnuk-api/handlers"
)

func main() {
	router := gin.Default()

	//Custom Logger Middleware untuk timestamping
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return param.TimeStamp.Format(time.RFC3339) + " | " + // Timestamp
			param.Latency.String() + " | " + // Latency
			param.ClientIP + " | " + // IP Klien
			param.Method + " " + // Metode HTTP
			param.Path + " " + // Path Request
			param.Request.Proto + " | " + // Protokol Request
			strconv.Itoa(param.StatusCode) + " | " + 
			param.ErrorMessage + "\n" // Pesan Error (jika ada)
	}))

	// User Endpoints
	router.POST("/users", handlers.CreateUser)
	router.GET("/users", handlers.GetAllUsers)
	router.GET("/users/:id", handlers.GetUserByID)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)

	// Post Endpoints
	router.POST("/posts", handlers.CreatePost)
	router.GET("/posts", handlers.GetAllPosts) // filter keyword
	router.GET("/posts/:id", handlers.GetPostByID)
	router.GET("/users/:id/posts", handlers.GetPostsByUserID)
	router.DELETE("/posts/:id", handlers.DeletePost)

	// Like Endpoints
	router.POST("/likes", handlers.CreateLike)
	router.GET("/posts/:id/likes", handlers.GetLikesByPostID)
	router.GET("/users/:id/likes", handlers.GetLikesByUserID)

	// Comment Endpoints
	router.POST("/comments", handlers.CreateComment)
	router.GET("/posts/:id/comments", handlers.GetCommentsByPostID)

	// Follower/Following Endpoints
	router.POST("/users/:id/follow", handlers.FollowUser)
	router.DELETE("/users/:id/unfollow", handlers.UnfollowUser)
	router.GET("/users/:id/followers", handlers.GetFollowers)
	router.GET("/users/:id/following", handlers.GetFollowing)

	log.Println("Fesnuk API Server running on port 8080")
	router.Run(":8080")
}