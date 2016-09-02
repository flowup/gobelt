// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/flowup/gobelt/operatorgen"
	"github.com/flowup/gogen"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		watcher, err := fsnotify.NewWatcher()
		panicIf(err)
		path, err := os.Getwd()
		panicIf(err)

		watcher.Add(path)

		for {
			select {
			case ev := <-watcher.Events:
				fmt.Println("Triggered generator for ", ev.Name)
				build, err := gogen.ParseFile(ev.Name)
				if err != nil {
					fmt.Println("Build error occured")
					return
				}

				fileBuild := build.Files[filepath.Base(ev.Name)]
				targetDir := filepath.Dir(path)

				if err = operatorgen.FromFile(fileBuild, targetDir); err != nil {
					fmt.Println("Error occured within generator: ", err.Error())
				}
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(watchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// watchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// watchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
