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

// Package adb provides methods for interacting with the ADB executable and reasoning
// about its output.
package adb

import (
	"os/exec"
)

// Command represents the execution of an ADB command.
type Command struct {
	cmd    *exec.Cmd
	args   []string
	StdOut []byte
	StdErr []byte
}

// Run executes the underlying ADB command and waits for it to finish.
// If the call is successful, Command.StdOut is populated with the output and the error is nil.
// If the call is unsuccessful, the error is returned, and if it is an exec.ExitError it is used
// to populate Command.StdErr.
func (a *Command) Run() error {
	output, err := a.cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			a.StdErr = exitError.Stderr
		}
		return err
	}

	a.StdOut = output
	return nil
}

// NewCommand initializes a Command struct with an exec.Cmd that will call ADB with the
// provided arguments and leaves the output bytes slices empty.
func NewCommand(args ...string) Command {
	command := exec.Command("adb", args...)
	return Command{cmd: command, args: args}
}
