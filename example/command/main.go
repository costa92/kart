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

import (
	"fmt"

	"github.com/spf13/viper"
	"kart-io/kart/internal/cmd"
)

type Server struct {
	Port string `json:"port,omitempty" yaml:"port" mapstructure:"port" toml:"port"`
}

func main() {
	config := Server{}
	cmd.NewCommand()
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
	fmt.Println(config)
}
