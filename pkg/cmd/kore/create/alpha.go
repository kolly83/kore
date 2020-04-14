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

package create

import (
	cmdutils "github.com/appvia/kore/pkg/cmd/utils"

	"github.com/spf13/cobra"
)

// NewCmdCreateAlpha creates and returns the create command
func NewCmdCreateAlpha(factory cmdutils.Factory) *cobra.Command {
	command := &cobra.Command{
		Use:                   "create",
		DisableFlagsInUseLine: true,
		Short:                 "Creates one of more resources in kore",
		Run:                   cmdutils.RunHelp,
	}

	return command
}