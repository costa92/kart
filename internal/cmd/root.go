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
	"github.com/spf13/cobra"
	"os"
	"runtime"
	"strings"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:                "kart",
	Short:              "kart 命令",
	Long:               `kart 框架提供的命令行工具，使用这个命令行工具能很方便执行框架自带命令，也能很方便编写业务命令`,
	DisableSuggestions: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	// 不需要出现cobra默认的completion子命令
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
}

func addConfigFlag() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.config.yaml | config.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func BuildCommand(noConfig bool) *cobra.Command {
	// cmd.SetUsageTemplate(usageTemplate)
	rootCmd.SetOut(os.Stdout)
	rootCmd.SetErr(os.Stderr)
	rootCmd.Flags().SortFlags = true
	InitFlags(rootCmd.Flags())

	if !noConfig { // 是否配置文件
		addConfigFlag()
	}
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(helpCommand(FormatBaseName("kart")))

	return rootCmd
}

// FormatBaseName is formatted as an executable file name under different
// operating systems according to the given name.
func FormatBaseName(basename string) string {
	// Make case-insensitive and strip executable suffix if present
	if runtime.GOOS == "windows" {
		basename = strings.ToLower(basename)
		basename = strings.TrimSuffix(basename, ".exe")
	}
	return basename
}
