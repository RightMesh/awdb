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

package adb

// Device represents a connected device as reported by `adb devices -l`.
// See https://developer.android.com/studio/command-line/adb#devicestatus
type Device struct {
	// This device's unique ID on this host machine.
	// e.g. 1
	TransportID int

	// The device's serial ID. Not guaranteed to be unique.
	// e.g. "0a388e93"
	SerialID string

	// The device's state.
	// One of "authorized", "unauthorized", "authorizing", "device", "offline", "no device".
	State string

	// The USB port the device is connected to.
	// e.g. "1-1.4.2"
	Usb string

	// Optional, the device's product name.
	// e.g., "razor"
	Product string

	// Optional, the device's model name.
	// e.g. "Nexus_7"
	Model string

	// Optional, the device's name.
	// e.g. "flo"
	Device string
}
