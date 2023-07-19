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
package main

// https://github.com/marmotedu/iam/blob/master/pkg/app/app.go
import "kart-io/kart/internal/cmd"

type Config struct {
	Server Server `json:"server,omitempty" yaml:"server" mapstructure:"server" toml:"server"`
}

type Server struct {
	Port string `json:"port,omitempty" yaml:"port" mapstructure:"port" toml:"port"`
}

type App struct {
	Commands []*cmd.Command
}

func main() {
	cmdRoot := cmd.BuildCommand(true)
	err := cmdRoot.Execute()
	if err != nil {
		return
	}
}
