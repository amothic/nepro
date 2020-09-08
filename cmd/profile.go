package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(profileCmd)
	profileCmd.Flags().StringVarP(&url, "url", "u", "", "neptune url")
	profileCmd.Flags().StringVarP(&query, "query", "q", "", "graph query")
	profileCmd.MarkFlagRequired("url")
	profileCmd.MarkFlagRequired("query")
}

var profileCmd = &cobra.Command{
	Use: "profile",
	RunE: func(cmd *cobra.Command, args []string) error {
		endpoint := url + "/gremlin/profile"
		resp, err := Request(endpoint, query)
		if err != nil {
			return err
		}
		fmt.Println(string(resp))
		return nil
	},
}
