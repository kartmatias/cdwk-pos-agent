/*
Copyright © 2024 Carlos Matias - carlos@codework.com.br
*/
package cmd

import (
	"fmt"

	"github.com/kartmatias/cdwk-pos-agent/api"
	"github.com/kartmatias/cdwk-pos-agent/database"
	"github.com/spf13/cobra"
)

// getproductCmd represents the getproduct command
var getproductCmd = &cobra.Command{
	Use:   "getproduct",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		productId := args[0]

		api.GetProduct(rootLogger, productId)

		database.Open(rootLogger)
		text, _ := database.RetrieveProduct("AE004")
		fmt.Printf("Product: %s", text)
	},
}

func init() {
	rootCmd.AddCommand(getproductCmd)
}
