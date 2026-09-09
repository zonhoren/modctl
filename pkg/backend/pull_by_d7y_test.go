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

package backend

import (
	"testing"

	"github.com/modelpack/modctl/pkg/config"

	"github.com/stretchr/testify/assert"
)

// TestBuildDownloadTaskRequestForceHardLink proves ForceHardLink threads
// straight through to the Dragonfly gRPC request in both directions --
// no live dfdaemon connection required, since buildDownloadTaskRequest is
// pure. See config.Pull.ForceHardLink / config.Fetch.ForceHardLink and
// pull_by_d7y.go / fetch_by_d7y.go, both of which call this helper instead
// of hardcoding the field.
func TestBuildDownloadTaskRequestForceHardLink(t *testing.T) {
	tests := []struct {
		name          string
		forceHardLink bool
	}{
		{name: "false stays false", forceHardLink: false},
		{name: "true threads through", forceHardLink: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := buildDownloadTaskRequest("https://example.com/v2/repo/blobs/sha256:abc", "/tmp/out", "Basic dGVzdA==", tt.forceHardLink)
			assert.Equal(t, tt.forceHardLink, req.Download.ForceHardLink)
		})
	}
}

// TestForceHardLinkZeroValue proves an unset config.Pull/config.Fetch (the
// zero value, e.g. a struct literal or the constructor defaults) keeps
// ForceHardLink false -- backward compatibility, not just the CLI flag's
// own default.
func TestForceHardLinkZeroValue(t *testing.T) {
	assert.False(t, config.Pull{}.ForceHardLink)
	assert.False(t, config.Fetch{}.ForceHardLink)
	assert.False(t, config.NewPull().ForceHardLink)
	assert.False(t, config.NewFetch().ForceHardLink)
}
