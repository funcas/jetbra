/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

// productsCmd represents the products command
var productsCmd = &cobra.Command{
	Use:   "products",
	Short: "print all supported products",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		products := []string{}
		for k, _ := range LIC_TEMPLATE {
			if k == "ALL" {
				continue
			}
			products = append(products, k)
		}
		sort.Strings(products)
		productStr := strings.Join(products, ",")
		fmt.Println("support products are: " + productStr)
	},
}

func init() {
	rootCmd.AddCommand(productsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// productsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// productsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
