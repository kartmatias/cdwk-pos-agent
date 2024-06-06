/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/kartmatias/cdwk-pos-agent/api"
	"github.com/spf13/cobra"
)

// syncproductsCmd represents the syncproducts command
var syncproductsCmd = &cobra.Command{
	Use:   "syncproducts",
	Short: "Start product syncronization",
	Long: `Start product Sync, example:

cdwk-pos-agent syncproducts`,
	Run: func(cmd *cobra.Command, args []string) {
		api.SyncProducts(rootLogger)
		api.SyncOrders(rootLogger)
	},
}

func init() {
	rootCmd.AddCommand(syncproductsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// syncproductsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// syncproductsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
