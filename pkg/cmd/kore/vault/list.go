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

package vault

import (
	"fmt"

	cmdutil "github.com/appvia/kore/pkg/cmd/utils"
	"github.com/appvia/kore/pkg/utils/render"

	"github.com/spf13/cobra"
)

// GetListOptions the are the options for a list command
type GetListOptions struct {
	cmdutil.Factory
	cmdutil.DefaultHandler
	// Output is the output format
	Output string
	// Headers indicates no headers on the table output
	Headers bool
}

// NewCmdVaultList creates and returns the vault list command
func NewCmdVaultList(factory cmdutil.Factory) *cobra.Command {
	o := &GetListOptions{Factory: factory}

	command := &cobra.Command{
		Use:     "list",
		Short:   "Returns all stored values within Kore's vault",
		Example: "kore vault list [options]",
		Run:     cmdutil.DefaultRunFunc(o),
	}

	return command
}

// Run implements the command
func (o *GetListOptions) Run() error {
	resp := o.ClientWithEndpoint("/vault").Get()
	fmt.Println(resp.Body())
	if resp.Error() != nil {
		return resp.Error()
	}

	return render.Render().
		Writer(o.Writer()).
		Resource(render.FromReader(resp.Body())).
		Printer(
			render.Column("Username", "username"),
			render.Column("Email", "email"),
			render.Column("Teams", "teams|@sjoin"),
		).Do()
}
