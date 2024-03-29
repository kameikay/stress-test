/*
Copyright © 2024 Victor Kamei Kay <kameikay@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/kameikay/stress-test/internal/core/service"
	"github.com/kameikay/stress-test/internal/repository"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stress-test",
	Short: "A stress test tool for web applications",
	Long: `A stress test tool for web applications.
				It sends multiple requests to a given URL and generates a report
				with the total execution time, total requests made,
				number of requests that returned HTTP status code 200 and
				number of requests that returned other HTTP status codes.`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt("requests")
		concurrency, _ := cmd.Flags().GetInt("concurrency")

		repo := repository.NewRequestsMemory()
		stressService := service.New(repo)
		stressService.Execute(url, requests, concurrency)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "URL to test")
	rootCmd.MarkFlagRequired("url")
	rootCmd.Flags().IntP("requests", "r", 1, "Number of requests to make")
	rootCmd.Flags().IntP("concurrency", "c", 1, "Number of concurrent requests")
}
