/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
    "log"
    "net/http"
    "io/ioutil"

	"github.com/spf13/cobra"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "A command to insert a word into dictionary",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("insert called")
        if len(args) != 1{
            log.Fatalln(fmt.Errorf("Wrong number of arguments to 'insert'"))
        }
        fmt.Println(args)
        resp, err := http.Get("http://localhost:8080/insert/"+args[0])
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
	rootCmd.AddCommand(insertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// insertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// insertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
