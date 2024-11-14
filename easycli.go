// Copyright (c) 2024 Tenebris Technologies Inc.
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package acli

import (
	"fmt"

	"github.com/spf13/cobra"
)

type ACLI struct {
	rootCmd *cobra.Command
}

func New(name string, short string, long string) *ACLI {
	return &ACLI{
		rootCmd: &cobra.Command{
			Use:   name,
			Short: short,
			Long:  long,
		},
	}
}

func (c *ACLI) AddCmd(newCmd func() *cobra.Command) {
	c.rootCmd.AddCommand(newCmd())
}

func (c *ACLI) AddBoolFlag(name string, value bool, usage string) {
	c.rootCmd.PersistentFlags().Bool(name, value, usage)
}

func (c *ACLI) AddStringFlag(name string, value string, usage string) {
	c.rootCmd.PersistentFlags().String(name, value, usage)
}

func (c *ACLI) AddIntFlag(name string, value int, usage string) {
	c.rootCmd.PersistentFlags().Int(name, value, usage)
}

func (c *ACLI) Execute() {
	if err := c.rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
