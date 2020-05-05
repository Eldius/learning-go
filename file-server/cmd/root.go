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
	"net/http"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"github.com/gorilla/handlers"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "file-server",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		//mux := http.NewServeMux()
		//fs := http.FileServer(http.Dir(getRootPath(serverPath)))
		//mux.Handle("/", fs)
		//mux.Handle("/", fs)
		//http.ListenAndServe(fmt.Sprintf(":%d", serverPort) , mux)
		//http.ListenAndServe(fmt.Sprintf(":%d", serverPort) , handlers.LoggingHandler(os.Stdout, mux))

		fs := http.FileServer(http.Dir(getRootPath(serverPath)))
		http.Handle("/", handlers.LoggingHandler(log.Writer(), fs))
		log.Printf("Serving %s on HTTP port: %d\n", getRootPath(serverPath), serverPort)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil))
	 },
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
