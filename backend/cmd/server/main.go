package main

import (
    "github.com/pp00x/foodiebaba/configs"
    "github.com/pp00x/foodiebaba/internal/controllers"
    "github.com/pp00x/foodiebaba/internal/db"
    "github.com/pp00x/foodiebaba/internal/middlewares"
    "github.com/pp00x/foodiebaba/internal/models"
    "github.com/pp00x/foodiebaba/pkg/logger"
    "net/http"
     "time"

    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"

    // Swagger imports
    _ "github.com/pp00x/foodiebaba/docs"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"

    // Prometheus
    ginprometheus "github.com/zsais/go-gin-prometheus"
)

// @title           FoodieBaba API
// @version         1.0
// @description     API documentation for FoodieBaba application.
// @termsOfService  http://foodiebaba.com/terms/

// @contact.name   API Support
// @contact.url    http://foodiebaba.com/support
// @contact.email  support@foodiebaba.com

// @license.name  MIT License
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
    configs.LoadConfig()
    db.Init()

    // Initialize the logger
    logger.Init()
    err := db.DB.AutoMigrate(
        &models.User{},
        &models.Restaurant{},
        &models.Review{},
        &models.Photo{},
    )

    if err != nil {
        logger.Log.Fatal("Migration failed: ", err)
    }



    r := gin.Default()

    // Configure CORS
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))


    // Monitoring
    p := ginprometheus.NewPrometheus("gin")
    p.Use(r)

    // Serve static files
    r.Static("/uploads", "./uploads")

    // Swagger docs
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Public routes
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)
    r.GET("/restaurants", controllers.GetRestaurants)

    // Protected routes
    auth := r.Group("/")
    auth.Use(middlewares.JWTAuth())
    {
        auth.POST("/restaurants", controllers.AddRestaurant)
        auth.POST("/restaurants/:id/photos", controllers.UploadPhotos)
        auth.POST("/reviews", controllers.AddReview)
    }

    // Admin routes
    admin := r.Group("/admin")
    admin.Use(middlewares.JWTAuth(), middlewares.AdminOnly())
    {
        admin.GET("/restaurants/pending", controllers.GetPendingRestaurants)
        admin.PUT("/restaurants/:id/approve", controllers.ApproveRestaurant)
        admin.PUT("/restaurants/:id/reject", controllers.RejectRestaurant)
    }

    // Health check
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "OK"})
    })

    logger.Log.Info("Starting server on port 8080")
    r.Run()
}