// Copyright (c) 2024 Tenebris Technologies Inc.
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package easycli

import (
	"fmt"

	"github.com/spf13/cobra"
)

type EasyCLI struct {
	rootCmd *cobra.Command
}

func New(name string, short string, long string) *EasyCLI {
	return &EasyCLI{
		rootCmd: &cobra.Command{
			Use:   name,
			Short: short,
			Long:  long,
		},
	}
}

func (c *EasyCLI) AddCmd(newCmd func() *cobra.Command) {
	c.rootCmd.AddCommand(newCmd())
}

func (c *EasyCLI) AddBoolFlag(name string, value bool, usage string) {
	c.rootCmd.PersistentFlags().Bool(name, value, usage)
}

func (c *EasyCLI) AddStringFlag(name string, value string, usage string) {
	c.rootCmd.PersistentFlags().String(name, value, usage)
}

func (c *EasyCLI) AddIntFlag(name string, value int, usage string) {
	c.rootCmd.PersistentFlags().Int(name, value, usage)
}

func (c *EasyCLI) Execute() {
	if err := c.rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
