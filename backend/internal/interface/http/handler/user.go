package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	// TODO: Add use case dependencies
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(useCase interface{}) *UserHandler {
	return &UserHandler{}
}

// Register handles user registration
// @Summary Register a new user
// @Description Register a new user account with email, password, and nickname
// @Tags auth
// @Accept json
// @Produce json
// @Param request body object{email=string,password=string,nickname=string} true "Registration request"
// @Success 201 {object} object{message=string,user=object{id=uint,email=string,nickname=string,created_at=string}}
// @Failure 400 {object} object{error=string}
// @Failure 409 {object} object{error=string} "Email or nickname already exists"
// @Router /auth/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Register endpoint - TODO"})
}

// Login handles user login
// @Summary User login
// @Description Authenticate user and return access/refresh tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param request body object{email=string,password=string} true "Login credentials"
// @Success 200 {object} object{access_token=string,refresh_token=string,user=object{id=uint,email=string,nickname=string,is_admin=bool}}
// @Failure 400 {object} object{error=string}
// @Failure 401 {object} object{error=string} "Invalid credentials"
// @Router /auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Login endpoint - TODO"})
}

// RefreshToken handles token refresh
// @Summary Refresh access token
// @Description Get a new access token using refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body object{refresh_token=string} true "Refresh token"
// @Success 200 {object} object{access_token=string}
// @Failure 400 {object} object{error=string}
// @Failure 401 {object} object{error=string} "Invalid or expired refresh token"
// @Router /auth/refresh [post]
func (h *UserHandler) RefreshToken(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "RefreshToken endpoint - TODO"})
}

// Logout handles user logout
// @Summary User logout
// @Description Invalidate refresh token and logout user
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} object{message=string}
// @Failure 401 {object} object{error=string}
// @Router /auth/logout [post]
func (h *UserHandler) Logout(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Logout endpoint - TODO"})
}

// GetProfile retrieves user profile
// @Summary Get user profile
// @Description Get current authenticated user's profile information
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} object{user=object{id=uint,email=string,nickname=string,avatar_url=string,bio=string,is_admin=bool,created_at=string}}
// @Failure 401 {object} object{error=string}
// @Router /auth/me [get]
func (h *UserHandler) GetProfile(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "GetProfile endpoint - TODO"})
}

// UpdateProfile updates user profile
// @Summary Update user profile
// @Description Update current user's profile (nickname, avatar_url, bio)
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body object{nickname=string,avatar_url=string,bio=string} true "Profile update request"
// @Success 200 {object} object{message=string,user=object{id=uint,email=string,nickname=string,avatar_url=string,bio=string}}
// @Failure 400 {object} object{error=string}
// @Failure 401 {object} object{error=string}
// @Failure 409 {object} object{error=string} "Nickname already exists"
// @Router /auth/me [put]
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "UpdateProfile endpoint - TODO"})
}
