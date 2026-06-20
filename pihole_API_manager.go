package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var ipRoute = "https://192.168.1.10:443/api"
var sID = "PLGkLhtsJo+5IL15be35yw="
var headers = [2]Header{
	{"accept", "application/json"},
	{"sid", sID},
}

func pihole_call[T any](route, method string, acceptExitCode int, payload any, headers []Header) (*T, error) {

	var responseStruct *T
	responseAPI := instAPICall(ipRoute+route, method, payload, headers)

	if responseAPI.err != nil {
		fmt.Printf("Failed to read body: %v", responseAPI.err)
		return nil, responseAPI.err
	}
	if acceptExitCode != responseAPI.exitCode {
		return nil, fmt.Errorf("Error, exit code %d: %s", responseAPI.exitCode, responseAPI.load)
	}
	responseStruct, errorParsing := parse_load[T](responseAPI.load)
	if errorParsing != nil {
		fmt.Printf("Error parsing : %v\n", errorParsing)
		return nil, errorParsing
	}

	return responseStruct, nil

}
func verify_sid() (bool, error) {

	response, errorInCall := pihole_call[get_auth]("/auth", http.MethodGet, OK, nil, headers[:])
	if errorInCall != nil {
		fmt.Printf("Error in pihole call: %v\n", errorInCall)
		return false, errorInCall
	}
	return response.Session.Valid, nil
}

func get_sid() (string, error) {
	password := os.Getenv("PASSWORD")
	load := map[string]string{"password": password}
	response, errorInCall := pihole_call[get_auth]("/auth", http.MethodPost, OK, load, headers[:])
	if errorInCall != nil {
		fmt.Printf("Error in pihole call: %v\n", errorInCall)
		return "", errorInCall
	}
	return *response.Session.Sid, nil
}

func get_history() ([]History, error) {

	response, errorInCall := pihole_call[history_database]("/auth", http.MethodGet, OK, nil, headers[:])
	if errorInCall != nil {
		fmt.Printf("Error in pihole call: %v\n", errorInCall)
		return nil, errorInCall
	}

	return response.History, nil
}
func main() {
	errorEnv := godotenv.Load()
	if errorEnv != nil {
		fmt.Printf("Error Loading the env file: %v\n", errorEnv)
	}
	sID = os.Getenv("SID")

	sID_valid, error_sID_Validation := verify_sid()
	if error_sID_Validation != nil || sID_valid == false {
		fmt.Printf("sID invalid!\nGenerating a new sID\n")
		sID_aux, error_sID_Generation := get_sid()
		if error_sID_Generation != nil {
			fmt.Printf("Error generating the new sid: %v\n", error_sID_Generation)
		}
		fmt.Printf("new sID generated: %s\n", sID_aux)
		sID = sID_aux
	} else {
		fmt.Printf("sID valid!\n")
	}
}
