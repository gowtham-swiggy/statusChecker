package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var updater int

var sites = make(map[string]int)

type SiteStruct struct {
	Website []string `json:"website"`
}

type StatusFinder interface {
	Checker(link string) (code int, err error)
}

type httpLink struct {
	link string
}

func (h httpLink) Checker() (stausCode int, err error) {
	client := http.Client{}
	r, err := http.NewRequest("GET", "http://"+h.link, nil)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}

func main() {
	updater = 0
	go StatusUpdaterUtility()
	fmt.Println("Hello Netizen!...")
	// Home Page API
	http.HandleFunc("/home", HomePageHandler)
	// Post API
	http.HandleFunc("/post", PostHandler)
	//Get Details API
	http.HandleFunc("/get", GetHandler)
	//Get Single Details API
	http.HandleFunc("/getsingle", GetSingleHandler)
	//Anything API
	http.HandleFunc("/", AnythingHandler)

	http.ListenAndServe(":3000", nil)

}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You are in Home Page")
	fmt.Println("You are in Home Handler")
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	site := SiteStruct{}
	err := json.NewDecoder(r.Body).Decode(&site)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, j := range site.Website {
		var tmpVar httpLink
		tmpVar.link = j
		sites[j], err = tmpVar.Checker()
		if err != nil {
			fmt.Println("Error occured for " + j)
			fmt.Println(err)
		}
	}
	fmt.Println("You are in Post Handler")
	fmt.Fprint(w, sites)
}

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

func AnythingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Error 404...")
	fmt.Println("You are in anything Handler")
}

// func StatusFinder(link string) int {
// 	client := http.Client{}
// 	r, err := http.NewRequest("GET", "http://"+link, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		return 0
// 	}
// 	resp, err := client.Do(r)
// 	if err != nil {
// 		fmt.Println(err)
// 		return 0
// 	}
// 	defer resp.Body.Close()
// 	return resp.StatusCode
// }

func StatusUpdater(link string) {
	var tmpVar httpLink
	tmpVar.link = link
	var err error
	sites[link], err = tmpVar.Checker()
	if err != nil {
		fmt.Println(err)
	}
}

func StatusUpdaterUtility() {
	for {
		for i := range sites {
			go StatusUpdater(i)
		}
		updater++
		var tmp string
		if updater == 1 {
			tmp = " time."
		} else {
			tmp = " times."
		}
		fmt.Println("Updated Data " + strconv.Itoa(updater) + tmp)

		time.Sleep(60 * time.Second)
	}
}
