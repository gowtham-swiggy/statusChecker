package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var sites = make(map[string]int)

type SiteStruct struct {
	Website []string `json:"website"`
}

func main() {
	fmt.Println("Hello Netizen!...")
	// Home Page API
	http.HandleFunc("/home",homePageHandler)
	// Post API
	http.HandleFunc("/post",postHandler)
	//Get Details API
	http.HandleFunc("/get",getHandler)
	//Get Single Details API
	http.HandleFunc("/getsingle",getSingleHandler)
	//Anything API
	http.HandleFunc("/",anythingHandler)

	http.ListenAndServe(":5000",nil)

}

func homePageHandler(w http.ResponseWriter,r * http.Request) {
	fmt.Fprint(w,"You are in Home Page")
	fmt.Println("You are in Home Page")
}

func postHandler(w http.ResponseWriter,r * http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	site := SiteStruct{}
	err := json.NewDecoder(r.Body).Decode(&site)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _,j := range site.Website {
		sites[j] = statusFinder(j)
	}
}

func getHandler(w http.ResponseWriter,r * http.Request) {
	for i := range sites {
		var resultString string
		if sites[i] == 200 {
			resultString = "The status of website "+i+" is UP"
		} else {
			resultString = "The status of website "+i+" is DOWN"
		}
		fmt.Fprint(w,resultString)
	}
}

func getSingleHandler(w http.ResponseWriter,r * http.Request) {
	fmt.Fprint(w,"Error Gets ..")
	fmt.Println("You are in anything handler")
}

func anythingHandler(w http.ResponseWriter,r * http.Request) {
	fmt.Fprint(w,"Error 404...")
	fmt.Println("You are in anything handler")
}

func statusFinder(link string) int {
	resp,err := http.Get("https://"+link)
	if err!=nil {
		fmt.Println(err)
		return 0
	}
	return resp.StatusCode
}

