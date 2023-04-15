/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "grep-cli-golang",
	Short: "grep command line utility in golang",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// grep from stdin
		if len(args) == 0 {
			searchString(os.Stdin, "abc")
		}

		// grep from files
		if len(args) > 0 {
			for _, arg := range args[1:] {
				file, err := os.Open(arg)
				if err != nil {
					fmt.Fprintf(os.Stderr, "failed to open file %q: %v\n", arg, err)
					continue // read the next file
				}
				defer file.Close()
				Strarr, _ := searchString(file, args[0])
				fmt.Println(Strarr)
			}

		}

	},
}

// write test cases for this function
func searchString(r io.Reader, str string) ([]string, error) {
	arr := []string{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), str) {
			arr = append(arr, scanner.Text())
		}

	}
	return arr, nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.grep-cli-golang.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
