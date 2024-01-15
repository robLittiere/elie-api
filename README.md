# ELIE API

## What is this?

Elie API is the backend supporting the Elie application. It is a RESTful API built with Golang.  
It consists of a martini-like API called Gin and currently supports a Postgress database (see docker-compose file for more details).  
The choice to use go coupled with Gin was made to have lightning fast and easy to build/maintain API so that the focus can be on the frontend.

## Installation
Download and install Go on your system, add it to your PATH.  

Download the project in the folder of your choice and run the following command:
```bash
go run main.go
```

## Docker

You can also run the project using docker.
> You will need to modify the config file for now as we do not have .env file support yet.  
> To do this, go to the config folder and modify the db.go file.  Modify the following line:

```go
    dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Paris"
    ```bash
    docker-compose up
```

Modify the hostname to be the one of the docker container (for now "db" in our docker-compose file).

```go
    dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Paris"
    ```bash
    docker-compose up
```

You will have indication on the port the API is running on. It should be http://localhost:8080 by default (you can modify this in the main.go file if you want).

## Build
If you ever want to build the project you can run the following command:
```bash
go build
```