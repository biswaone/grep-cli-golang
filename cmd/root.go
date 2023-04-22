/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/spf13/cobra"
)

type flags struct {
	outputToFile  bool
	caseSensitive bool
	matchCount    bool
}

var flag flags

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "grep-cli-golang",
	Short: "grep command line utility in golang",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("No arguments provided")
		}
		// grep from stdin
		if len(args) == 1 {
			searchStr := args[0]
			Strarr, err := grep(os.Stdin, searchStr, flag)
			if err != nil {
				log.Fatal(err)
			}
			if flag.outputToFile {
				outputFile := args[2]
				err := writeToFile(Strarr, outputFile)
				if err != nil {
					log.Fatal(err)
				}
			} else if flag.matchCount {
				fmt.Println(len(Strarr)) // print line count

			} else { // print to stdout
				for _, line := range Strarr {
					fmt.Println(line)
				}

			}
		}

		// grep from directory or file
		if len(args) > 1 {
			searchStr := args[0]
			path := args[1]

			fileInfo, err := os.Stat(path)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			mode := fileInfo.Mode()
			// grep from directory
			if mode.IsDir() {

				result := make(map[string][]string)

				var mu sync.Mutex
				var wg sync.WaitGroup
				err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
					if err != nil {
						fmt.Printf("Error: %v\n", err)
						return err
					}
					if !d.IsDir() {
						wg.Add(1)

						go func() {
							defer wg.Done()
							file, err := os.Open(path)
							if err != nil {
								fmt.Fprintf(os.Stderr, "failed to open file %q: %v\n", path, err)
								return
							}
							defer file.Close()
							Strarr, _ := grep(file, searchStr, flag)
							mu.Lock()
							result[path] = Strarr
							mu.Unlock()
						}()
					}
					return nil

				})
				if err != nil {
					log.Fatal(err)
				}
				wg.Wait()
				for key, value := range result {
					for _, v := range value {
						fmt.Println(key, v)
					}
				}
			}

			// grep from a file
			if !mode.IsDir() {
				file, err := os.Open(path)
				if err != nil {
					fmt.Fprintf(os.Stderr, "failed to open file %q: %v\n", path, err)
				}
				defer file.Close()
				Strarr, _ := grep(file, searchStr, flag)

				if flag.outputToFile && !flag.matchCount {
					outputFile := args[2]
					err := writeToFile(Strarr, outputFile)
					if err != nil {
						log.Fatal(err)
					}
				} else if flag.matchCount {
					fmt.Println(len(Strarr)) // print line count

				} else { //print to stdout
					for _, line := range Strarr {
						fmt.Println(line)
					}
				}

			}
		}

	},
}

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

func writeToFile(Strarr []string, filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range Strarr {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
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
	rootCmd.Flags().BoolVarP(&flag.matchCount, "match-count", "c", false, "match count")
}
