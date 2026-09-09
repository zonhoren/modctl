/*
 *     Copyright 2025 The CNAI Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"fmt"
	"io"
	"os"
)

const (
	// defaultFetchConcurrency is the default number of concurrent fetch operations.
	defaultFetchConcurrency = 5
)

type Fetch struct {
	Concurrency       int
	PlainHTTP         bool
	Proxy             string
	Insecure          bool
	Output            string
	Patterns          []string
	DragonflyEndpoint string
	ProgressWriter    io.Writer
	DisableProgress   bool
	Hooks             PullHooks
	// ForceHardLink requires Dragonfly to hard-link the downloaded blob into
	// the output path rather than silently falling back to a copy when a
	// hard link isn't possible. See config.Pull.ForceHardLink for the full
	// rationale -- same flag, same default (false), same Dragonfly-only
	// scope, mirrored here for the Fetch path.
	ForceHardLink bool
}

func NewFetch() *Fetch {
	return &Fetch{
		Concurrency:       defaultFetchConcurrency,
		PlainHTTP:         false,
		Proxy:             "",
		Insecure:          false,
		Output:            "",
		Patterns:          []string{},
		DragonflyEndpoint: "",
		ProgressWriter:    os.Stdout,
		DisableProgress:   false,
		Hooks:             &emptyPullHook{},
		ForceHardLink:     false,
	}
}

func (f *Fetch) Validate() error {
	if f.Concurrency < 1 {
		return fmt.Errorf("invalid concurrency: %d", f.Concurrency)
	}

	if f.Output == "" {
		return fmt.Errorf("output is required")
	}

	if len(f.Patterns) == 0 {
		return fmt.Errorf("patterns are required")
	}

	// ForceHardLink only applies to the Dragonfly download path.
	if f.ForceHardLink && f.DragonflyEndpoint == "" {
		return fmt.Errorf("force hard link only can work with dragonfly endpoint scenario")
	}

	return nil
}
