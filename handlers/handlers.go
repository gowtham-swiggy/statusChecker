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
		return 0, err
	}
	resp, err := client.Do(r) // Hitting
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close() // Closing
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
	models.Sites[link], _ = tmpVar.Checker() // Calling method to update the sites status.
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) { // Home Page Handler
	fmt.Fprint(w, "Welcome to Home Page")
	fmt.Println("You are in Home Page Handler")
}

func AnythingHandler(w http.ResponseWriter, r *http.Request) { // Invalid or 404 API Handler
	fmt.Fprint(w, "Error 404...")
	fmt.Println("You are in anything(404) Handler")
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	site := SiteStruct{}
	err := json.NewDecoder(r.Body).Decode(&site) //Decoding
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, j := range site.Website {
		var tmpVar httpLink
		tmpVar.link = j
		models.Sites[j], _ = tmpVar.Checker() // Adding the data
	}
	fmt.Println("You are in Post Handler")
	fmt.Fprint(w, models.Sites)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You are in Get Handler")
	var resultString string
	if len(models.Sites) == 0 {
		fmt.Fprint(w, "No Data in map")

	}
	for i := range models.Sites { // Looping over the map
		if models.Sites[i] == 200 { // Customize the result
			resultString = "The status of website " + i + " is UP\n"
		} else {
			resultString = "The status of website " + i + " is DOWN\n"
		}
		fmt.Fprint(w, resultString)
	}

}

func GetSingleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You are in Get Single Site Details Handler")
	website := r.URL.Query().Get("name") // Extracting the Query details
	var resultString string
	stats, ok := models.Sites[website] // Checking status and the existence
	if ok == true {                    // Customize the result
		if stats == 200 {
			resultString = "The status of website " + website + " is UP"
		} else {
			resultString = "The status of website " + website + " is DOWN"
		}
	} else {
		resultString = "The Website is not present, add the site First"
	}
	fmt.Fprint(w, resultString)
}
