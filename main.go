package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/spf13/viper"
)

func main() {
	// Init config
	viper.AddConfigPath("application.properties.yaml")

	url := "http://che4ppayapi01.cheadds.ca:8708/timeoff/pphr-monitor/health"
	fmt.Println(url)

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + viper.GetString("accessToken")

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panicln("Error on request creation!\n[ERROR] -", err)
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)

	// Unmarshal results into our info struct
	var result Info
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	// Format & Print
	var finalResult string = Beautify(result)
	fmt.Println(finalResult)

	// Write results ot the log file
	createLogDirectory()
	writeToLogFile(finalResult)
	log.Printf("\n" + sb)

	// Format status of Microservices
	logMicroserviceStatus(result)
}

func Beautify(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func createLogDirectory() {
	path, _ := os.Getwd()
	path += "/microserviceLogs"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

func writeToLogFile(response string) {
	var fileName string = "microservice_status_" + time.Now().String() + ".txt"
	os.Chdir("microserviceLogs")
	f, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(response)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Done writing to the file")
}

func logMicroserviceStatus(data Info) {
	results := data.Details.DiscoveryComposite.Details.Eureka.Details.Applications
	v := reflect.ValueOf(results)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Microservice: %s\t Status: %v\n", typeOfS.Field(i).Name, getStatus(v.Field(i).Interface()))
	}
}

func getStatus(i interface{}) string {
	if i == 1 {
		return "OPERATIONAL :)"
	} else {
		return "OFFLINE :("
	}
}
