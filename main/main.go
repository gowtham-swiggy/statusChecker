package main

import (
	"fmt"
	"net/http"

	"github.com/statusfinder/handlers"
	"github.com/statusfinder/models"
)

func main() { // Main function
	models.Updater = 0
	go handlers.StatusUpdaterUtility() // Go routine to update status of websites every one minute
	fmt.Println("Hello Netizen!...")   // Indicates server started
	// Home Page API
	http.HandleFunc("/home", handlers.HomePageHandler) // Home Page handler
	// Post API
	http.HandleFunc("/post", handlers.PostHandler) // Post page handler
	//Get Details API
	http.HandleFunc("/get", handlers.GetHandler) // Get the details of the all sites posted till time
	//Get Single Details API
	http.HandleFunc("/getsingle", handlers.GetSingleHandler) // Get details of single site
	//Anything API
	http.HandleFunc("/", handlers.AnythingHandler) // 404 error handler

	http.ListenAndServe(":3000", nil) // server running on port 3000

}
