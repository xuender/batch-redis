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
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete key by pattern.",
	Long: `Delete key by pattern. For example:

batch-redis del prefix*`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			panic(errors.New("Defect a SCAN pattern"))
		}
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
			panic(err)
		}
		fmt.Printf("Del pattern: %s\n", args[0])
		iter := client.Scan(0, args[0], int64(c.Int(_count))).Iterator()
		count := 0
		keys := make([]string, 100)
		for iter.Next() {
			key := iter.Val()
			if len(keys) > 99 {
				if err := client.Del(keys...).Err(); err != nil {
					panic(err)
				}
				keys = append([]string{})
				fmt.Printf("%d: del %s\n", count, key)
			}
			count++
			keys = append(keys, key)
		}
		if len(keys) > 0 {
			if err := client.Del(keys...).Err(); err != nil {
				panic(err)
			}
		}
		if err := iter.Err(); err != nil {
			panic(err)
		}
		fmt.Printf("Del %s matching %d\n", args[0], count)
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
