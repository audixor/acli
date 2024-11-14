// Copyright (c) 2024 Tenebris Technologies Inc.
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package testCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Register() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "test",
		Long:  "test",
		Run: func(cmd *cobra.Command, args []string) {
			dryRun, _ := cmd.Flags().GetBool("dry")
			test(dryRun)
		},
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "one",
		Short: "test one",
		Long:  "test one",
		Run: func(cmd *cobra.Command, args []string) {
			dryRun, _ := cmd.Flags().GetBool("dry")
			one(dryRun)
		}})

	cmd.AddCommand(&cobra.Command{
		Use:   "two",
		Short: "test two",
		Long:  "test two",
		Run: func(cmd *cobra.Command, args []string) {
			dryRun, _ := cmd.Flags().GetBool("dry")
			two(dryRun)
		}})

	return cmd
}

func test(dryRun bool) {
	fmt.Println("root test")
}

func one(dryRun bool) {
	fmt.Println("test one")
}

func two(dryRun bool) {
	fmt.Println("test two")
}
