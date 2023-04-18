/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

type flags struct {
	outputToFile  bool
	caseSensitive bool
}

var flag flags

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "grep-cli-golang",
	Short: "grep command line utility in golang",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// grep from stdin
		if len(args) == 1 {
			Strarr, err := grep(os.Stdin, args[0], flag)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(Strarr)
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
				Strarr, _ := grep(file, args[0], flag)
				fmt.Println(Strarr)
			}

		}

	},
}

// write test cases for this function
func grep(r io.Reader, str string, flag flags) ([]string, error) {
	arr := []string{}
	scanner := bufio.NewScanner(r)
	searchPattern, err := regexp.Compile(str)
	if err != nil {
		return nil, err
	}
	if flag.caseSensitive {
		searchPattern = regexp.MustCompile("(?i)" + str)
	}

	for scanner.Scan() {
		if searchPattern.MatchString(scanner.Text()) {
			arr = append(arr, scanner.Text())
		}

	}
	if err := scanner.Err(); err != nil {
		return nil, err
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolVarP(&flag.outputToFile, "output", "o", false, "output to file")
	rootCmd.Flags().BoolVarP(&flag.caseSensitive, "case-sensitive", "i", false, "case sensitive")
}
