package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kevingentile/keyscrape/core"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", "", "path to search directory")
	rootCmd.PersistentFlags().StringVarP(&key, "key", "k", "", "search key")
	rootCmd.MarkFlagRequired("path")
	rootCmd.MarkFlagRequired("key")
}

var (
	key, path string
)

var rootCmd = &cobra.Command{
	Use: "Scrape a directory for the specified key",
	Run: func(cmd *cobra.Command, args []string) {
		s, err := core.NewDefaultScraper(key, path)
		if err != nil {
			fmt.Println(err)
			cmd.Usage()
			os.Exit(1)
		}
		r, err := s.Scrape()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		rb, err := json.MarshalIndent(r, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(string(rb))
	},
}

func Execute() error {
	return rootCmd.Execute()
}
