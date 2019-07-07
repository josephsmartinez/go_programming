package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// APIKey JSON for Google API
type APIKey struct {
	Key string `json:"key"`
}

// Address data for url creatation
type Address struct {
	StreetNumber string
	StreetName   string
	City         string
	State        string
}

// GeoLocation data for url creatation
type GeoLocation struct {
	Latitude  string
	Longitude string
}

func (g GeoLocation) addressFormatter(geoloc *GeoLocation, apiKey string) string {
	url := "https://maps.googleapis.com/maps/api/geocode/json?latlng=" + geoloc.Latitude + "," + geoloc.Longitude + "&key=" + apiKey
	return url
}

func (a Address) addressFormatter(addr *Address, apiKey string) string {
	var str strings.Builder
	str.WriteString(addr.StreetNumber)
	streetNameSplit := strings.Split(addr.StreetName, " ")
	for i, w := range streetNameSplit {
		if i < len(streetNameSplit) {
			str.WriteString("+" + w)
		}
	}
	str.WriteString(",")
	cityNameSplit := strings.Split(addr.City, " ")
	for i, w := range cityNameSplit {
		if i < len(cityNameSplit) {
			str.WriteString("+" + w)
		}
	}
	str.WriteString(",")
	str.WriteString(addr.State)
	url := "https://maps.googleapis.com/maps/api/geocode/json?address=" + str.String() + "&key=" + apiKey
	return url

}

// GeocodingRequest HTTP request function
func GeocodingRequest(url string) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		os.Exit(2)
	}

	log.Println(string(body))
}

// ReverseGeocodingRequest HTTP request function
func ReverseGeocodingRequest(url string) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		os.Exit(2)
	}

	log.Println(string(body))
}

func main() {

	apikeyJSONFile, err := os.Open("apikey.json")
	if err != nil {
		log.Fatal(err)
	}
	defer apikeyJSONFile.Close()

	byteValue, _ := ioutil.ReadAll(apikeyJSONFile)
	apikey := new(APIKey)
	err = json.Unmarshal(byteValue, &apikey)
	if err != nil {
		log.Fatal(err)
	}

	address := Address{
		StreetNumber: "1600",
		StreetName:   "Amphitheatre Parkway",
		City:         "Mountain View",
		State:        "CA",
	}

	formatedURL := address.addressFormatter(&address, apikey.Key)
	fmt.Println("---------ADDRESS LOOKUP----------------")
	GeocodingRequest(formatedURL)

	geoLocation := GeoLocation{
		Latitude:  "40.714224",
		Longitude: "-73.961452",
	}
	formatedURL = geoLocation.addressFormatter(&geoLocation, apikey.Key)
	fmt.Println("---------REVERSE LOOKUP----------------")
	ReverseGeocodingRequest(formatedURL)
}
