package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/statusfinder/models"
)

type SiteStruct struct {
	Website []string `json:"website"` // Struct to marshall and unmarshall the information
}

type StatusFinder interface {
	Checker(link string) (code int, err error) // Interface to handle future cases
}

type httpLink struct { // type
	link string
}

func (h httpLink) Checker() (stausCode int, err error) { // website Status Finder site.
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

func StatusUpdaterUtility() {
	for {
		for i := range models.Sites {
			go StatusUpdater(i) // Create N go routines
		}
		models.Updater++ // updater value increased
		var tmp string
		if models.Updater == 1 {
			tmp = " time."
		} else {
			tmp = " times."
		}
		fmt.Println("Updated Data " + strconv.Itoa(models.Updater) + tmp)

		time.Sleep(60 * time.Second) // Used sleep to make this run every one minute
	}
}

func StatusUpdater(link string) {
	var tmpVar httpLink
	tmpVar.link = link
	var err error
	models.Sites[link], err = tmpVar.Checker() // Calling method to update the sites status.
	if err != nil {
		fmt.Println(err)
	}
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Home Page")
	fmt.Println("You are in Home Page Handler")
}

func AnythingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Error 404...")
	fmt.Println("You are in anything(404) Handler")
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
		models.Sites[j], err = tmpVar.Checker()
		if err != nil {
			fmt.Println("Error occured for " + j)
			fmt.Println(err)
		}
	}
	fmt.Println("You are in Post Handler")
	fmt.Fprint(w, models.Sites)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You are in Get Handler")
	for i := range models.Sites {
		var resultString string
		if models.Sites[i] == 200 {
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
	if models.Sites[website] == 200 {
		resultString = "The status of website " + website + " is UP"
	} else {
		resultString = "The status of website " + website + " is DOWN"
	}
	fmt.Fprint(w, resultString)
}
