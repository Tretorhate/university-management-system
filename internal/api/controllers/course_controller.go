package controllers

import (
	"net/http"
	"strconv"

	"github.com/Tretorhate/university-management-system/internal/dto"
	"github.com/Tretorhate/university-management-system/internal/service"
	"github.com/gin-gonic/gin"
)

type CourseController struct {
    courseService *service.CourseService
}

func NewCourseController(courseService *service.CourseService) *CourseController {
    return &CourseController{courseService: courseService}
}

func (c *CourseController) Create(ctx *gin.Context) {
    var request dto.CourseCreateDTO
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    response, err := c.courseService.Create(&request)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, response)
}
func (c *CourseController) GetAll(ctx *gin.Context) {
    courses, err := c.courseService.GetAll(ctx.Query("search"))
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, courses)
}

func (c *CourseController) GetByID(ctx *gin.Context) {
    id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    course, err := c.courseService.GetByID(uint(id))
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
        return
    }

    ctx.JSON(http.StatusOK, course)
}

func (c *CourseController) Update(ctx *gin.Context) {
    id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    var request dto.CourseUpdateDTO
    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    course, err := c.courseService.Update(uint(id), &request)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, course)
}

func (c *CourseController) Delete(ctx *gin.Context) {
    id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    if err := c.courseService.Delete(uint(id)); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}