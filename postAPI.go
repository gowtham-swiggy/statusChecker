package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
