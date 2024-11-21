package controllers

import (
    "github.com/pp00x/foodiebaba/internal/db"
    "github.com/pp00x/foodiebaba/internal/models"
    "github.com/pp00x/foodiebaba/internal/utils"
    "github.com/pp00x/foodiebaba/pkg/logger"
    "net/http"
    "os"
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

// Register godoc
// @Summary      Register a new user
// @Description  Register a new user with username, email, and password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user body     models.User true "User"
// @Success      201  {object} map[string]string
// @Failure      400  {object} map[string]string
// @Router       /register [post]
func Register(c *gin.Context) {
    var input models.User
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

    // Hash the password using bcrypt
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
    if err != nil {
        logger.Log.Error("Error hashing password: ", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
        return
    }

    // Create a new user instance with the hashed password
    user := models.User{
        Username: input.Username,
        Email:    input.Email,
        Password: string(hashedPassword),
        Role:     "user",
    }

    // Save the user to the database
    if err := db.DB.Create(&user).Error; err != nil {
        logger.Log.Error("Error creating user: ", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email or username already exists"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login godoc
// @Summary      User login
// @Description  Authenticate user and return a JWT token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        credentials body     map[string]string true "Email and Password"
// @Success      200  {object} map[string]string
// @Failure      400  {object} map[string]string
// @Failure      401  {object} map[string]string
// @Router       /login [post]
func Login(c *gin.Context) {
    var input struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        logger.Log.Error("Invalid input: ", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := db.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        logger.Log.Warnf("User not found: %s", input.Email)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        logger.Log.Warnf("Invalid password for user: %s", input.Email)
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "role":    user.Role,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })
    tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

    logger.Log.Infof("User logged in: %s", user.Email)
    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}