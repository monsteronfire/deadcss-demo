/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("scan called")
		testPythonService()
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Test if Python service is running
func testPythonService() error {
	// resp, err := http.Get("http://localhost:8000/health")
	payload := map[string]interface{}{
		"css_content":   "body { color: red; }",
		"project_files": []string{"index.html", "styles.css"},
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}

	resp, err := http.Post(
		"http://localhost:8000/api/analyse-css",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		fmt.Println("Error connecting to Python AI service:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		fmt.Println("Python AI service is not found!")
	}

	if resp.StatusCode == 200 {
		fmt.Println("Python AI service is running!")
	}
	return nil
}
