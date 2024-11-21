package controllers

import (
    "github.com/pp00x/foodiebaba/internal/db"
    "github.com/pp00x/foodiebaba/internal/models"
    "github.com/pp00x/foodiebaba/internal/utils"
    "github.com/pp00x/foodiebaba/pkg/logger"
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// AddReview godoc
// @Summary      Add a review
// @Description  Users can add a review to a restaurant
// @Tags         Reviews
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        review body     models.Review true "Review"
// @Success      201    {object} models.Review
// @Failure      400    {object} map[string]string
// @Failure      500    {object} map[string]string
// @Router       /reviews [post]
func AddReview(c *gin.Context) {
    var input models.Review
    if err := c.ShouldBindJSON(&input); err != nil {
        logger.Log.Error("Invalid input: ", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Input validation
    if err := utils.Validate.Struct(input); err != nil {
        logger.Log.Error("Validation error: ", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := c.GetUint("userID")
    input.UserID = userID

    if err := db.DB.Create(&input).Error; err != nil {
        logger.Log.Error("Error adding review: ", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add review"})
        return
    }

    // Increment user reputation
    if err := db.DB.Model(&models.User{}).Where("id = ?", userID).Update("reputation", gorm.Expr("reputation + ?", 5)).Error; err != nil {
        logger.Log.Error("Error updating user reputation: ", err)
    }

    c.JSON(http.StatusCreated, input)
}