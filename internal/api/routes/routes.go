package routes

import (
	"github.com/Tretorhate/university-management-system/internal/api/controllers"
	"github.com/Tretorhate/university-management-system/internal/api/middleware"
	"github.com/Tretorhate/university-management-system/internal/domain"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	authMiddleware *middleware.AuthMiddleware,
	authController *controllers.AuthController,
	studentController *controllers.StudentController,
	teacherController *controllers.TeacherController,
	courseController *controllers.CourseController,
	enrollmentController *controllers.EnrollmentController,
) {
	// Global middleware
	r.Use(middleware.ErrorHandler())

	// Public routes
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
	}

	// Protected routes
	api := r.Group("/api")
	api.Use(authMiddleware.AuthRequired())
	{
		// Students routes
		students := api.Group("/students")
		{
			students.POST("", authMiddleware.RoleRequired(domain.RoleAdmin), studentController.Create)
			students.GET("", studentController.GetAll)
			students.GET("/:id", studentController.GetByID)
			students.PUT("/:id", authMiddleware.RoleRequired(domain.RoleAdmin), studentController.Update)
			students.DELETE("/:id", authMiddleware.RoleRequired(domain.RoleAdmin), studentController.Delete)
		}

		// Teachers routes
		teachers := api.Group("/teachers")
		{
			teachers.POST("", authMiddleware.RoleRequired(domain.RoleAdmin), teacherController.Create)
			teachers.GET("", teacherController.GetAll)
			teachers.GET("/:id", teacherController.GetByID)
			teachers.PUT("/:id", authMiddleware.RoleRequired(domain.RoleAdmin), teacherController.Update)
			teachers.DELETE("/:id", authMiddleware.RoleRequired(domain.RoleAdmin), teacherController.Delete)
		}

		// Courses routes
		courses := api.Group("/courses")
		{
			courses.POST("", authMiddleware.RoleRequired(domain.RoleAdmin, domain.RoleTeacher), courseController.Create)
			courses.GET("", courseController.GetAll)
			courses.GET("/:id", courseController.GetByID)
			courses.PUT("/:id", authMiddleware.RoleRequired(domain.RoleAdmin, domain.RoleTeacher), courseController.Update)
			courses.DELETE("/:id", authMiddleware.RoleRequired(domain.RoleAdmin), courseController.Delete)
		}

		// Enrollments routes
		enrollments := api.Group("/enrollments")
		{
			enrollments.POST("", authMiddleware.RoleRequired(domain.RoleAdmin, domain.RoleTeacher), enrollmentController.Create)
			enrollments.GET("", enrollmentController.GetAll)
			enrollments.GET("/:id", enrollmentController.GetByID)
			enrollments.PUT("/:id", authMiddleware.RoleRequired(domain.RoleAdmin, domain.RoleTeacher), enrollmentController.Update)
			enrollments.DELETE("/:id", authMiddleware.RoleRequired(domain.RoleAdmin, domain.RoleTeacher), enrollmentController.Delete)
		}
	}
}
