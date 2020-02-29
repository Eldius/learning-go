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
	"strings"
	"github.com/james4k/rcon"
	"github.com/spf13/cobra"
)

// cmdCmd represents the cmd command
var cmdCmd = &cobra.Command{
	Use:   "cmd [command]",
	Short: "Executes a command in the server using the RCON protocol",
	Long:  `Executes a command in the server using the RCON protocol.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("executing command '%s' at %s", strings.Join(args[:], " "), cmdHost))
		executeCommand(cmdHost, cmdPort, cmdPass, args)
	},
}

var cmdHost string
var cmdPort int
var cmdPass string

func init() {
	rootCmd.AddCommand(cmdCmd)

	cmdCmd.PersistentFlags().StringVarP(&cmdHost, "server", "s", "", "The server host or IP")
	cmdCmd.PersistentFlags().IntVarP(&cmdPort, "port", "p", 0, "The server port")
	cmdCmd.PersistentFlags().StringVarP(&cmdPass, "pass", "W", "", "The server RCON password")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cmdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cmdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func executeCommand(host string, port int, pass string, cmd []string) {
	//response := rconhelper.NewRconConnection(host, port, pass).Execute(cmd)

	//fmt.Println(fmt.Sprintf("command result:\n%s", response))
	hostPort := fmt.Sprintf("%s:%d", host, port)
	fmt.Println(fmt.Sprintf("host: %s", hostPort))
	remoteConsole, err := rcon.Dial(hostPort, pass)
	if err != nil {
		panic(err.Error())
	}
	defer remoteConsole.Close()

	reqId, err := remoteConsole.Write(strings.Join(cmd[:], " "))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to send command:", err.Error())
		//continue
	}

	resp, respReqId, err := remoteConsole.Read()
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Fprintln(os.Stderr, "Failed to read command:", err.Error())
		//continue
	}

	if reqId != respReqId {
		fmt.Println("Weird. This response is for another request.")
	}

	fmt.Println(resp)
	//fmt.Fprintln(out, resp)
	//out.Write([]byte("> "))
}
