// This is a code test which tests fetching an API, grouping the data and finally outputing a CSV file. It should be ran as a Command Line Interface tool.

// The URL is http://api.open-notify.org/astros.json, it provides a live list of people in space.

// Once fetched, please sort the data a-z by craft and then name.

// Finally output the file in a CSV format similar to below or following the CSV standard.

// Name,Craft
// Fei Junlong,Shenzhou 15
// Alexander Skvortsov,ISS
// Andrew Morgan,ISS
// Christina Koch,ISS
// Jessica Meir,ISS
// Luca Parmitano,ISS
// Oleg Skripochka,ISS

package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// A Response struct to map the Entire Response
type Response struct {
	People []People `json:"people"`
}

// A People Struct to map the craft and name to.
type People struct {
	Craft string `json:"craft"`
	Name  string `json:"name"`
}

func main() {
	// get the api
	response, err := http.Get("http://api.open-notify.org/astros.json")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	// read from the api
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// convert byte data into original data structure : List of People
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	// get length of data structure
	var length int = len(responseObject.People)
	// create an empty result list/slice to output
	// sort by name
	for i := 1; i < length; i++ {
		j := i
		// compare the name before with current name
		for j > 0 && responseObject.People[j-1].Name > responseObject.People[j].Name {
			// if previous name is bigger, swap them
			responseObject.People[j].Name, responseObject.People[j-1].Name = responseObject.People[j-1].Name, responseObject.People[j].Name
			j--
		}
	}
	// sort by craft
	for i := 1; i < length; i++ {
		j := i
		// compare the craft before with current craft
		for j > 0 && responseObject.People[j-1].Craft > responseObject.People[j].Craft {
			// if previous craft is bigger, swap them
			responseObject.People[j].Craft, responseObject.People[j-1].Craft = responseObject.People[j-1].Craft, responseObject.People[j].Craft
			j--
		}
	}
	// This is to create the CSV file
	file, err := os.Create("People.csv")
	// defer performed after the current function is executed
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	// using to write into the CSV file
	w := csv.NewWriter(file)
	// Flush is to ensure that the name and craft is written to the underlying writer.
	defer w.Flush()

	// adding all the data into the right format into a slice
	var data [][]string
	data = append(data, []string{"Name", "Craft"})
	for _, person := range responseObject.People {
		row := []string{person.Name, person.Craft}
		data = append(data, row)
	}
	// writing the data into the file
	w.WriteAll(data)
}
