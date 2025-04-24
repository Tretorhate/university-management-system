// internal/api/controllers/teacher_controller.go
package controllers

import (
	"net/http"
	"strconv"

	"github.com/Tretorhate/university-management-system/internal/dto"
	"github.com/Tretorhate/university-management-system/internal/service"
	"github.com/gin-gonic/gin"
)

type TeacherController struct {
	teacherService *service.TeacherService
}

func NewTeacherController(teacherService *service.TeacherService) *TeacherController {
	return &TeacherController{teacherService: teacherService}
}

func (c *TeacherController) Create(ctx *gin.Context) {
	var request dto.TeacherCreateDTO
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.teacherService.Create(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (c *TeacherController) GetAll(ctx *gin.Context) {
	teachers, err := c.teacherService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, teachers)
}

func (c *TeacherController) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	teacher, err := c.teacherService.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
		return
	}

	ctx.JSON(http.StatusOK, teacher)
}

func (c *TeacherController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var request dto.TeacherUpdateDTO
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teacher, err := c.teacherService.Update(uint(id), &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, teacher)
}

func (c *TeacherController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := c.teacherService.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Teacher deleted successfully"})
}
