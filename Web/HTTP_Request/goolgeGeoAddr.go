package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const(
	BUFFLIMIT int (1024 * 5)
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
func GeocodingRequest(url string) *[]byte {

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
	//log.Println(string(body))
	return &body
}

// ReverseGeocodingRequest HTTP request function
func ReverseGeocodingRequest(url string) *[]byte {

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
	//log.Println(string(body))
	return &body
}

// Read JSON
func readJSONFile(keyFile string) *string {

	apikeyJSONFile, err := os.Open(keyFile)
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

	return &apikey.Key

}

func unmarshalJSON() {

}

func printToFile(data *string) (bool, int) {
	var dataWritten bool
	f, err := os.Create("locationdata.json")
	if err != nil {
		dataWritten = false
		log.Fatal(err)
	}
	buffSize := (BUFFLIMIT)
	w := bufio.NewWriterSize(f, buffSize)
	bytesWritten, err := w.WriteString(*data)
	if err != nil {
		dataWritten = false
		log.Fatal(err)
	}
	dataWritten = true
	w.Flush()
	return dataWritten, bytesWritten
}

func main() {

	apikey := readJSONFile("apikey.json")

	address := Address{
		StreetNumber: "1600",
		StreetName:   "Amphitheatre Parkway",
		City:         "Mountain View",
		State:        "CA",
	}

	geoLocation := GeoLocation{
		Latitude:  "40.714224",
		Longitude: "-73.961452",
	}

	formatedURL := address.addressFormatter(&address, *apikey)
	fmt.Println("---------ADDRESS LOOKUP----------------")
	GeocodingRequest(formatedURL)

	formatedURL = geoLocation.addressFormatter(&geoLocation, *apikey)
	fmt.Println("---------REVERSE LOOKUP----------------")
	reverBytes := ReverseGeocodingRequest(formatedURL)
	//fmt.Println(string(*reverBytes))

	stringData := string(*reverBytes)
	//fmt.Println(stringData)
	printToFile(&stringData)

	// var results map[string]interface{}
	// json.Unmarshal([]byte(*byteStr), &results)
	// loc := results["results"].(map[string]interface{})
	//unmarshalJSON(results)

}
