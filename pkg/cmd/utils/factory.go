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

package utils

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/appvia/kore/pkg/client"
	"github.com/appvia/kore/pkg/client/config"
	"github.com/appvia/kore/pkg/cmd/errors"
)

type factory struct {
	client    client.Interface
	streams   Streams
	cfg       config.Config
	resources Resources
}

// NewFactory returns a default factory
func NewFactory(client client.Interface, streams Streams, config config.Config) (Factory, error) {
	resources, err := newResourceManager(client)
	if err != nil {
		return nil, err
	}

	return &factory{
		cfg:       config,
		client:    client,
		resources: resources,
		streams:   streams,
	}, nil
}

// Client returns the api client
func (f *factory) Client() client.RestInterface {
	return f.client.Request()
}

// CheckError handles the cli errors for us
func (f *factory) CheckError(kerror error) {
	err := func() error {
		if client.IsNotAuthentication(kerror) {

			return errors.ErrAuthentication
		}

		return kerror
	}()
	if err != nil {
		fmt.Fprintf(f.Stderr(), "Error: %s\n", err)

		os.Exit(1)
	}
}

// UpdateConfig is responsible for updating the configuration
func (f *factory) UpdateConfig() error {
	return config.UpdateConfig(&f.cfg, config.GetClientConfigurationPath())
}

// Config returns the factory client configuration
func (f *factory) Config() *config.Config {
	return &f.cfg
}

// SetStdin allows you to set the stdin for the factory
func (f *factory) SetStdin(in io.Reader) {
	f.streams.Stdin = in
}

// Stdin return the standard input
func (f *factory) Stdin() io.Reader {
	return f.streams.Stdin
}

// Stderr returns the io.Writer for errors
func (f *factory) Stderr() io.Writer {
	return f.streams.Stderr
}

// Writer returns the io.Writer for output
func (f *factory) Writer() io.Writer {
	return f.streams.Stdout
}

// Printf writes a message to the io.Writer
func (f *factory) Printf(message string, args ...interface{}) {
	fmt.Fprintf(f.Writer(), message, args...)
}

// Println writes a message to the io.Writer
func (f *factory) Println(message string, args ...interface{}) {
	filtered := strings.TrimRight(message, "\n")

	fmt.Fprintf(f.Writer(), filtered+"\n", args...)
}

// Resources returns the resources contract
func (f *factory) Resources() Resources {
	return f.resources
}