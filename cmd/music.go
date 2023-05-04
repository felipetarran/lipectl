/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// musicCmd represents the music command
var musicCmd = &cobra.Command{
	Use:   "music",
	Short: "Get good Music",
	Long: `With this command you will receive a artist recommendation based on the music style you choose between: 
	Rock
	Jazz
	Blues.`,
	Run: func(cmd *cobra.Command, args []string) {
		musicStyle, _ := cmd.Flags().GetString("style")
		var artist string
		if musicStyle == "jazz" {
			artist = "Hank Mobley"
		} else if musicStyle == "rock" {
			artist = "Free"
		} else if musicStyle == "blues" {
			artist = "Clarence Gatemouth Brown"
		}

		fmt.Println("Artista recomendado:", artist)

	},
}

func init() {
	rootCmd.AddCommand(musicCmd)
	musicCmd.Flags().StringP("style", "s", "", "specify style recomendation")
}
