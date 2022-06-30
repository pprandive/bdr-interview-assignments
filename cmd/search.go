/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
    "log"
    "io/ioutil"
    "net/http"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A command to search a word in the dictionary",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")
        if len(args) != 1{
            log.Fatalln(fmt.Errorf("Wrong number of arguments to 'search'"))
        }
        fmt.Println(args)
        resp, err := http.Get("http://localhost:8080/search/"+args[0])
        if err != nil {
            log.Fatalln(err)
        }
        //We Read the response body on the line below.
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            log.Fatalln(err)
        }
        //Convert the body to type string
        sb := string(body)
        log.Printf(sb)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
