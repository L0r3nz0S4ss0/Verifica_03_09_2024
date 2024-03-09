# README

## Introduction
This project is a simple RESTful API implemented in Go using the Gin web framework and MySQL database. It provides endpoints to perform CRUD operations on a comments table in the database.

## Requirements
- Go programming language (version 1.16 or later)
- MySQL database server
- Gin (Gin-Gonic) web framework
- Go MySQL driver

## Installation
1. Install Go from the [official website](https://golang.org/doc/install).
2. Install Gin and Go MySQL driver using the following `go get` commands:
   - `go get -u github.com/gin-gonic/gin`
   - `go get -u github.com/go-sql-driver/mysql`
3. Ensure MySQL server is installed and running.

## Configuration
- Make sure to update the MySQL connection string in `main()` function to connect to your MySQL server:
  ```go
  db, err = sql.Open("mysql", "username:password@tcp(hostname:port)/databasename")
  ```

## Usage
1. Run the Go program:
   ```go
   go run main.go
   ```
2. Access the API endpoints using a REST client such as cURL or Postman.

## API Endpoints

### 1. Get Comments by Date
- **Endpoint:** `/comments/api/get/date/:date`
- **Method:** GET
- **Parameters:**
- `:date` - Date in the format 'YYYY-MM-DD'
- **Description:** Retrieves comments added on the specified date.
- **Example:** `/comments/api/get/date/2024-03-08`

### 2. Get Comments by Time
- **Endpoint:** `/comments/api/get/time/:time`
- **Method:** GET
- **Parameters:**
- `:time` - Time in 24-hour format 'HHMM'
- **Description:** Retrieves comments added at the specified time.
- **Example:** `/comments/api/get/time/1430`

### 3. Update Comment
- **Endpoint:** `/comments/api/update/:id`
- **Method:** PATCH
- **Parameters:**
- `:id` - ID of the comment to be updated
- **Request Body:** JSON object containing the fields to be updated (`name`, `email`, `body`)
- **Description:** Updates the specified comment with new data.
- **Example Request Body:**
```json
{
 "name": "New Name",
 "email": "new@email.com",
 "body": "Updated comment body"
}
```
4. Delete Comment

- **Endpoint:** `/comments/api/delete/:id`
- **Method:** DELETE
- **Parameters:**
  - `:id` - ID of the comment to be deleted
- **Description:** Deletes the specified comment from the database.

### 5. Create Comment

- **Endpoint:** `/comments/api/create`
- **Method:** POST
- **Request Body:** JSON object containing comment data (`postId`, `id`, `name`, `email`, `body`)
- **Description:** Creates a new comment in the database.
- **Example Request Body:**
  ```json
  {
    "postId": "1",
    "id": "101",
    "name": "John Doe",
    "email": "john@example.com",
    "body": "This is a new comment"
  }
  ```
