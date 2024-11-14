// Copyright (c) 2024 Tenebris Technologies Inc.
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package main

import (
	"github.com/audixor/acli/example/infoCmd"
	"github.com/audixor/acli/example/testCmd"

	"github.com/audixor/acli"
)

func main() {
	c := acli.New("easy-test", "testing", "still testing")
	c.AddBoolFlag("dry", false, "Dry run")
	c.AddCmd(testCmd.Register)
	c.AddCmd(infoCmd.Register)
	c.Execute()
}
