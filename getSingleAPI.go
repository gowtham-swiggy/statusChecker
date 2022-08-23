package main

import (
	"fmt"
	"net/http"
)

func GetSingleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You are in Get Single Site Details Handler")
	website := r.URL.Query().Get("name")
	var resultString string
	if sites[website] == 200 {
		resultString = "The status of website " + website + " is UP"
	} else {
		resultString = "The status of website " + website + " is DOWN"
	}
	fmt.Fprint(w, resultString)
}