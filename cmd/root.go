/*
Copyright © 2023 NAME HERE kartmatias@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// rootCmd represents the base command when called without any subcommands
var rootLogger *zap.Logger

var rootCmd = &cobra.Command{
	Use:   "cdwk-pos-agent",
	Short: "Eversa e-commerce API integration",
	Long: `Eversa e-commerce API integration
with Sitex ERP backend:

Connects with SQL Server from Sitex.
Generate Products, Sales, Stock and Tags.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(mylogger *zap.Logger) {
	rootLogger = mylogger
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cdwk-pos-agent.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
