# Digital Channel Monitoring System

A backend-focused monitoring system built with Go (Golang), Echo Framework, PostgreSQL, and GORM.

## Features

- User Authentication
- JWT Authentication
- Role-Based Access Control (RBAC)
- Bug Report Management
- Bug Status Workflow
- Pagination & Filtering
- API Monitoring
- Transaction Monitoring
- Activity Logging
- Notification Service
- Asynchronous API Monitoring Logging
- Dashboard Frontend

## Tech Stack

- Go (Golang)
- Echo Framework
- PostgreSQL
- GORM
- JWT
- REST API
- HTML
- CSS
- JavaScript
- Git & GitHub
- Postman

## Project Architecture

The project follows a layered architecture:

- Handler
- Service
- Repository
- Middleware
- Models
- Routes

## Monitoring

The system automatically records API monitoring data including:

- Endpoint
- HTTP Method
- Status Code
- Response Time
- Timestamp

The monitoring logs are stored asynchronously using Go Goroutines to avoid blocking the main request flow.

## Bug Report

The system supports:

- Create Bug Report
- Get Bug Report
- Get All Bug Reports
- Search & Filtering
- Pagination
- Update Bug Status
- Delete Bug Report (Admin)

## Transaction Monitoring

Transaction monitoring records:

- Transaction ID
- MSISDN
- Transaction Type
- Amount
- Status
- Created At

## Activity Logging

The system records user activities such as:

- Create Bug Report
- Update Bug Report
- Delete Bug Report

## Authentication

Authentication uses JWT with role-based authorization.

Available roles:

- Admin
- User

## Running Locally

Clone the repository:

```bash
git clone https://github.com/dewialvi/digital-channel-monitoring.git