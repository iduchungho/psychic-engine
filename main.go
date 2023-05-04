// backend for smart home projects
// author:
//
//	@yesic4n
//	@lamdienchinh
//	@Nguyenleminhbao-tt5
//
// /////////////////////////////////
// main.go file
// Author: Ho Duc Hung
// Modify .env.example to .env
// Start api:  go run main.go
package main

import (
	appli "smhome/app"
	"smhome/pkg/utils"
)

// ! @title smhome: main
// ! @version 1.0
// ! @description: main function where every beginning.
// ! @BasePath /api
// ! @yesic4n GitHub
// ! @name Authorization
func main() {
	////////////////////////////////
	// comment line for heroku deployment
	utils.LoadEnvFile()
	/////////////////////////////////
	// create application
	app := appli.GetApplication()
	// app run localhost:8080
	app.Run()
}
