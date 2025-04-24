# University Management System (UMS)

A RESTful API for university management built with Golang.

## Project Overview

The University Management System (UMS) is a comprehensive solution for managing university data including students, teachers, courses, and enrollments. It features role-based authentication and authorization using JWT tokens and provides a complete set of CRUD operations for all entities.

## Technologies Used

- Golang 1.21+
- Gin Web Framework
- GORM (with PostgreSQL driver)
- JWT for Authentication
- Docker and Docker Compose

## Design Patterns Implemented

- **DTO (Data Transfer Object)**: Separates internal models from external views
- **Factory Pattern**: Creates entities or DTOs from incoming requests
- **Strategy Pattern**: Different course sorting strategies (by date, by student count, by name)
- **Service Layer**: Clear separation of business logic

## Project Structure

```
university-management-system/
├── cmd/
│   └── api/
│       └── main.go         # Application entry point
├── internal/
│   ├── api/                # HTTP handlers
│   │   ├── controllers/    # Route handlers
│   │   ├── middleware/     # Auth middleware, logging, etc.
│   │   └── routes/         # Route definitions
│   ├── config/             # Configuration management
│   ├── domain/             # Domain models/entities
│   ├── dto/                # Data Transfer Objects
│   ├── repository/         # Database interaction
│   ├── service/            # Business logic
│   │   ├── factory/        # Factory pattern implementations
│   │   └── strategy/       # Strategy pattern implementations
│   └── util/               # Utility functions
├── pkg/                    # Reusable packages
│   ├── auth/               # Authentication utilities
│   └── validator/          # Custom validators
├── migrations/             # SQL migration files
├── tests/                  # Test files
├── go.mod                  # Dependencies
├── go.sum                  # Dependency checksums
├── .env                    # Environment variables (not committed)
├── .env.example            # Example environment variables
├── Dockerfile              # Docker configuration
├── docker-compose.yml      # Docker compose configuration
└── README.md               # Project documentation
```

## How to Run the Project

### Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose (for containerized deployment)
- PostgreSQL (if running without Docker)

### Setup and Run

1. Clone the repository

   ```bash
   git clone https://github.com/yourusername/university-management-system.git
   cd university-management-system
   ```

2. Set up environment variables

   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. Run with Docker

   ```bash
   docker-compose up --build
   ```

4. Or run locally
   ```bash
   go mod download
   go run cmd/api/main.go
   ```

## API Endpoints

### Authentication

- `POST /auth/register` - Register a new user
- `POST /auth/login` - Login and get JWT token

### Students

- `POST /api/students` - Create a student (ADMIN)
- `GET /api/students` - List all students (All roles)
- `GET /api/students/:id` - Get student by ID (All roles)
- `PUT /api/students/:id` - Update student (ADMIN)
- `DELETE /api/students/:id` - Delete student (ADMIN)

### Teachers

- `POST /api/teachers` - Create a teacher (ADMIN)
- `GET /api/teachers` - List all teachers (All roles)
- `GET /api/teachers/:id` - Get teacher by ID (All roles)
- `PUT /api/teachers/:id` - Update teacher (ADMIN)
- `DELETE /api/teachers/:id` - Delete teacher (ADMIN)

### Courses

- `POST /api/courses` - Create a course (ADMIN, TEACHER)
- `GET /api/courses` - List all courses (All roles)
- `GET /api/courses/:id` - Get course by ID (All roles)
- `PUT /api/courses/:id` - Update course (ADMIN, TEACHER)
- `DELETE /api/courses/:id` - Delete course (ADMIN)

### Enrollments

- `POST /api/enrollments` - Create enrollment (ADMIN, TEACHER)
- `GET /api/enrollments` - List all enrollments (All roles)
- `GET /api/enrollments/:id` - Get enrollment by ID (All roles)
- `PUT /api/enrollments/:id` - Update enrollment (ADMIN, TEACHER)
- `DELETE /api/enrollments/:id` - Delete enrollment (ADMIN, TEACHER)

## Authentication

All protected endpoints require a valid JWT token in the Authorization header:

```
Authorization: Bearer {your_jwt_token}
```

## License

This project is licensed under the MIT License.

## Contact Information

For any questions or suggestions, please contact:

- Email: tretorhate@gmail.com | tretorhate@outlook.com
- GitHub: [Your GitHub Username](https://github.com/Tretorhate)
