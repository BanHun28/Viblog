package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/viblog/internal/interface/http/dto"
	"github.com/yourusername/viblog/internal/interface/http/presenter"
	"github.com/yourusername/viblog/internal/usecase/admin"
)

// AdminHandler handles admin-related HTTP requests
type AdminHandler struct {
	dashboardUC       *admin.GetDashboardUseCase
	listUsersUC       *admin.ListUsersUseCase
	deleteUserUC      *admin.DeleteUserUseCase
	listCommentsUC    *admin.ListCommentsUseCase
	deleteCommentUC   *admin.DeleteCommentUseCase
	listCategoriesUC  *admin.ListCategoriesUseCase
	createCategoryUC  *admin.CreateCategoryUseCase
	updateCategoryUC  *admin.UpdateCategoryUseCase
	deleteCategoryUC  *admin.DeleteCategoryUseCase
	listTagsUC        *admin.ListTagsUseCase
	createTagUC       *admin.CreateTagUseCase
	updateTagUC       *admin.UpdateTagUseCase
	deleteTagUC       *admin.DeleteTagUseCase
}

// NewAdminHandler creates a new AdminHandler
func NewAdminHandler(
	dashboardUC *admin.GetDashboardUseCase,
	listUsersUC *admin.ListUsersUseCase,
	deleteUserUC *admin.DeleteUserUseCase,
	listCommentsUC *admin.ListCommentsUseCase,
	deleteCommentUC *admin.DeleteCommentUseCase,
	listCategoriesUC *admin.ListCategoriesUseCase,
	createCategoryUC *admin.CreateCategoryUseCase,
	updateCategoryUC *admin.UpdateCategoryUseCase,
	deleteCategoryUC *admin.DeleteCategoryUseCase,
	listTagsUC *admin.ListTagsUseCase,
	createTagUC *admin.CreateTagUseCase,
	updateTagUC *admin.UpdateTagUseCase,
	deleteTagUC *admin.DeleteTagUseCase,
) *AdminHandler {
	return &AdminHandler{
		dashboardUC:      dashboardUC,
		listUsersUC:      listUsersUC,
		deleteUserUC:     deleteUserUC,
		listCommentsUC:   listCommentsUC,
		deleteCommentUC:  deleteCommentUC,
		listCategoriesUC: listCategoriesUC,
		createCategoryUC: createCategoryUC,
		updateCategoryUC: updateCategoryUC,
		deleteCategoryUC: deleteCategoryUC,
		listTagsUC:       listTagsUC,
		createTagUC:      createTagUC,
		updateTagUC:      updateTagUC,
		deleteTagUC:      deleteTagUC,
	}
}

// GetDashboard retrieves dashboard data
// @Summary Get admin dashboard
// @Description Get dashboard statistics and metrics (Admin only)
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.AdminDashboardResponse
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /admin/dashboard [get]
func (h *AdminHandler) GetDashboard(c *gin.Context) {
	stats, err := h.dashboardUC.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve dashboard statistics"})
		return
	}

	response := presenter.PresentDashboardStats(stats)
	c.JSON(http.StatusOK, response)
}

// ListUsers lists all users
// @Summary List all users
// @Description Get paginated list of all users (Admin only)
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(20)
// @Success 200 {object} dto.AdminUsersListResponse
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /admin/users [get]
func (h *AdminHandler) ListUsers(c *gin.Context) {
	// Parse pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	users, total, err := h.listUsersUC.Execute(c.Request.Context(), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	response := presenter.PresentUsersList(users, total, page, limit)
	c.JSON(http.StatusOK, response)
}

// DeleteUser deletes a user
// @Summary Delete user
// @Description Delete a user account (Admin only)
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} dto.MessageResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /admin/users/{id} [delete]
func (h *AdminHandler) DeleteUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.deleteUserUC.Execute(c.Request.Context(), uint(userID)); err != nil {
		if err == admin.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		if err == admin.ErrCannotDeleteAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Cannot delete admin user"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, dto.MessageResponse{Message: "User deleted successfully"})
}

// ListComments lists all comments
// @Summary List all comments
// @Description Get paginated list of all comments (Admin only)
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(20)
// @Success 200 {object} dto.AdminCommentsListResponse
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /admin/comments [get]
func (h *AdminHandler) ListComments(c *gin.Context) {
	// Parse pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	comments, total, err := h.listCommentsUC.Execute(c.Request.Context(), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comments"})
		return
	}

	response := presenter.PresentCommentsList(comments, total, page, limit)
	c.JSON(http.StatusOK, response)
}

// DeleteComment deletes a comment
// @Summary Delete comment
// @Description Delete a comment (Admin only)
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Comment ID"
// @Success 200 {object} dto.MessageResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /admin/comments/{id} [delete]
func (h *AdminHandler) DeleteComment(c *gin.Context) {
	commentID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	if err := h.deleteCommentUC.Execute(c.Request.Context(), uint(commentID)); err != nil {
		if err == admin.ErrCommentNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, dto.MessageResponse{Message: "Comment deleted successfully"})
}

