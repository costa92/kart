/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

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
	"github.com/fatih/color"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	flagHelp          = "help"
	flagHelpShorthand = "H"
)

func init() {
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
	pflag.BoolP(
		flagHelp,
		flagHelpShorthand,
		false,
		fmt.Sprintf("Help for the %s command.", color.GreenString(strings.Split("kart", " ")[0])),
	)
}

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "kart 帮助命令",
	Long: `Help provides help for any command in the application.
Simply type kart help [path to command] for full details.`,
	Run: func(c *cobra.Command, args []string) {
		cmd, _, e := c.Root().Find(args)
		if cmd == nil || e != nil {
			c.Printf("Unknown help topic %#q\n", args)
			_ = c.Root().Usage()
		} else {
			cmd.InitDefaultHelpFlag() // make possible 'help' flag to be shown
			_ = cmd.Help()
		}
	},
}
