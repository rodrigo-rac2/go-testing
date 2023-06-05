# rodrigo_sandbox

## Description

## Go Testing

Purpose: test end functions (not integrations)

### Unit Testing

Run all tests in the current directory:
`` go test ``

Coverage in CLI:
`` go test -cover ``

Coverage in html format:
`` go test -coverprofile=coverage.out && go tool cover -html=coverage.out ``

# Integration/API Testing

Located at `go-restful-jw-api`

Instructions:

1. Clone the repo
2. Cd into `cd go-restful-jw-api`
3. Next, install the project dependencies using the following command: `go mod tidy`
4. Create a new Postgres database called `diary_app`: 
    
    ```
   createdb -h localhost -p <DB_PORT> -U <DB_USER> diary_app
   ```
   Where prompted, provide the password associated with DB_USER. 
      
   Alternatively, create the database using your preferred tool of choice.
5. Next, make a copy of .env.example named .env.local using the following command.

    ```
   cp .env.example .env.local
   ```
   
### Set up the test environment
Testify is one of the most popular testing packages for Golang. 

1. Install it using the following command.

```
go get github.com/stretchr/testify
```

2. Next, create a new database named diary_app_test.

```
createdb -h localhost -p <DB_PORT> -U <DB_USER> diary_app_test --password
```

3. With that done, create a .env.test.local file from the .env.local file using the command below.

```
cp .env.local .env.test.local
```

Doing this will provide a separate, local test environment configuration, preventing unexpected behavior.

5. Run the API tests housed in the ``controller`` folder:

```
go test -v -cover ./controller
```
