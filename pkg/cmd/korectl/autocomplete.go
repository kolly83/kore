/**
 * Copyright 2020 Appvia Ltd <info@appvia.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package korectl

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var (
	autocompleteLongDescription = `
Provides bash autocompletion to make working with korectl quickier.

# View the source code
$ korectl autocomplete bash

# Source into the shell (which can be placed into your ${HOME}/.bash_profile
$ source <(korectl autocomplete bash)
`
)

// GetAutoCompleteCommand returns the autocomplete command
func GetAutoCompleteCommand(config *Config) *cli.Command {
	return &cli.Command{
		Name:        "autocomplete",
		Usage:       "Provides the autocomplete output so you can source into your bash shell",
		Description: formatLongDescription(autocompleteLongDescription),

		Subcommands: []*cli.Command{
			{
				Name:  "bash",
				Usage: "generates the bash auto completion",
				Action: func(ctx *cli.Context) error {
					return printAutocompletion(ctx.Command.Name)
				},
			},
			{
				Name:  "zsh",
				Usage: "generates the zsh autocompletion",
				Action: func(ctx *cli.Context) error {
					return printAutocompletion(ctx.Command.Name)
				},
			},
		},
	}
}

// printAutocompletion prints the source code
func printAutocompletion(name string) error {
	if code, found := completions[name]; found {
		fmt.Println(code)
	}

	return nil
}
