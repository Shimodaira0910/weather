/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Shimodaira0910/weather/api"
	"github.com/spf13/cobra"
)

// weatherCmd represents the weather command
var weatherCmd = &cobra.Command{
	Use:   "weather",
	Run: func(cmd *cobra.Command, args []string) {
		api := api.Weather{}
		api.GetWeatherInfo("Osaka")
	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)
}
