package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Tretorhate/university-management-system/internal/api/controllers"
	"github.com/Tretorhate/university-management-system/internal/api/middleware"
	"github.com/Tretorhate/university-management-system/internal/api/routes"
	"github.com/Tretorhate/university-management-system/internal/config"
	"github.com/Tretorhate/university-management-system/internal/repository"
	"github.com/Tretorhate/university-management-system/internal/service"
	"github.com/Tretorhate/university-management-system/pkg/auth"
	"github.com/Tretorhate/university-management-system/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	goValidator "github.com/go-playground/validator/v10"
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

	// Note: AutoMigrate is disabled as schema is managed by golang-migrate
	// If you need to verify the schema matches your models, you can add a check here
	// For example, ensure tables exist, but do not modify the schema

	// Register custom validations with Gin's validator
	if v, ok := binding.Validator.Engine().(*goValidator.Validate); ok {
		if err := v.RegisterValidation("password", validator.ValidatePassword); err != nil {
			log.Fatalf("Failed to register password validation: %v", err)
		}
		if err := v.RegisterValidation("student_id", validator.ValidateStudentID); err != nil {
			log.Fatalf("Failed to register student_id validation: %v", err)
		}
		if err := v.RegisterValidation("employee_id", validator.ValidateEmployeeID); err != nil {
			log.Fatalf("Failed to register employee_id validation: %v", err)
		}
		if err := v.RegisterValidation("course_code", validator.ValidateCourseCode); err != nil {
			log.Fatalf("Failed to register course_code validation: %v", err)
		}
		if err := v.RegisterValidation("date_range", validator.ValidateDateRange); err != nil {
			log.Fatalf("Failed to register date_range validation: %v", err)
		}
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
