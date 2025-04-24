package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Tretorhate/university-management-system/internal/api/controllers"
	"github.com/Tretorhate/university-management-system/internal/api/middleware"
	"github.com/Tretorhate/university-management-system/internal/api/routes"
	"github.com/Tretorhate/university-management-system/internal/config"
	"github.com/Tretorhate/university-management-system/internal/domain"
	"github.com/Tretorhate/university-management-system/internal/repository"
	"github.com/Tretorhate/university-management-system/internal/service"
	"github.com/Tretorhate/university-management-system/pkg/auth"
	"github.com/Tretorhate/university-management-system/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	goValidator "github.com/go-playground/validator/v10" // Use alias to avoid conflict
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to database
	db, err := gorm.Open(postgres.Open(cfg.GetDBConnString()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate schema
	err = db.AutoMigrate(
		&domain.User{},
		&domain.Student{},
		&domain.Teacher{},
		&domain.Course{},
		&domain.Enrollment{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Register custom validations with Gin's validator
	if v, ok := binding.Validator.Engine().(*goValidator.Validate); ok {
		// Register custom validations
		_ = v.RegisterValidation("password", validator.ValidatePassword)
		_ = v.RegisterValidation("student_id", validator.ValidateStudentID)
		_ = v.RegisterValidation("employee_id", validator.ValidateEmployeeID)
		_ = v.RegisterValidation("course_code", validator.ValidateCourseCode)
		_ = v.RegisterValidation("date_range", validator.ValidateDateRange)
	}

	// Initialize repositories
	baseRepo := repository.NewRepository(db)
	userRepo := repository.NewUserRepository(baseRepo)
	studentRepo := repository.NewStudentRepository(baseRepo)
	teacherRepo := repository.NewTeacherRepository(baseRepo)
	courseRepo := repository.NewCourseRepository(baseRepo)
	enrollmentRepo := repository.NewEnrollmentRepository(baseRepo)

	// Initialize JWT service
	jwtService := auth.NewJWTService(cfg.JWTSecret)

	// Initialize services
	authService := service.NewAuthService(userRepo, jwtService)
	studentService := service.NewStudentService(studentRepo, userRepo)
	teacherService := service.NewTeacherService(teacherRepo, userRepo)
	courseService := service.NewCourseService(courseRepo, teacherRepo)
	enrollmentService := service.NewEnrollmentService(enrollmentRepo, studentRepo, courseRepo)

	// Initialize middleware
	authMiddleware := middleware.NewAuthMiddleware(jwtService)

	// Initialize controllers
	authController := controllers.NewAuthController(authService)
	studentController := controllers.NewStudentController(studentService)
	teacherController := controllers.NewTeacherController(teacherService)
	courseController := controllers.NewCourseController(courseService)

	// Check if the enrollment controller constructor exists
	// You'll need to implement this controller
	enrollmentController := controllers.NewEnrollmentController(enrollmentService)

	// Setup gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(
		router,
		authMiddleware,
		authController,
		studentController,
		teacherController,
		courseController,
		enrollmentController,
	)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.ServerPort
	}

	log.Printf("Server running on port %s", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
