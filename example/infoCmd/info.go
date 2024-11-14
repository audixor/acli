// Copyright (c) 2024 Tenebris Technologies Inc.
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package infoCmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Register() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "info",
		Long:  "info",
		Run: func(cmd *cobra.Command, args []string) {
			dryRun, _ := cmd.Flags().GetBool("dry")
			test(dryRun)
		},
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "one",
		Short: "info one",
		Long:  "info one",
		Run: func(cmd *cobra.Command, args []string) {
			dryRun, _ := cmd.Flags().GetBool("dry")
			one(dryRun)
		}})

	cmd.AddCommand(&cobra.Command{
		Use:   "two",
		Short: "info two",
		Long:  "info two",
		Run: func(cmd *cobra.Command, args []string) {
			dryRun, _ := cmd.Flags().GetBool("dry")
			two(dryRun)
		}})

	return cmd
}

func test(dryRun bool) {
	fmt.Println("info root")
}

func one(dryRun bool) {
	fmt.Println("info one")
}

func two(dryRun bool) {
	fmt.Println("info two")
}
