/*
main.go file
Author: Ho Duc Hung
Start api:  go run main.go
*/
package main

import appli "smhome/app"

func main() {
	// create application
	app := appli.GetApplication()
	// app run localhost:8080
	app.Run()
}
