package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"shorturl/pkg/utils"
	"strings"
)

// unrollCmd represents the unroll command
var unrollCmd = &cobra.Command{
	Use:   "unroll",
	Short: "Unrolls a shorted link",
	Long: `Unrolls a given short link that it gets from the storage,
which can be used if the server is up.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Please provide a link as a single argument.")
			return
		}

		err := utils.CheckServerState()
		if err != nil {
			fmt.Printf("Server is not running, error: %s\n", err)
		}

		storage := utils.StartUpStorage()

		str := args[0]
		if strings.Contains(str, "/u/") {
			i := strings.Index(str, "/u/")
			str = str[i+3:]
		}

		link, err := storage.Get(str)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Your link: %s\n", link)
	},
}

func init() {
	rootCmd.AddCommand(unrollCmd)
}