// ListCategories lists all categories
// @Summary List all categories (Admin)
// @Description Get all categories for management (Admin only)
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.CategoriesListResponse
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /admin/categories [get]
func (h *AdminHandler) ListCategories(c *gin.Context) {
	categories, err := h.listCategoriesUC.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve categories"})
		return
	}

	response := presenter.PresentCategoriesList(categories)
	c.JSON(http.StatusOK, response)
}

// CreateCategory creates a new category
// @Summary Create category
// @Description Create a new blog category (Admin only)
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateCategoryRequest true "Category creation request"
// @Success 201 {object} dto.CategoryResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 409 {object} map[string]interface{}
// @Router /admin/categories [post]
func (h *AdminHandler) CreateCategory(c *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := admin.CreateCategoryInput{
		Name:        req.Name,
		Description: req.Description,
	}

	category, err := h.createCategoryUC.Execute(c.Request.Context(), input)
	if err != nil {
		if err == admin.ErrCategoryNameRequired {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Category name is required"})
			return
		}
		if err == admin.ErrCategoryNameExists || err == admin.ErrCategorySlugExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	response := presenter.PresentCategory(category)
	c.JSON(http.StatusCreated, response)
}

// UpdateCategory updates a category
// @Summary Update category
// @Description Update an existing category (Admin only)
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Param request body dto.UpdateCategoryRequest true "Category update request"
// @Success 200 {object} dto.CategoryResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 409 {object} map[string]interface{}
// @Router /admin/categories/{id} [put]
func (h *AdminHandler) UpdateCategory(c *gin.Context) {
	categoryID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var req dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := admin.UpdateCategoryInput{
		Name:        req.Name,
		Description: req.Description,
	}

	category, err := h.updateCategoryUC.Execute(c.Request.Context(), uint(categoryID), input)
	if err != nil {
		if err == admin.ErrCategoryNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}
		if err == admin.ErrCategoryNameExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	response := presenter.PresentCategory(category)
	c.JSON(http.StatusOK, response)
}

// DeleteCategory deletes a category
// @Summary Delete category
// @Description Delete a category (Admin only)
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Success 200 {object} dto.MessageResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /admin/categories/{id} [delete]
func (h *AdminHandler) DeleteCategory(c *gin.Context) {
	categoryID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	if err := h.deleteCategoryUC.Execute(c.Request.Context(), uint(categoryID)); err != nil {
		if err == admin.ErrCategoryNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	c.JSON(http.StatusOK, dto.MessageResponse{Message: "Category deleted successfully"})
}

// ListTags lists all tags
// @Summary List all tags (Admin)
// @Description Get all tags for management (Admin only)
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dto.TagsListResponse
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /admin/tags [get]
func (h *AdminHandler) ListTags(c *gin.Context) {
	tags, err := h.listTagsUC.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tags"})
		return
	}

	response := presenter.PresentTagsList(tags)
	c.JSON(http.StatusOK, response)
}

// CreateTag creates a new tag
// @Summary Create tag
// @Description Create a new blog tag (Admin only)
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateTagRequest true "Tag creation request"
// @Success 201 {object} dto.TagResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 409 {object} map[string]interface{}
// @Router /admin/tags [post]
func (h *AdminHandler) CreateTag(c *gin.Context) {
	var req dto.CreateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := admin.CreateTagInput{
		Name: req.Name,
	}

	tag, err := h.createTagUC.Execute(c.Request.Context(), input)
	if err != nil {
		if err == admin.ErrTagNameRequired {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Tag name is required"})
			return
		}
		if err == admin.ErrTagNameExists || err == admin.ErrTagSlugExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tag"})
		return
	}

	response := presenter.PresentTag(tag)
	c.JSON(http.StatusCreated, response)
}

// UpdateTag updates a tag
// @Summary Update tag
// @Description Update an existing tag (Admin only)
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Tag ID"
// @Param request body dto.UpdateTagRequest true "Tag update request"
// @Success 200 {object} dto.TagResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 409 {object} map[string]interface{}
// @Router /admin/tags/{id} [put]
func (h *AdminHandler) UpdateTag(c *gin.Context) {
	tagID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tag ID"})
		return
	}

	var req dto.UpdateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := admin.UpdateTagInput{
		Name: req.Name,
	}

	tag, err := h.updateTagUC.Execute(c.Request.Context(), uint(tagID), input)
	if err != nil {
		if err == admin.ErrTagNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
			return
		}
		if err == admin.ErrTagNameExists {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tag"})
		return
	}

	response := presenter.PresentTag(tag)
	c.JSON(http.StatusOK, response)
}

// DeleteTag deletes a tag
// @Summary Delete tag
// @Description Delete a tag (Admin only)
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Tag ID"
// @Success 200 {object} dto.MessageResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /admin/tags/{id} [delete]
func (h *AdminHandler) DeleteTag(c *gin.Context) {
	tagID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tag ID"})
		return
	}

	if err := h.deleteTagUC.Execute(c.Request.Context(), uint(tagID)); err != nil {
		if err == admin.ErrTagNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tag"})
		return
	}

	c.JSON(http.StatusOK, dto.MessageResponse{Message: "Tag deleted successfully"})
}
