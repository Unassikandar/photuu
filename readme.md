# Photuu

Photuu is a simple web application built with Go and the Gin framework. It provides basic user authentication, session management, and CSRF protection. The application provides CRUD api for user posts

## Features

- User registration and login
- Session-based authentication
- CSRF token validation
- RESTful API endpoints

## Prerequisites

- Go 1.23.2 or higher
- A terminal or command-line interface
- A web browser or HTTP client (e.g., Postman, curl)

## Setup

### 1. Clone the Repository

```bash
git clone https://github.com/unassikandar/photuu.git
cd photuu
```

### 2. Configure the Environment Variables

Create a .env file in the root directory of the project and add the following environment variables:

```bash
OAUTH_CLIENT_ID=your-google-oauth-client-id
OAUTH_CLIENT_SECRET=your-google-oauth-client-secret
OAUTH_REDIRECT_URL=http://localhost:8080
```

### 3. Install Dependencies

Run the following command to install the required dependencies:

```bash
go mod tidy
```

### 4. Run the Application

Start the application by running:

```bash
go run main.go
```

The server will start on http://localhost:8080.