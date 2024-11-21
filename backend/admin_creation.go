package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "github.com/pp00x/foodiebaba/configs"
    "github.com/pp00x/foodiebaba/internal/db"
    "github.com/pp00x/foodiebaba/internal/models"
    "golang.org/x/crypto/bcrypt"
)

func main() {
    configs.LoadConfig()
    db.Init()

    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter admin username: ")
    username, _ := reader.ReadString('\n')
    username = strings.TrimSpace(username)

    fmt.Print("Enter admin email: ")
    email, _ := reader.ReadString('\n')
    email = strings.TrimSpace(email)

    fmt.Print("Enter admin password: ")
    password, _ := reader.ReadString('\n')
    password = strings.TrimSpace(password)

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        fmt.Println("Error hashing password:", err)
        os.Exit(1)
    }

    adminUser := models.User{
        Username: username,
        Email:    email,
        Password: string(hashedPassword),
        Role:     "admin",
    }

    if err := db.DB.Create(&adminUser).Error; err != nil {
        fmt.Println("Error creating admin user:", err)
        os.Exit(1)
    }

    fmt.Println("Admin user created successfully")
}