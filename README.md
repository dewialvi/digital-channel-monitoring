# Digital Channel Monitoring System

A backend-based monitoring system designed to simulate and manage digital channel operations, including API monitoring, transaction monitoring, bug reporting, user activity logging, and notifications.

This project was developed as a mini project to practice backend software engineering, RESTful API development, database design, authentication, monitoring, and API documentation using Go and PostgreSQL.

---

## Features

### Authentication & Authorization
- User registration
- User login
- JWT-based authentication
- Logout
- Role-based access control (Admin & Staff)

### Bug Report Management
- Create bug reports
- View bug report details
- List bug reports with pagination
- Search and filtering
- Bug severity and priority management
- Bug lifecycle status management
- Admin-only bug deletion

### API Monitoring
- Automatically record API request activity
- Record HTTP method
- Record HTTP status code
- Record response time
- Record request timestamp

### Transaction Monitoring
- Create transaction monitoring records
- Generate unique transaction identifiers
- Monitor transaction status
- Store customer MSISDN
- Store transaction type and amount
- List transaction records with pagination

### Activity Logging
- Record user activities
- Track bug creation
- Track bug updates
- Track bug deletion
- Store activity descriptions
- Store IP address information

### Notifications
- Create notifications
- Store notifications for users
- Track read/unread notification status

---

## Tech Stack

### Backend
- Go (Golang)
- Echo Framework
- GORM

### Database
- PostgreSQL

### Authentication & Security
- JWT (JSON Web Token)
- bcrypt password hashing
- Role-Based Access Control (RBAC)

### API Testing & Documentation
- Postman
- RESTful API
- Postman Collection
- Postman Environment

### Development Tools
- Visual Studio Code
- Git
- GitHub
- PowerShell

---

## Project Architecture

The project follows a layered architecture to separate application responsibilities.

```text
digital-channel-monitoring/
│
├── config/
│   ├── config.go
│   └── database.go
│
├── handler/
│   ├── auth_handler.go
│   ├── bug_report_handler.go
│   ├── api_monitor_handler.go
│   ├── transaction_monitor_handler.go
│   ├── activity_log_handler.go
│   └── notification_handler.go
│
├── middleware/
│   ├── jwt_middleware.go
│   ├── role_middleware.go
│   └── monitor_middleware.go
│
├── models/
│   ├── user.go
│   ├── bug_report.go
│   ├── api_monitor.go
│   ├── transaction_monitor.go
│   ├── activity_log.go
│   ├── notification.go
│   └── user_feedback.go
│
├── repository/
│   ├── user_repository.go
│   ├── bug_report_repository.go
│   ├── api_monitor_repository.go
│   ├── transaction_monitor_repository.go
│   ├── activity_log_repository.go
│   └── notification_repository.go
│
├── service/
│   ├── auth_service.go
│   ├── bug_report_service.go
│   ├── activity_log_service.go
│   └── notification_service.go
│
├── routes/
│   └── routes.go
│
├── docs/
│   ├── API.md
│   ├── DATABASE.md
│   └── postman/
│       ├── Digital-Channel-Monitoring.postman_collection.json
│       └── Digital-Channel-Local.postman_environment.json
│
├── main.go
├── go.mod
├── go.sum
└── README.md