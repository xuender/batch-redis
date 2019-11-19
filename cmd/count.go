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
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/spf13/cobra"
)

const (
	_pattern = "pattern"
)

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Defect a SCAN pattern.")
		} else {
			c := NewCommand(cmd)
			client := redis.NewClient(&redis.Options{
				Addr:        fmt.Sprintf("%s:%d", c.String(_hostname), c.Int(_port)),
				Password:    c.String(_password),
				DB:          c.Int(_db),
				ReadTimeout: time.Second * time.Duration(c.Int(_timeout)),
			})
			defer client.Close()
			fmt.Printf("Connection %s:%d...\n", c.String(_hostname), c.Int(_port))
			if pong, err := client.Ping().Result(); err == nil {
				fmt.Println(pong)
			} else {
				fmt.Println(err.Error())
			}
			val, err := client.Get("key").Result()
			if err != nil {
				fmt.Println(err.Error())
				panic(err)
			} else {
				fmt.Println(val)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(countCmd)
}
