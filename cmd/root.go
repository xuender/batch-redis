/*
Copyright © 2019 妙音 <xuender@139.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	_hostname = "redis.hostname"
	_port     = "redis.port"
	_password = "redis.password"
	_db       = "redis.db"
	_timeout  = "redis.timeout"
	_count    = "redis.count"
)

var _cfgFile string

var rootCmd = &cobra.Command{
	Use:   "batch-redis",
	Short: "redis batch operation",
	Long: `redis count or redis delete. For example:
batch-redis count prefix*
batch-redis del prefix*`,
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

	pflags := rootCmd.PersistentFlags()
	pflags.StringVarP(&_cfgFile, "config", "c", "", "config file (default $HOME/.batch-redis.yaml)")
	pflags.StringP(_hostname, "o", "127.0.0.1", "Server hostname")
	pflags.Int32P(_port, "p", 6379, "Server port")
	pflags.StringP(_password, "a", "", "Password to use when connecting to the server")
	pflags.Int32P(_db, "n", 0, "Database number")
	pflags.Int32P(_timeout, "t", 10, "Timeout for socket reads")
	pflags.Int32P(_count, "u", 1000, "SCAN count")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if _cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(_cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".batch-redis" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".batch-redis")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
