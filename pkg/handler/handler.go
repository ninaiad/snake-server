package handler

import (
	"log"
	"net/http"
	"os"
	"snake/pkg/service"

	"github.com/gin-gonic/gin"

	cors "github.com/rs/cors/wrapper/gin"
)

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type updateScoreInput struct {
	Score	 uint64	`json:"score" binding:"required"`
}

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) SetupRouter() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{os.Getenv("GAME_CLIENT_ORIGIN")},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowCredentials: true,
		AllowedHeaders: []string{"[Authorization]", "Authorization", "Content-Type", "Bearer", "Bearer ", "content-type", "Origin", "Accept"},
		OptionsPassthrough: true,
		// Enable Debugging for testing, consider disabling in production
		// Debug: true,
	}))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.DELETE("/delete", h.userIdentity, h.deleteUser)
	}

	scores := router.Group("/scores", h.userIdentity)
	{
		scores.GET("/", h.getAllScores)
		scores.GET("/:id", h.getScore)
		scores.POST("/:id", h.updateScore)
	}

	log.Println("router set up");

	return router
}
