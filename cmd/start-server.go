package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"shorturl/internal/server"
)

// start-serverCmd represents the start-server command
var startServerCmd = &cobra.Command{
	Use:   "start-server",
	Short: "Starts the server required for the short links to work",
	Long: `Starts an HTTP server which allows you to use the short links
both using the CLI and the browser (using localhost:3000/s/ for making
short links and localhost:3000/u/ for getting unrolled versions"`,
	Run: func(cmd *cobra.Command, args []string) {
		s := server.ShortURLServer{Address: "localhost:3000"}
		s.Init()
		fmt.Println("Started the server at http://localhost:3000")
		log.Fatal(s.Run())
	},
}

func init() {
	rootCmd.AddCommand(startServerCmd)
}
