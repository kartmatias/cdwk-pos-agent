/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/kartmatias/cdwk-pos-agent/api"
	"github.com/spf13/cobra"
)

// getproductlistCmd represents the getproductlist command
var getproductlistCmd = &cobra.Command{
	Use:   "getproductlist",
	Short: "Get a product list",
	Long: `Gets a product list from
	Woocommerce server.`,
	Run: func(cmd *cobra.Command, args []string) {
		api.GetProducts(rootLogger)
	},
}

func init() {
	rootCmd.AddCommand(getproductlistCmd)
}
