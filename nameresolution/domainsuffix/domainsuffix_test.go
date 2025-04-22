/*
Copyright 2024 The Dapr Authors
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

package domainsuffix

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	nr "github.com/dapr/components-contrib/nameresolution"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name          string
		metadata      nr.Metadata
		expectedError string
	}{
		{
			name: "valid metadata with domain suffix",
			metadata: nr.Metadata{
				Configuration: map[string]string{
					"domainSuffix": "example.dev",
				},
			},
		},
		{
			name: "valid metadata with domain suffix with leading dot",
			metadata: nr.Metadata{
				Configuration: map[string]string{
					"domainSuffix": ".example.dev",
				},
			},
		},
		{
			name: "missing domain suffix",
			metadata: nr.Metadata{
				Configuration: map[string]string{},
			},
			expectedError: "domainSuffix is required in metadata",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewDomainSuffixResolver()
			err := r.Init(context.Background(), tt.metadata)

			if tt.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestResolveID(t *testing.T) {
	tests := []struct {
		name           string
		domainSuffix   string
		request        nr.ResolveRequest
		expectedResult string
		expectedError  string
	}{
		{
			name:         "valid app name",
			domainSuffix: "-my.example.dev",
			request: nr.ResolveRequest{
				ID: "some-app",
			},
			expectedResult: "some-app-my.example.dev",
		},
		{
			name:         "valid app name with leading dot in suffix",
			domainSuffix: ".example.dev",
			request: nr.ResolveRequest{
				ID: "some-app",
			},
			expectedResult: "some-app.example.dev",
		},
		{
			name:         "empty app name",
			domainSuffix: "example.dev",
			request: nr.ResolveRequest{
				ID: "",
			},
			expectedError: "empty ID not allowed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewDomainSuffixResolver()
			err := r.Init(context.Background(), nr.Metadata{
				Configuration: map[string]string{
					"domainSuffix": tt.domainSuffix,
				},
			})
			require.NoError(t, err)

			result, err := r.ResolveID(context.Background(), tt.request)

			if tt.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}
		})
	}
}

func TestClose(t *testing.T) {
	r := NewDomainSuffixResolver()
	err := r.Close()
	require.NoError(t, err)
} 