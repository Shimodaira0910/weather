/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/Shimodaira0910/weather/api"
	"github.com/spf13/cobra"
)

// weatherCmd represents the weather command
var weatherCmd = &cobra.Command{
	Use:   "weather",
	Run: func(cmd *cobra.Command, args []string) {
		cityName := ""
		
		for cityName == ""{
			fmt.Println("検索したい都市名を、ローマ字で入力してください(頭文字を大文字で入力してください)")
			_, err := fmt.Scanln(&cityName)

			if err != nil {
				fmt.Println("都市名を入力してください。")
			}
		}
		api := api.Weather{}
		api.GetWeatherInfo(cityName)
	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)
}
