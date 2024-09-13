/*
 * Project: QoolQR
 * Filename: /cmd/create.go
 * Created Date: Friday September 13th 2024 11:27:36 +0800
 * Author: Sallehuddin Abdul Latif (salleh@sallehuddin.com)
 * Company: Sallehuddin Industries
 * --------------------------------------
 * Last Modified: Friday September 13th 2024 11:40:35 +0800
 * Modified By: Sallehuddin Abdul Latif (salleh@sallehuddin.com)
 * --------------------------------------
 * Copyright (c) 2024 Sallehuddin Abdul Latif
 */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")

	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
