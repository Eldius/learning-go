/*
Package cmd is where commands live...
*/
package cmd

import (
	"path/filepath"
	"fmt"
	"log"
	"github.com/spf13/cobra"
	"os"
	"net"
	"net/http"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"github.com/gorilla/handlers"
)

var cfgFile string


var buildTime string
var commitHash string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "file-server",
	Short: "An HTTP server to share files from disk",
	Long: fmt.Sprintf(`An HTTP server to share files from disk

build time: %s
%s
`, buildTime, commitHash),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		//mux := http.NewServeMux()
		//fs := http.FileServer(http.Dir(getRootPath(serverPath)))
		//mux.Handle("/", fs)
		//mux.Handle("/", fs)
		//http.ListenAndServe(fmt.Sprintf(":%d", serverPort) , mux)
		//http.ListenAndServe(fmt.Sprintf(":%d", serverPort) , handlers.LoggingHandler(os.Stdout, mux))

		path := getRootPath(serverPath)
		fs := http.FileServer(http.Dir(path))
		r := handlers.LoggingHandler(log.Writer(), fs)
		printServerStartedLog(path, serverPort)
		if *serverCompress {
			r = handlers.CompressHandler(r)
		}
		http.Handle("/", r)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil))
	 },
}

func printServerStartedLog(path string, serverPort int) {
	msg := fmt.Sprintf("\n\n---\nServing %s on HTTP at:\n", path)
	for _, ip := range getIPAddress() {
		msg += fmt.Sprintf("- http://%s:%d\n", ip, serverPort)
	}
	msg += "---\n"
	log.Println(msg)
}

func getIPAddress() []string {
	results := make([]string, 0)
	ifaces, err := net.Interfaces()
	// handle err
	if err != nil {

	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Panic(err.Error())
		}
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
					ip = v.IP
			case *net.IPAddr:
					ip = v.IP
			}
			results = append(results, ip.String())
		}
	}

	return results
}

func getRootPath(serverPath string) string {
	log.Println("parsing root path", serverPath)
	if path, err := filepath.Abs(serverPath); err != nil {
		panic(err.Error())
	} else {
		log.Println("returning", path)
		return path
	}
}

var serverPort int
var serverPath string
var serverCompress *bool

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
	rootCmd.Flags().IntVarP(&serverPort, "port", "p", 8080, "-p 8080")
	rootCmd.Flags().StringVarP(&serverPath, "folder", "f", ".", "-p 8080")

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.file-server.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	serverCompress = rootCmd.Flags().BoolP("compress", "c", false, "Compress data")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".file-server" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".file-server")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
