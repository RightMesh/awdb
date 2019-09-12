/*
   Copyright 2019 Left Technologies Inc.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

// Package server maps routes to handler functions.
package server

import (
	"bytes"
	"encoding/json"
	"github.com/rightmesh/awdb/pkg/adb"
	"io"
	"log"
	"net/http"
)

// helpHandler returns the contents of `adb help` as plaintext.
func helpHandler(response http.ResponseWriter, request *http.Request) {
	adbRun := adb.NewRun("help")
	if err := proxyAdbRun(response, &adbRun); err != nil {
		return
	}

	response.Header().Set("Content-Type", "text/plain; charset=utf-8")
	response.WriteHeader(http.StatusOK)
	response.Write(adbRun.StdOut)
}

// devicesHandler returns the contents of `adb devices -l` as JSON.
// TODO: Add example JSON here.
func devicesHandler(response http.ResponseWriter, request *http.Request) {
	adbRun := adb.NewRun("devices", "-l")
	if err := proxyAdbRun(response, &adbRun); err != nil {
		return
	}

	deviceList, err := adb.ParseDeviceList(adbRun.StdOut)
	if err != nil {
		// TODO: Make a 502 helper method
		return
	}

	writeResponseAsJSON(response, deviceList)
}

// proxyAdbRun executes the command stored in the provided adb.Run instance and handles
// reporting an error if one occurs by writing a 502 error to the response along with the
// contents of stderr, and returning the error.
// If no errors occur, nil is returned and nothing is written to the response.
func proxyAdbRun(response http.ResponseWriter, adbRun *adb.Run) (err error) {
	err = adbRun.Output()
	if err != nil {
		response.Header().Set("Content-Type", "text/plain; charset=utf-8")
		response.WriteHeader(http.StatusBadGateway)
		response.Write(adbRun.StdErr)
	}

	// TODO: Trim out ADB debugging lines. E.g.:
	//     * daemon not running; starting now at tcp:5037
	//     * daemon started successfully

	return err
}

// writeResponseAsJSON attempts to marshal the provided data to JSON, writing it to the provided
// response with an HTTP 200 code if successful, or writing a 502 to the provided response if not.
func writeResponseAsJSON(response http.ResponseWriter, data interface{}) error {
	temp := new(bytes.Buffer)

	encoder := json.NewEncoder(temp)
	err := encoder.Encode(data)
	if err != nil {
		response.Header().Set("Content-Type", "text/plain; charset=utf-8")
		response.WriteHeader(http.StatusBadGateway)
		return err
	}

	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	response.WriteHeader(http.StatusOK)
	io.Copy(response, temp)
	return nil
}

// Start sets up the HTTP routes and serves the service, crashing if any errors are encountered.
func Start() {
	http.HandleFunc("/help/", helpHandler)
	http.HandleFunc("/devices/", devicesHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
