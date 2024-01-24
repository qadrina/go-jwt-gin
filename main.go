package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qadrina/go-jwt-gin/controllers"
	"github.com/qadrina/go-jwt-gin/initializers"
	"github.com/qadrina/go-jwt-gin/middleware"
)

/*
1. create main.go -> create .env file, enter PORT -> create initializers folder -> create loadEnvVariables.go inside the init folder.
2. create database di mssql
3. create models -> migrate to db by creating syncDatabase.go -> tambah sync db di main init
4. create controllers folder (for the signup thing)
5. connect the signup function to the router -> test on postman
6. create login function di controller -> add var secret key to decode and encode inside the .env file -> tambah route di main -> test in postman
7. add cookies
8. create Auth middleware  (func validate di ctrler) -> add to router -> SECURE the route by creating a new folder named MIDDLEWARE -> require Auth.go
9. ke controller -> tambah di validate
*/
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	//initializers.SyncDatabase()
}

func main() {
	// create a gin app
	r := gin.Default()

	// the routes
	r.POST("/api/signup", controllers.SignUp)
	r.POST("/api/login", controllers.Login)
	r.GET("/api/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}
