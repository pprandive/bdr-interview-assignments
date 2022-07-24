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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A command to delete a word in the dictionary",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
        if len(args) != 1{
            log.Fatalln(fmt.Errorf("Wrong number of arguments to 'delete'"))
        }
        fmt.Println(args)
        resp, err := http.Get("http://localhost:8080/delete/"+args[0])
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
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
