package main

import (
	"fmt"
	"net/http"
)


func main() {
	fmt.Println("Hello Netizen!...")
	// Home Page API

	// Post API

	//Get Details API

	//Get Single Details API

	//Anything API

	http.ListenAndServe(":5000",nil)

}

func homePageHandler(w http.ResponseWriter,r * http.Request) {
	fmt.Fprint(w,"Error Home...")
	fmt.Println("You are in anything handler")
}

func postHandler(w http.ResponseWriter,r * http.Request) {
	fmt.Fprint(w,"Error 4Post...")
	fmt.Println("You are in anything handler")
}

func getHandler(w http.ResponseWriter,r * http.Request) {
	fmt.Fprint(w,"Error Get...")
	fmt.Println("You are in anything handler")
}

func getSingleHandler(w http.ResponseWriter,r * http.Request) {
	fmt.Fprint(w,"Error Gets ..")
	fmt.Println("You are in anything handler")
}

func anythingHandler(w http.ResponseWriter,r * http.Request) {
	fmt.Fprint(w,"Error 404...")
	fmt.Println("You are in anything handler")
}