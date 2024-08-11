# Golang DataStore

This project is a Golang-based system that allows you to import data from an
Excel file, store it in a MySQL database, and cache the data in Redis.

## Software used

- **Go+**: The programming language used to build this application.
- **MySQL**: The relational database management system where data will be stored.
- **Redis**: The in-memory data structure store used for caching.

## Dependencies

This project relies on several Golang packages:

- **Gin**: A web framework used to handle HTTP requests
and route them to the appropriate handlers.
- **GORM**: An ORM library used for interacting with the MySQL database.
- **Redis**: A Redis client used for caching.
- **Excelize**: A library to parse Excel files.

You can install these dependencies by running `go mod tidy`.

## API Endpoints

1. **Import Excel Data**
    - **Endpoint**: POST `/import`
    - **Description**: Upload an Excel file to import data.
    - **Request**: Multipart form data with an Excel file.
    - **Response**:
        - 200 OK on successful import
        - 400 Bad Request if the file is not in the correct format
        - 500 Internal Server Error if any other error occurs during processing
2. **View Imported Data**
    - **Endpoint**: GET `/view`
    - **Description**: View the list of imported data.
    The data is fetched from Redis cache if available;
otherwise, it is retrieved from MySQL and then cached.
    - **Response**:
        - 200 OK with JSON array of records
3. **Edit a Record**
    - **Endpoint**: PUT `/edit/:id`
    - **Description**: Edit a specific record by ID.
    Updates both the MySQL database and the Redis cache.

    - **Request Body**:

    ```json
    {
        "FirstName": "Jane",
        "LastName": "Doe",
        "CompanyName": "Example Corp.",
        "Address": "5678 Oak St.",
        "City": "Shelbyville",
        "County": "Shelby",
        "Postal": "01104",
        "Phone": "555-5678",
        "Email": "jane.doe@example.com",
        "Web": "http://examplecorp.com"
    }
    ```

    - **Response**:

        - 200 OK on successful update
        - 400 Bad Request if the input data is invalid
        - 404 Not Found if the record with the given ID doesn't exist
