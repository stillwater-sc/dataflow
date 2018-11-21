package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
)

var portNr string

var serverCmd = &cobra.Command{
	Use: "server",
	Short: "start data flow server",
	Run: serverExecute,
}

func init() {
	serverCmd.PersistentFlags().StringVarP(&portNr, "portNr", "p", "9090", "the port number the server listens to")
	rootCmd.AddCommand(serverCmd)
}

func work(w http.ResponseWriter, r *http.Request) {
	log.Println("Work to be done at this URL")
	w.Write([]byte("OK"))
}

func serverExecute(cmd *cobra.Command, args []string) {
	workHandler := http.HandlerFunc(work)

	http.Handle("/", workHandler)
	if err := http.ListenAndServe(":"+portNr, nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}