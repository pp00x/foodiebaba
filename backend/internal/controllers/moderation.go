package controllers

import (
    "github.com/pp00x/foodiebaba/internal/db"
    "github.com/pp00x/foodiebaba/internal/models"
    "github.com/pp00x/foodiebaba/pkg/logger"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

// GetPendingRestaurants godoc
// @Summary      Get pending restaurants
// @Description  Admins can get a list of restaurants pending approval
// @Tags         Moderation
// @Security     BearerAuth
// @Produce      json
// @Success      200  {array}   models.Restaurant
// @Failure      500  {object}  map[string]string
// @Router       /admin/restaurants/pending [get]
func GetPendingRestaurants(c *gin.Context) {
    var restaurants []models.Restaurant
    if err := db.DB.Where("status = ?", "pending").Find(&restaurants).Error; err != nil {
        logger.Log.Error("Error fetching pending restaurants: ", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch pending restaurants"})
        return
    }
    c.JSON(http.StatusOK, restaurants)
}

// ApproveRestaurant godoc
// @Summary      Approve a restaurant
// @Description  Admins can approve a pending restaurant
// @Tags         Moderation
// @Security     BearerAuth
// @Param        id   path      int  true  "Restaurant ID"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /admin/restaurants/{id}/approve [put]
func ApproveRestaurant(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        logger.Log.Error("Invalid restaurant ID: ", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid restaurant ID"})
        return
    }
    if err := db.DB.Model(&models.Restaurant{}).Where("id = ?", id).Update("status", "approved").Error; err != nil {
        logger.Log.Error("Error approving restaurant: ", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to approve restaurant"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Restaurant approved"})
}

// RejectRestaurant godoc
// @Summary      Reject a restaurant
// @Description  Admins can reject a pending restaurant
// @Tags         Moderation
// @Security     BearerAuth
// @Param        id   path      int  true  "Restaurant ID"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /admin/restaurants/{id}/reject [put]
func RejectRestaurant(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        logger.Log.Error("Invalid restaurant ID: ", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid restaurant ID"})
        return
    }
    if err := db.DB.Model(&models.Restaurant{}).Where("id = ?", id).Update("status", "rejected").Error; err != nil {
        logger.Log.Error("Error rejecting restaurant: ", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reject restaurant"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Restaurant rejected"})
}