package main

import (
	"fmt"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You are in Get Handler")
	for i := range sites {
		var resultString string
		if sites[i] == 200 {
			resultString = "The status of website " + i + " is UP\n"
		} else {
			resultString = "The status of website " + i + " is DOWN\n"
		}
		fmt.Fprint(w, resultString)
	}

}