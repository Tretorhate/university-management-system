package controllers

import (
	"strconv"

	"github.com/Tretorhate/university-management-system/internal/dto"
	"github.com/Tretorhate/university-management-system/internal/service"
	"github.com/Tretorhate/university-management-system/pkg/errors"
	"github.com/gin-gonic/gin"
)

type EnrollmentController struct {
	enrollmentService *service.EnrollmentService
}

func NewEnrollmentController(enrollmentService *service.EnrollmentService) *EnrollmentController {
	return &EnrollmentController{enrollmentService: enrollmentService}
}

func (c *EnrollmentController) Create(ctx *gin.Context) {
	var request dto.EnrollmentCreateDTO
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.Error(errors.BadRequest("Invalid request body", err))
		return
	}

	response, err := c.enrollmentService.Create(&request)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(201, response)
}

func (c *EnrollmentController) GetAll(ctx *gin.Context) {
	// Check for optional query parameters
	studentID := ctx.Query("studentId")
	courseID := ctx.Query("courseId")

	var studentIDUint, courseIDUint uint
	var err error

	if studentID != "" {
		id, err := strconv.ParseUint(studentID, 10, 32)
		if err != nil {
			ctx.Error(errors.BadRequest("Invalid student ID format", err))
			return
		}
		studentIDUint = uint(id)
	}

	if courseID != "" {
		id, err := strconv.ParseUint(courseID, 10, 32)
		if err != nil {
			ctx.Error(errors.BadRequest("Invalid course ID format", err))
			return
		}
		courseIDUint = uint(id)
	}

	var enrollments []dto.EnrollmentResponseDTO

	if studentID != "" && courseID != "" {
		// Both studentID and courseID are provided
		enrollments, err = c.enrollmentService.GetByStudentAndCourseID(studentIDUint, courseIDUint)
	} else if studentID != "" {
		// Only studentID is provided
		enrollments, err = c.enrollmentService.GetByStudentID(studentIDUint)
	} else if courseID != "" {
		// Only courseID is provided
		enrollments, err = c.enrollmentService.GetByCourseID(courseIDUint)
	} else {
		// No filters provided, get all enrollments
		enrollments, err = c.enrollmentService.GetAll()
	}

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, enrollments)
}

func (c *EnrollmentController) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.Error(errors.BadRequest("Invalid ID format", err))
		return
	}

	enrollment, err := c.enrollmentService.GetByID(uint(id))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, enrollment)
}

func (c *EnrollmentController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.Error(errors.BadRequest("Invalid ID format", err))
		return
	}

	var request dto.EnrollmentUpdateDTO
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.Error(errors.BadRequest("Invalid request body", err))
		return
	}

	enrollment, err := c.enrollmentService.Update(uint(id), &request)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, enrollment)
}

func (c *EnrollmentController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.Error(errors.BadRequest("Invalid ID format", err))
		return
	}

	if err := c.enrollmentService.Delete(uint(id)); err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, gin.H{"message": "Enrollment deleted successfully"})
}
