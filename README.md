# BACKEND TO DO APP

This backend serves the managed tasks from FRONTEND TO DO APP.

Here you can find endpoints to be possible the CRUD for management of tasks and the creation of a postgres database using an image with docker.

## INSTALL AND RUN BACKEND

1. Install packages locally

````
go mod vendor
go mod tidy
````
2. Connect to database

````
sudo docker run -p 54321:5432 tasks-db
````

3. Go to cmd and run the project 

````
go run main.go
````