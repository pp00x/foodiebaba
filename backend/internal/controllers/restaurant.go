package controllers

import (
    "github.com/pp00x/foodiebaba/internal/db"
    "github.com/pp00x/foodiebaba/internal/models"
    "github.com/pp00x/foodiebaba/internal/utils"
    "github.com/pp00x/foodiebaba/pkg/logger"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// GetRestaurants godoc
// @Summary      List restaurants
// @Description  Get a list of approved restaurants with optional pagination, search, and filtering
// @Tags         Restaurants
// @Accept       json
// @Produce      json
// @Param        page     query     int    false  "Page number"
// @Param        limit    query     int    false  "Page size"
// @Param        name     query     string false  "Search by name"
// @Param        category query     string false  "Filter by category"
// @Success      200      {array}   models.Restaurant
// @Failure      500      {object}  map[string]string
// @Router       /restaurants [get]
func GetRestaurants(c *gin.Context) {
    var restaurants []models.Restaurant

    // Pagination parameters
    page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
    if err != nil || page < 1 {
        page = 1
    }
    limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
    if err != nil || limit < 1 {
        limit = 10
    }
    offset := (page - 1) * limit

    // Search and filter parameters
    name := c.Query("name")
    category := c.Query("category")

    // Build the query
    query := db.DB.Preload("Photos").Preload("Reviews").Where("status = ?", "approved")
    if name != "" {
        query = query.Where("name ILIKE ?", "%"+name+"%")
    }
    if category != "" {
        query = query.Where("category ILIKE ?", "%"+category+"%")
    }

    // Execute the query with pagination
    if err := query.Limit(limit).Offset(offset).Find(&restaurants).Error; err != nil {
        logger.Log.Error("Error fetching restaurants: ", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch restaurants"})
        return
    }

    c.JSON(http.StatusOK, restaurants)
}

// AddRestaurant godoc
// @Summary      Add a new restaurant
// @Description  Users can add a new restaurant listing (requires approval)
// @Tags         Restaurants
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        restaurant body     models.CreateRestaurantInput true "Restaurant"
// @Success      201       {object}  models.Restaurant
// @Failure      400       {object}  map[string]string
// @Failure      500       {object}  map[string]string
// @Router       /restaurants [post]
func AddRestaurant(c *gin.Context) {
    var input models.CreateRestaurantInput
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

    // Create a new Restaurant instance with the validated input and additional fields
    restaurant := models.Restaurant{
        Name:        input.Name,
        Address:     input.Address,
        Category:    input.Category,
        Description: input.Description,
        CreatedByID: userID,
        Status:      "pending",
    }

    if err := db.DB.Create(&restaurant).Error; err != nil {
        logger.Log.Error("Error adding restaurant: ", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add restaurant"})
        return
    }

    // Increment user reputation
    if err := db.DB.Model(&models.User{}).Where("id = ?", userID).Update("reputation", gorm.Expr("reputation + ?", 10)).Error; err != nil {
        logger.Log.Error("Error updating user reputation: ", err)
    }

    logger.Log.Infof("Restaurant added: %s by user %d", restaurant.Name, userID)
    c.JSON(http.StatusCreated, restaurant)
}

// UploadPhotos godoc
// @Summary      Upload photos for a restaurant
// @Description  Users can upload photos for a restaurant
// @Tags         Restaurants
// @Security     BearerAuth
// @Accept       multipart/form-data
// @Produce      json
// @Param        id      path      int                   true  "Restaurant ID"
// @Param        photos  formData  file                  true  "Photos"
// @Success      200     {array}   models.Photo
// @Failure      400     {object}  map[string]string
// @Failure      500     {object}  map[string]string
// @Router       /restaurants/{id}/photos [post]
func UploadPhotos(c *gin.Context) {
    restaurantID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        logger.Log.Error("Invalid restaurant ID: ", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid restaurant ID"})
        return
    }

    form, err := c.MultipartForm()
    if err != nil {
        logger.Log.Error("Error retrieving form data: ", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    files := form.File["photos"]
    var photos []models.Photo
    for _, file := range files {
        url, err := utils.UploadFile(file)
        if err != nil {
            logger.Log.Error("Error uploading file: ", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        photo := models.Photo{
            URL:          url,
            RestaurantID: uint(restaurantID),
        }
        photos = append(photos, photo)
    }
    if err := db.DB.Create(&photos).Error; err != nil {
        logger.Log.Error("Error saving photos: ", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, photos)
}