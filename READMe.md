# Aquafarm Management Application
#### By Abdul Salam

This project aims to develop a prototype of an aquafarm management application. The application manages two entities: farms and ponds. Farms can have multiple ponds, but each pond can only be registered to one farm.

## Problem Statement

The application should fulfill the following requirements:

1. Handle a POST request to create a farm or a pond. Duplicate entries should be denied.
2. Handle a PUT request to update an existing farm or pond, or create a new one if the specified entity doesn't exist.
3. Handle a DELETE request to soft delete an existing farm or pond. Throw an error if the specified entity doesn't exist. Note that hard delete (complete data deletion) should not be used.
4. Handle a GET request to retrieve a list of all currently existing farms or ponds in the system. If no entities are found, return the HTTP status 404 Not Found.
5. Handle a GET request for a specific farm or pond using its ID as a path parameter. If the specified ID doesn't exist, return the HTTP status 404 Not Found. Additionally, ensure the correct relationship between ponds and farms is reflected in this API.

## Technology Stack

This application is developed using the following technologies:

- Go programming language (version 1.20)
- SQLite database

## Note on Database Design

Please note that for the purpose of this prototype, I have made the following design decisions:

- To prioritize performance optimization, I have intentionally avoided using joins and foreign key constraints between the farms and ponds tables in this prototype. While this simplifies the implementation, it's important to note that in a production system, establishing proper relationship constraints and utilizing joins and foreign keys would be recommended for maintaining data integrity and enforcing referential integrity. However, for the purpose of this prototype, omitting these features helps to enhance the application's performance.
- To enhance database performance, I recommend setting indexes for several fields that are commonly used in queries. By setting indexes on these fields, you can speed up data retrieval operations. Consider indexing fields such as farmID, pondID, and any other frequently used columns.

## Instructions

To run the application, make sure you have Go 1.20 and SQLite installed on your system. Follow the steps below:

1. Clone the repository:

   ```bash
   git clone https://github.com/abdulsalam01/delos-farm-go.git
   ```

2. Navigate to the project directory:

   ```bash
   cd delos-farm-go
   ```

3. Build and run the application:

   ```bash
   go run cmd/app/main.go
   ```

4. Or simply (if using docker):
    ```
    docker-compose up
    ```

   This will start the application, and it will be accessible at `http://localhost:3000`.


## Documentation

Documentation that useful for run and tests the API appropriately is attached below:

1. Attached Postman Collection here that can be accessed on:
    ```
    https://api.postman.com/collections/1535500-6fe0a9f1-59f0-4d18-b01a-9ae550110467?access_key=PMAT-01H3YPJ3GQETTM05KSPCH1KPQV
    ```

2. There's also swagger documentation that can be found on root project named:
    ```
    open-api.yaml
    ```

## API Endpoints

The following API endpoints are available:

- **POST /farm** - Create a new farm
- **POST /pond** - Create a new pond
- **PUT /farm/{farmID}** - Update an existing farm or create a new one
- **PUT /pond/{pondID}** - Update an existing pond or create a new one
- **DELETE /farm/{farmID}** - Soft delete an existing farm
- **DELETE /pond/{pondID}** - Soft delete an existing pond
- **GET /farm** - Retrieve a list of all farms
- **GET /pond** - Retrieve a list of all ponds
- **GET /farm/{farmID}** - Retrieve a specific farm by ID
- **GET /pond/{pondID}** - Retrieve a specific pond by ID

Please note that the appropriate request payloads and responses should be provided when interacting with these endpoints.

---

With the provided instructions, you can set up and run the aquafarm management application using Go 1.20 and SQLite. Feel free to modify and extend the application as needed. If you have any questions or need further assistance, please don't hesitate to ask me at @abdulsalam.
