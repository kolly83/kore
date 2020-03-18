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

// GetCreateGCPCommand provides the gcp create commands
func GetCreateGCPCommand(config *Config) *cli.Command {
	return &cli.Command{
		Name:    "gcp",
		Aliases: []string{"google"},
		Usage:   "Provides the ability to create GCP related cloud resources",
		Subcommands: []*cli.Command{
			GetCreateGCPOrganization(config),
			GetCreateGCPProject(config),
		},
	}
}
