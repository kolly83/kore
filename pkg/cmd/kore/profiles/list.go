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

package profiles

import (
	cmdutil "github.com/appvia/kore/pkg/cmd/utils"
	"github.com/appvia/kore/pkg/utils/render"

	"github.com/spf13/cobra"
)

type ListOptions struct {
	cmdutil.Factory
	cmdutil.DefaultHandler
	// Headers indicates no headers on the table output
	Headers bool
	// Output is the output format
	Output string
}

// NewCmdProfilesList creates and returns the profile list command
func NewCmdProfilesList(factory cmdutil.Factory) *cobra.Command {
	o := &ListOptions{Factory: factory}

	return &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "lists all the profiles which has been configured thus far",
		Run:     cmdutil.DefaultRunFunc(o),
	}
}

// Run implements the action
func (o *ListOptions) Run() error {
	type profile struct {
		Name   string `json:"name"`
		Server string `json:"server"`
		Team   string `json:"team"`
	}
	var list []profile

	for k, v := range o.Config().Profiles {
		if o.Config().HasServer(v.Server) {
			list = append(list, profile{
				Name:   k,
				Server: o.Config().Servers[v.Server].Endpoint,
				Team:   v.Team,
			})
		}
	}
	current := o.Config().CurrentProfile

	return render.Render().
		Writer(o.Writer()).
		Format(o.Output).
		ShowHeaders(o.Headers).
		Resource(render.FromStruct(&list)).
		Printer(
			render.Column("Profile", "name"),
			render.Column("Endpoint", "server"),
			render.Column("Default Team", "team", render.Default("None")),
			render.Column("Active", "name", render.IfEqualOr(current, "yes", `-`)),
		).Do()
}
