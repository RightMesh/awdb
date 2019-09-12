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

package server

import (
	"encoding/json"
	"github.com/rightmesh/awdb/pkg/adb"
	"net/http"
)

// helpHandler returns the contents of `adb help` as plaintext.
// May return a 502 if communicating with ADB fails.
func helpHandler(response http.ResponseWriter, request *http.Request) {
	adbRun := adb.NewRun("help")
	if err := adbRun.Output(); err != nil {
		writeResponse(response, http.StatusBadGateway, contentTypeText, adbRun.StdErr)
		return
	}

	writeResponse(response, http.StatusOK, contentTypeText, adbRun.StdOut)
}

// devicesHandler returns the contents of `adb devices -l` as JSON.
// May return a 502 if communicating with ADB fails, or a 500 if marshalling JSON fails.
func devicesHandler(response http.ResponseWriter, request *http.Request) {
	adbRun := adb.NewRun("devices", "-l")
	if err := adbRun.Output(); err != nil {
		writeResponse(response, http.StatusBadGateway, contentTypeText, adbRun.StdErr)
		return
	}

	deviceList, err := adb.ParseDeviceList(adbRun.StdOut)
	if err != nil {
		writeResponse(response, http.StatusBadGateway, contentTypeText, []byte(err.Error()))
		return
	}

	jsonBytes, err := json.Marshal(deviceList)
	if err != nil {
		writeResponse(response, http.StatusInternalServerError, contentTypeText, []byte(err.Error()))
		return
	}

	writeResponse(response, http.StatusOK, contentTypeJSON, jsonBytes)
}
