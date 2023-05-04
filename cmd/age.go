/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"io"
	//"encoding/json"
)

// ageCmd represents the age command
var ageCmd = &cobra.Command{
	Use:   "age",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		resp, err := http.Get("https://api.agify.io?name=" + name)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}
		fmt.Println(string(body))
		defer resp.Body.Close()
	},
}

func init() {
	rootCmd.AddCommand(ageCmd)
	ageCmd.Flags().StringP("name", "n", "", "specify name for age")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
