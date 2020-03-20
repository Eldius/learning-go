/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"time"
	"io"
	"path/filepath"

	"github.com/spf13/cobra"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename files",
	Long: `Rename files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rename called")
		proccess(source, destination, minutesAhead)
	},
}

var source string
var destination string
var minutesAhead int64

var timeFormatPattern = "20060102150405"

func init() {
	rootCmd.AddCommand(renameCmd)

	renameCmd.Flags().StringVarP(&source, "source", "s", ".", "Files source folder")
	renameCmd.Flags().StringVarP(&destination, "dest", "d", ".", "Files destination folder")
	renameCmd.Flags().Int64VarP(&minutesAhead, "minutes", "m", 0, "Minutes to add")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func proccess(source, destination string, minutesAhead int64) {
	var files []string
	createDestinationFolder(destination)
    err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if (!info.IsDir()) {
			files = append(files, path)
		}
        return nil
    })
    if err != nil {
        panic(err)
    }
    for i, file := range files {
			sourceFile, err := os.Open(file)
			if err != nil {
				panic(err.Error())
			}
			defer sourceFile.Close()
			destFileWithTimestamp := createFileNameWithTimestamp(minutesAhead, i, destination)
			destinationFile, err := os.Create(destFileWithTimestamp)
			if err != nil {
                panic(err.Error())
        	}
        	defer destinationFile.Close()
			fmt.Println(fmt.Sprintf("copying file '%s' => '%s'", file, destFileWithTimestamp))
			io.Copy(sourceFile, destinationFile)
    }
}

func createFileNameWithTimestamp(minutesAhead int64, secondsToAdd int, destination string) string {
	fileName, err := filepath.Abs(fmt.Sprintf(
		"%s/OHB.%s",
		destination,
		time.Now().Local().Add(time.Minute * time.Duration(minutesAhead)).Add(time.Second * time.Duration(secondsToAdd)).Format(timeFormatPattern),
	))
	if err != nil {
		panic(err.Error())
	}
	return fileName
}

func createDestinationFolder(destFolder string) {
	os.MkdirAll(destination, os.ModePerm)
}
