package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/url"
	"shorturl/pkg/shortener"
	"shorturl/pkg/utils"
)

// shortCmd represents the short command
var shortCmd = &cobra.Command{
	Use:   "short",
	Short: "Makes a short link from the given URL",
	Long: `Makes a short link from the given URL, requires the server to be running.
			Look at the server command which explains how to start the server.
			Without it you won't be able to use the given link'`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Please provide a link as a single argument.")
		}

		u, err := url.ParseRequestURI(args[0])
		if err != nil {
			panic(err)
		}

		err = utils.CheckServerState()
		if err != nil {
			fmt.Printf("Server is not running, error: %s\n", err)
		}

		storage := utils.StartUpStorage()

		short, full := shortener.Make(u.String())

		storage.Store(u.String(), short)
		fmt.Printf("Your shortened link: %s", full)
	},
}

func init() {
	rootCmd.AddCommand(shortCmd)
}
