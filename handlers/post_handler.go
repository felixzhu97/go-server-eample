package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-server-example/database"
	"go-server-example/models"
	"go-server-example/utils/logger"
)

// PostHandler 文章处理器
type PostHandler struct{}

// NewPostHandler 创建文章处理器实例
func NewPostHandler() *PostHandler {
	return &PostHandler{}
}

// GetPosts 获取文章列表
func (h *PostHandler) GetPosts(c *gin.Context) {
	var posts []models.Post
	
	// 预加载用户信息
	result := database.GetDB().Preload("User").Find(&posts)
	if result.Error != nil {
		logger.Error("Failed to get posts:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get posts",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data": posts,
		"total": len(posts),
	})
}

// GetPost 获取单个文章
func (h *PostHandler) GetPost(c *gin.Context) {
	id := c.Param("id")
	postID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid post ID",
		})
		return
	}
	
	var post models.Post
	result := database.GetDB().Preload("User").First(&post, postID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}

// CreatePost 创建文章
func (h *PostHandler) CreatePost(c *gin.Context) {
	var post models.Post
	
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
		return
	}
	
	// 从请求中获取用户ID（实际应用中应该从JWT token中获取）
	userID := c.GetUint("user_id")
	if userID == 0 {
		userID = 1 // 默认用户ID，实际应用中应该从认证中间件获取
	}
	post.UserID = userID
	
	result := database.GetDB().Create(&post)
	if result.Error != nil {
		logger.Error("Failed to create post:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create post",
		})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"data": post,
		"message": "Post created successfully",
	})
}

// UpdatePost 更新文章
func (h *PostHandler) UpdatePost(c *gin.Context) {
	id := c.Param("id")
	postID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid post ID",
		})
		return
	}
	
	var post models.Post
	if err := database.GetDB().First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}
	
	var updateData models.Post
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
		return
	}
	
	result := database.GetDB().Model(&post).Updates(updateData)
	if result.Error != nil {
		logger.Error("Failed to update post:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update post",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data": post,
		"message": "Post updated successfully",
	})
}

// DeletePost 删除文章
func (h *PostHandler) DeletePost(c *gin.Context) {
	id := c.Param("id")
	postID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid post ID",
		})
		return
	}
	
	var post models.Post
	if err := database.GetDB().First(&post, postID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}
	
	result := database.GetDB().Delete(&post)
	if result.Error != nil {
		logger.Error("Failed to delete post:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete post",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
	})
}

// GetPostsByUser 获取用户的文章
func (h *PostHandler) GetPostsByUser(c *gin.Context) {
	userID := c.Param("user_id")
	uid, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}
	
	var posts []models.Post
	result := database.GetDB().Where("user_id = ?", uid).Find(&posts)
	if result.Error != nil {
		logger.Error("Failed to get user posts:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user posts",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data": posts,
		"total": len(posts),
	})
} 