/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Shimodaira0910/weather/api"
	"github.com/spf13/cobra"
)

// weatherCmd represents the weather command
var weatherCmd = &cobra.Command{
	Use:   "weather",
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("検索したい都市名を、ローマ字で入力してください(文字の最初は大文字です)")
		scanner.Scan()
		cityName := scanner.Text()
		
		api := api.Weather{}
		api.GetWeatherInfo(cityName)
	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)
}
