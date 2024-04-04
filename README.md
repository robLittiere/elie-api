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

You will have indication on the port the API is running on. It should be http://localhost:8080 by default (you can modify this in the main.go file if you want).

## Migrations

We are using a simple tool called [Goose](https://github.com/pressly/goose#go-migrations) to run migrations.
Simply do a `go install github.com/pressly/goose/v3/cmd/goose@latest` to install it.

Then you can run the following command to run the migrations:
```bash
goose -dir migrations postgres "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Paris" up
```

If you wish to create a migration, you can do so with the following command : 
```bash
goose -dir migrations postgres "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Paris" create <migration_name> sql
```
Then simply add the SQL  code you want in the script

## Docker

You can also run the project using docker.
> You will need to modify the config file for now as we do not have .env file support yet.  
> To do this, go to the config folder and modify the db.go file.  Modify the following line:

```go
func ConnectDB() {
    ...
    dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Paris"
```

Modify the hostname to be the one of the docker container (for now "db" in our docker-compose file).

```go
func ConnectDB() {
    ...
    dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Paris"
```

## Live reload
Elie API uses [Air](https://github.com/cosmtrek/air) for live reload.  

If you wish to perform a live reload, install `air` on your machine and run the following command:
```bash
air
```
`Air` uses the `.air.toml` file to know what to do. You can modify it to your needs. You can specify another config file by running `air -c <config_file>`.

## Build
If you ever want to build the project you can run the following command:
```bash
go build
```