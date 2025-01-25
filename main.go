package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"github.com/songwaad/cs-activity-backend/adapters"
	"github.com/songwaad/cs-activity-backend/config"
	_ "github.com/songwaad/cs-activity-backend/docs"
	"github.com/songwaad/cs-activity-backend/entities"
	"github.com/songwaad/cs-activity-backend/middleware"
	"github.com/songwaad/cs-activity-backend/usecases"
)

func main() {
	// โหลดไฟล์ .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// ดึง SECRET_KEY จาก environment
	secretKey := os.Getenv("JWT_SECRET_KEY")

	dbHost := os.Getenv("DB_HOST")
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Invalid DB_PORT value")
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// เชื่อมต่อฐานข้อมูล
	db := config.LoadDatabase(dbHost, dbPort, dbUser, dbPassword, dbName)

	db.AutoMigrate(&entities.User{})

	// Initialize repository and use case
	userRepo := &adapters.GormUserRepo{DB: db}
	userUseCase := &usecases.UserUseCase{UserRepo: userRepo}
	userHandler := &adapters.UserHandler{UserUseCase: userUseCase}

	// Setup Fiber
	app := fiber.New()

	// เปิดใช้งาน CORS Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",       // อนุญาตเฉพาะ Origin ของ Frontend
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS", // อนุญาต Methods ที่ใช้
		AllowHeaders:     "Content-Type, Authorization", // อนุญาต Headers ที่ต้องการ
		AllowCredentials: true,                          // อนุญาตการใช้ Cookie
	}))

	app.Post("/register", userHandler.Register)
	app.Post("/login", func(c *fiber.Ctx) error {
		// Pass the secret key to the login handler
		return userHandler.Login(c, secretKey)
	})

	// Swagger route
	app.Get("/swagger/*", swagger.HandlerDefault)

	// JWT Middleware (ใช้กับทุกเส้นทางที่ต้องการตรวจสอบ JWT)
	app.Use(middleware.AuthMiddleware())

	// Protected Route สำหรับ "admin" เท่านั้น
	app.Get("/admin", middleware.RoleMiddleware("admin"), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome, admin!",
		})
	})

	// Protected Route สำหรับ "user" เท่านั้น
	app.Get("/user", middleware.RoleMiddleware("user"), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome, user!",
		})
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	app.Listen(":8080")
}
