package cmd

import (
	"fmt"
	"log"
	"shorturl/internal/server"

	"github.com/spf13/cobra"
)

// start-serverCmd represents the start-server command
var startServerCmd = &cobra.Command{
	Use:   "start-server",
	Short: "Starts the server required for the short links and connects to mongoDB.",
	Long: `Starts an HTTP server which allows you to use the short links
		   using the browser (using localhost:3000/s/ for making
  		   short links and localhost:3000/u/ for getting unrolled versions
		   Connects to MongoDB if the argument provides the right password."`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			s := server.ShortURLServer{Address: "localhost:3000", DBKey: args[0]}
			s.Init()
			fmt.Println("Started the server at http://localhost:3000")
			log.Fatal(s.Run())
		} else {
			fmt.Println("Please provide mongodb key as an argument, or no arguments for the local file version.")
		}
	},
}

func init() {
	rootCmd.AddCommand(startServerCmd)
}
