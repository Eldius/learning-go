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
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "file-renamer",
	Short: "Rename files",
	Long:  `Rename files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rename called")
		proccess(source, destination, minutesAhead)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVarP(&source, "source", "s", ".", "Files source folder")
	rootCmd.Flags().StringVarP(&destination, "dest", "d", ".", "Files destination folder")
	rootCmd.Flags().Int64VarP(&minutesAhead, "minutes", "m", 0, "Minutes to add")
}

var source string
var destination string
var minutesAhead int64

var timeFormatPattern = "20060102150405"

// initConfig reads in config file and ENV variables if set.
func initConfig() {
}

func proccess(source, destination string, minutesAhead int64) {
	var files []string
	createDestinationFolder(destination)
	err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
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
		time.Now().Local().Add(time.Minute*time.Duration(minutesAhead)).Add(time.Second*time.Duration(secondsToAdd)).Format(timeFormatPattern),
	))
	if err != nil {
		panic(err.Error())
	}
	return fileName
}

func createDestinationFolder(destFolder string) {
	os.MkdirAll(destination, os.ModePerm)
}
