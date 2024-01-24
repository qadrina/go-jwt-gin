package main

import "github.com/qadrina/go-jwt-gin/initializers"

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.SyncDatabase()
}
