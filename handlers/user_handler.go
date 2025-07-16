package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-server-example/database"
	"go-server-example/models"
	"go-server-example/utils/logger"
)

// UserHandler 用户处理器
type UserHandler struct{}

// NewUserHandler 创建用户处理器实例
func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// GetUsers 获取用户列表
func (h *UserHandler) GetUsers(c *gin.Context) {
	var users []models.User
	
	result := database.GetDB().Find(&users)
	if result.Error != nil {
		logger.Error("Failed to get users:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get users",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data": users,
		"total": len(users),
	})
}

// GetUser 获取单个用户
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}
	
	var user models.User
	result := database.GetDB().First(&user, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// CreateUser 创建用户
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
		return
	}
	
	result := database.GetDB().Create(&user)
	if result.Error != nil {
		logger.Error("Failed to create user:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"data": user,
		"message": "User created successfully",
	})
}

// UpdateUser 更新用户
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}
	
	var user models.User
	if err := database.GetDB().First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}
	
	var updateData models.User
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
		return
	}
	
	result := database.GetDB().Model(&user).Updates(updateData)
	if result.Error != nil {
		logger.Error("Failed to update user:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data": user,
		"message": "User updated successfully",
	})
}

// DeleteUser 删除用户
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}
	
	var user models.User
	if err := database.GetDB().First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}
	
	result := database.GetDB().Delete(&user)
	if result.Error != nil {
		logger.Error("Failed to delete user:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete user",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
} 