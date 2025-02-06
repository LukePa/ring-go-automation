/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/playwright-community/playwright-go"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ring-go-automation",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		test()
	},
}

func getPlaywrightInstance() (*playwright.Playwright, error) {
	pw, err := playwright.Run()
	if err != nil {
		fmt.Println("Playwright installation not found, would you like to install it?")
		var input string
		fmt.Print("Enter (y) to install: ")
		fmt.Scan(&input)

		if input == "y" {
			playwright.Install()
			return playwright.Run()
		}
	}

	return pw, err
}

func test() {
	pw, err := getPlaywrightInstance()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}

	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	if _, err = page.Goto("https://myringgo.co.uk/account/login"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	content, err := page.Content()
	if err != nil {
		log.Fatalf("could not get page content: %v", err)
	}
	fmt.Print(content)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ring-go-automation.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
