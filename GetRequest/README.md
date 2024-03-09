# README

## Introduction
This Java program fetches comments data from the [JSONPlaceholder](https://jsonplaceholder.typicode.com/) API and stores it into a MySQL database. It prompts the user to specify the number of comments to retrieve, then makes HTTP GET requests to the API and parses the JSON response to extract comments data. Finally, it inserts the comments data into the MySQL database.

## Requirements
- Java Development Kit (JDK)
- MySQL database server
- MySQL JDBC driver
- JSONPlaceholder API (for fetching comments data)

## Usage
1. Ensure that you have a MySQL database set up.
2. Compile the Java code using the JDK.
3. Run the compiled Java program.
4. Follow the prompts to specify the number of comments to retrieve.
5. The program will fetch comments data from the JSONPlaceholder API and store it in the MySQL database.

## Configuration
- Update the MySQL connection details (`url`, `username`, `password`) in the code to match your database configuration.

## Error Handling
- The program handles exceptions gracefully and prints stack traces in case of errors during execution.

## Notes
- This program demonstrates basic HTTP GET request handling, JSON parsing, and JDBC usage in Java.
- Make sure to handle any potential rate limits or restrictions imposed by the JSONPlaceholder API.
