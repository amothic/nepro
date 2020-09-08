package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(explainCmd)
	explainCmd.Flags().StringVarP(&url, "url", "u", "", "neptune url")
	explainCmd.Flags().StringVarP(&query, "query", "q", "", "graph query")
	explainCmd.MarkFlagRequired("url")
	explainCmd.MarkFlagRequired("query")
}

var explainCmd = &cobra.Command{
	Use: "explain",
	RunE: func(cmd *cobra.Command, args []string) error {
		endpoint := url + "/gremlin/explain"
		resp, err := Request(endpoint, query)
		if err != nil {
			return err
		}
		fmt.Println(resp)
		return nil
	},
}
