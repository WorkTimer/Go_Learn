package routes

import (
	"blog-system/internal/handlers"
	"blog-system/internal/middleware"
	"blog-system/pkg/logger"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, jwtSecret string) {
	userHandler := handlers.NewUserHandler()
	postHandler := handlers.NewPostHandler()
	commentHandler := handlers.NewCommentHandler()

	r.Use(middleware.CORSMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", userHandler.Register)
			auth.POST("/login", userHandler.Login)
		}

		authenticated := v1.Group("")
		authenticated.Use(middleware.AuthMiddleware(jwtSecret))
		{
			authenticated.GET("/profile", userHandler.GetProfile)

			posts := authenticated.Group("/posts")
			{
				posts.POST("", postHandler.CreatePost)
				posts.PUT("/:id", postHandler.UpdatePost)
				posts.DELETE("/:id", postHandler.DeletePost)
			}

			comments := authenticated.Group("/comments")
			{
				comments.POST("", commentHandler.CreateComment)
				comments.DELETE("/:id", commentHandler.DeleteComment)
			}
		}

		public := v1.Group("")
		public.Use(middleware.OptionalAuthMiddleware(jwtSecret))
		{
			public.GET("/posts", postHandler.GetPosts)
			public.GET("/posts/:id", postHandler.GetPost)

			public.GET("/posts/:id/comments", commentHandler.GetCommentsByPostID)
		}
	}

	logger.Info("Routes setup completed")
}
