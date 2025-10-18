package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/viblog/internal/infrastructure/auth"
	"github.com/yourusername/viblog/internal/interface/http/dto"
	"github.com/yourusername/viblog/internal/usecase/user"
	"github.com/yourusername/viblog/pkg/errors"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	registerUseCase      *user.RegisterUseCase
	loginUseCase         *user.LoginUseCase
	getProfileUseCase    *user.GetProfileUseCase
	updateProfileUseCase *user.UpdateProfileUseCase
	jwtService           *auth.JWTService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(
	registerUseCase *user.RegisterUseCase,
	loginUseCase *user.LoginUseCase,
	getProfileUseCase *user.GetProfileUseCase,
	updateProfileUseCase *user.UpdateProfileUseCase,
	jwtService *auth.JWTService,
) *UserHandler {
	return &UserHandler{
		registerUseCase:      registerUseCase,
		loginUseCase:         loginUseCase,
		getProfileUseCase:    getProfileUseCase,
		updateProfileUseCase: updateProfileUseCase,
		jwtService:           jwtService,
	}
}

// Register handles user registration
// @Summary Register a new user
// @Description Register a new user account with email, password, and nickname
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Registration request"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} object{error=string}
// @Failure 409 {object} object{error=string} "Email or nickname already exists"
// @Router /auth/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Execute registration use case
	newUser, err := h.registerUseCase.Execute(c.Request.Context(), user.RegisterInput{
		Email:    req.Email,
		Password: req.Password,
		Nickname: req.Nickname,
	})

	if err != nil {
		statusCode := errors.GetStatusCode(err)
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	// Convert to response DTO
	userResp := dto.UserResponse{
		ID:        newUser.ID,
		Email:     newUser.Email,
		Nickname:  newUser.Nickname,
		AvatarURL: newUser.AvatarURL,
		Bio:       newUser.Bio,
		IsAdmin:   newUser.IsAdmin,
		CreatedAt: newUser.CreatedAt,
	}

	c.JSON(http.StatusCreated, userResp)
}

// Login handles user login
// @Summary User login
// @Description Authenticate user and return access/refresh tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login credentials"
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} object{error=string}
// @Failure 401 {object} object{error=string} "Invalid credentials"
// @Router /auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Execute login use case
	authenticatedUser, err := h.loginUseCase.Execute(c.Request.Context(), user.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		statusCode := errors.GetStatusCode(err)
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	// Generate tokens
	accessToken, err := h.jwtService.GenerateAccessToken(authenticatedUser.ID, authenticatedUser.Email, authenticatedUser.IsAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	refreshToken, err := h.jwtService.GenerateRefreshToken(authenticatedUser.ID, authenticatedUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	// Convert to response DTO
	authResp := dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: dto.UserResponse{
			ID:          authenticatedUser.ID,
			Email:       authenticatedUser.Email,
			Nickname:    authenticatedUser.Nickname,
			AvatarURL:   authenticatedUser.AvatarURL,
			Bio:         authenticatedUser.Bio,
			IsAdmin:     authenticatedUser.IsAdmin,
			CreatedAt:   authenticatedUser.CreatedAt,
			LastLoginAt: authenticatedUser.LastLoginAt,
		},
	}

	c.JSON(http.StatusOK, authResp)
}

// RefreshToken handles token refresh
// @Summary Refresh access token
// @Description Get a new access token using refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.RefreshTokenRequest true "Refresh token"
// @Success 200 {object} dto.TokenResponse
// @Failure 400 {object} object{error=string}
// @Failure 401 {object} object{error=string} "Invalid or expired refresh token"
// @Router /auth/refresh [post]
func (h *UserHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate refresh token
	claims, err := h.jwtService.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	// Generate new access token
	accessToken, err := h.jwtService.GenerateAccessToken(claims.UserID, claims.Email, claims.IsAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	c.JSON(http.StatusOK, dto.TokenResponse{
		AccessToken: accessToken,
	})
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
	// In a stateless JWT setup, logout is handled client-side by removing tokens
	// In production, you might want to implement token blacklisting
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// GetProfile retrieves user profile
// @Summary Get user profile
// @Description Get current authenticated user's profile information
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.UserResponse
// @Failure 401 {object} object{error=string}
// @Router /auth/me [get]
func (h *UserHandler) GetProfile(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Execute get profile use case
	userProfile, err := h.getProfileUseCase.Execute(c.Request.Context(), user.GetProfileInput{
		UserID: userID.(uint),
	})

	if err != nil {
		statusCode := errors.GetStatusCode(err)
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	// Convert to response DTO
	userResp := dto.UserResponse{
		ID:          userProfile.ID,
		Email:       userProfile.Email,
		Nickname:    userProfile.Nickname,
		AvatarURL:   userProfile.AvatarURL,
		Bio:         userProfile.Bio,
		IsAdmin:     userProfile.IsAdmin,
		CreatedAt:   userProfile.CreatedAt,
		LastLoginAt: userProfile.LastLoginAt,
	}

	c.JSON(http.StatusOK, userResp)
}

// UpdateProfile updates user profile
// @Summary Update user profile
// @Description Update current user's profile (nickname, avatar_url, bio)
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.UpdateProfileRequest true "Profile update request"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} object{error=string}
// @Failure 401 {object} object{error=string}
// @Failure 409 {object} object{error=string} "Nickname already exists"
// @Router /auth/me [put]
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	// Get user ID from context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Execute update profile use case
	updatedUser, err := h.updateProfileUseCase.Execute(c.Request.Context(), user.UpdateProfileInput{
		UserID:    userID.(uint),
		Nickname:  req.Nickname,
		AvatarURL: req.AvatarURL,
		Bio:       req.Bio,
	})

	if err != nil {
		statusCode := errors.GetStatusCode(err)
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	// Convert to response DTO
	userResp := dto.UserResponse{
		ID:          updatedUser.ID,
		Email:       updatedUser.Email,
		Nickname:    updatedUser.Nickname,
		AvatarURL:   updatedUser.AvatarURL,
		Bio:         updatedUser.Bio,
		IsAdmin:     updatedUser.IsAdmin,
		CreatedAt:   updatedUser.CreatedAt,
		LastLoginAt: updatedUser.LastLoginAt,
	}

	c.JSON(http.StatusOK, userResp)
}
