/*
Copyright © 2024 Carlos Matias - carlos@codework.com.br
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
		//api.SyncOrders(rootLogger)
	},
}

func init() {
	rootCmd.AddCommand(syncproductsCmd)
}
