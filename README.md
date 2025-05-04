# Test-Task-For-Effective-Mobile
REST API for managing a list of people with support for filtering, pagination, and CRUD operations.
## Setup
1. Clone repository
```sh
git clone https://github.com/Iskhazov/Test-Task-For-Effective-Mobile.git
cd awesomeProject2
```
2. Install Dependencies
 ```sh
go mod tidy
```
3. Database Configuration  
* Set up PostgreSQL database.  
* Configure connection details in envs/.env
4. Run Migrations
 ```sh
make migrate-up
```
5. Start application
 ```sh
make run
```
## Swagger
Link to Swagger: [http://localhost:8080/swagger/index.html#/](http://localhost:8080/swagger/index.html#/)
## API Endpoints
POST /api/v1/person - Creates a new person.  
POST /api/v1/persons - Get all people.  
PUT /api/v1/persons/{id} - Update a person.  
DELETE /api/v1/persons/{id} - Delete a person.  

