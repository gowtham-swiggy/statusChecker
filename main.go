package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var sites = make(map[string]int)

type SiteStruct struct {
	Website []string `json:"website"`
}

func main() {
	go statusUpdaterUtility()
	fmt.Println("Hello Netizen!...")
	// Home Page API
	http.HandleFunc("/home", homePageHandler)
	// Post API
	http.HandleFunc("/post", postHandler)
	//Get Details API
	http.HandleFunc("/get", getHandler)
	//Get Single Details API
	http.HandleFunc("/getsingle", getSingleHandler)
	//Anything API
	http.HandleFunc("/", anythingHandler)

	http.ListenAndServe(":5000", nil)

}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You are in Home Page")
	fmt.Println("You are in Home Page")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"HI\n")
	w.Header().Set("Content-Type", "Application/json")
	site := SiteStruct{}
	err := json.NewDecoder(r.Body).Decode(&site)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, j := range site.Website {
		sites[j] = statusFinder(j)
	}
	fmt.Print(w, sites)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	for i := range sites {
		var resultString string
		if sites[i] == 200 {
			resultString = "The status of website " + i + " is UP"
		} else {
			resultString = "The status of website " + i + " is DOWN"
		}
		fmt.Fprint(w, resultString)
	}
}

func getSingleHandler(w http.ResponseWriter, r *http.Request) {
	website := r.URL.Query().Get("name")
	var resultString string
	if sites[website] == 200 {
		resultString = "The status of website " + website + " is UP"
	} else {
		resultString = "The status of website " + website + " is DOWN"
	}
	fmt.Fprint(w, resultString)
}

func anythingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Error 404...")
	fmt.Println("You are in anything handler")
}

// func statusFinder(link string) int {
// 	resp,err := http.Get("https://"+link)
// 	if err!=nil {
// 		fmt.Println(err)
// 		return 0
// 	}
// 	return resp.StatusCode
// }
func statusFinder(link string) int {
	client := http.Client{}
	r, err := http.NewRequest("GET", "http://"+link, nil)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer resp.Body.Close()
	return resp.StatusCode
}

func statusUpdater(link string) {
	sites[link] = statusFinder(link)
}

func statusUpdaterUtility() {
	for {
		for i := range sites {
			go statusUpdater(i)
		}
		fmt.Println("Inside statusUpdater")
		time.Sleep(60 * time.Second)
	}
}
