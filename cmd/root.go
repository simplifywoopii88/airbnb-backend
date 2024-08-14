/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: setUpDatabase,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var database string

func setUpDatabase(cmd *cobra.Command, args []string) {
	dbbool := false
	if database == "true" || database == "false" {
		dbbool, _ = strconv.ParseBool(database)
	} else {
		log.Println("database: true and false")
		os.Exit(1)
		return
	}

	if dbbool {
		fmt.Println("database is set to true")
		command := exec.Command("sh", "-c", "docker compose -f ./docker-compose/dev/docker-compose.yml down && docker compose -f ./docker-compose/dev/docker-compose.yml up -d")
		command.Stdout = os.Stdout

		if err := command.Run(); err != nil {
			fmt.Println(err)
			return
		}
	}

}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.airbnb-backend.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringVarP(&database, "database", "d", "false", "This is Command's Persistent Flag")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
