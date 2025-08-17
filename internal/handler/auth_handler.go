package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	v1 "github.com/nitinjangam/pos-receipt-system/api/v1"
	"github.com/nitinjangam/pos-receipt-system/internal/service"
	"go.uber.org/zap"
)

type AuthHandlerInterface interface {
	PostAuthLogin(c *gin.Context)
	PostAuthRegister(c *gin.Context)
}

type AuthHandler struct {
	authService service.AuthServiceInterface
	logger      *zap.SugaredLogger
}

func NewAuthHandler(authService service.AuthServiceInterface, logger *zap.SugaredLogger) AuthHandlerInterface {
	return &AuthHandler{
		authService: authService,
		logger:      logger,
	}
}

func (s *AuthHandler) PostAuthLogin(c *gin.Context) {
	user := v1.PostAuthLoginJSONRequestBody{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}
	if user.Username == nil || user.Password == nil {
		c.JSON(400, gin.H{"message": "Username and password are required"})
		return
	}
	if err := s.authService.Login(c.Request.Context(), *user.Username, *user.Password); err != nil {
		c.JSON(401, gin.H{"message": "Unauthorized"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"username": user.Username,
		},
	})

	s.logger.Infow("User logged in", "username", user.Username)
}

func (s *AuthHandler) PostAuthRegister(c *gin.Context) {
	user := v1.PostAuthRegisterJSONRequestBody{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}
	fmt.Println(user)
	if user.Username == nil || user.Password == nil {
		c.JSON(400, gin.H{"message": "Username and password are required"})
		return
	}
	if err := s.authService.Register(c.Request.Context(), *user.Username, *user.Password); err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(201, gin.H{
		"message": "Registration successful",
		"user": gin.H{
			"username": user.Username,
		},
	})

	s.logger.Infow("User registered", "username", user.Username)
}
