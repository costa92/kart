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
	"github.com/costa92/logger"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:                "kart",
	Short:              "kart 命令",
	Long:               `kart 框架提供的命令行工具，使用这个命令行工具能很方便执行框架自带命令，也能很方便编写业务命令`,
	DisableSuggestions: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.InitDefaultHelpFlag()
		return cmd.Help()
	},
	// 不需要出现cobra默认的completion子命令
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
			viper.SetConfigType("yaml")
		} else {
			viper.SetConfigType("yaml")
			home, err := homedir.Dir()
			cobra.CheckErr(err)
			viper.AddConfigPath(home)
			viper.SetConfigName("config")
		}
		viper.AutomaticEnv()
		if err := viper.ReadInConfig(); err != nil {
			logger.Errorw("Using config file", "file", viper.ConfigFileUsed(), "err", err)
			os.Exit(1)
		}
	})
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.config.yaml | config.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(helpCmd)
	rootCmd.AddCommand(serveCmd)
}
