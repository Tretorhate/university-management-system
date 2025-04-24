package validator

import (
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

// CustomValidator extends the default validator with custom validation functions
type CustomValidator struct {
	validator *validator.Validate
}

// New creates a new CustomValidator
func New() *CustomValidator {
	v := validator.New()

	// Register custom validations
	_ = v.RegisterValidation("password", ValidatePassword)
	_ = v.RegisterValidation("student_id", ValidateStudentID)
	_ = v.RegisterValidation("employee_id", ValidateEmployeeID)
	_ = v.RegisterValidation("course_code", ValidateCourseCode)
	_ = v.RegisterValidation("date_range", ValidateDateRange)

	return &CustomValidator{
		validator: v,
	}
}

// Validate validates a struct
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// ValidatePassword ensures password meets security requirements
// Export this function by capitalizing the first letter
func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// At least 8 characters, 1 uppercase, 1 lowercase, 1 number, 1 special character
	hasMinLen := len(password) >= 8
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password)

	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

// ValidateStudentID ensures student ID follows the required format
// Export this function by capitalizing the first letter
func ValidateStudentID(fl validator.FieldLevel) bool {
	studentID := fl.Field().String()

	// Format: YYYY-XXXXX where YYYY is year and XXXXX is a 5-digit number
	pattern := `^\d{4}-\d{5}$`
	return regexp.MustCompile(pattern).MatchString(studentID)
}

// ValidateEmployeeID ensures employee ID follows the required format
// Export this function by capitalizing the first letter
func ValidateEmployeeID(fl validator.FieldLevel) bool {
	employeeID := fl.Field().String()

	// Format: DEP-XXXXX where DEP is department code and XXXXX is a 5-digit number
	pattern := `^[A-Z]{3}-\d{5}$`
	return regexp.MustCompile(pattern).MatchString(employeeID)
}

// ValidateCourseCode ensures course code follows the required format
// Export this function by capitalizing the first letter
func ValidateCourseCode(fl validator.FieldLevel) bool {
	courseCode := fl.Field().String()

	// Format: SUBJ-XXX where SUBJ is subject code and XXX is a 3-digit number
	pattern := `^[A-Z]{4}-\d{3}$`
	return regexp.MustCompile(pattern).MatchString(courseCode)
}

// ValidateDateRange ensures end date is after start date
// Export this function by capitalizing the first letter
func ValidateDateRange(fl validator.FieldLevel) bool {
	// This is a struct-level validation, so we need to get the struct
	startDate := fl.Parent().FieldByName("StartDate").Interface().(time.Time)
	endDate := fl.Parent().FieldByName("EndDate").Interface().(time.Time)

	return endDate.After(startDate)
}
