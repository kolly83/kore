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
	"github.com/urfave/cli/v2"
)

const (
	localEndpoint  string = "http://127.0.0.1:10080"
	localManifests string = "./manifests/local"
	localCompose   string = "./hack/compose"
)

// GetLocalCommand return local command
func GetLocalCommand(config *Config) *cli.Command {
	commands := append([]*cli.Command{
		GetLocalConfigureSubCommand(config),
		GetLocalLogsSubCommand(config),
	}, GetLocalRunSubCommands(config)...)

	return &cli.Command{
		Name:  "local",
		Usage: "Used to configure and run a local instance of Kore.",

		Subcommands: DefaultCompletion(commands...),
	}
}
