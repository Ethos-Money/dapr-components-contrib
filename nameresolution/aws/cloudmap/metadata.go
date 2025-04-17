/*
Copyright 2021 The Dapr Authors
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

package cloudmap

import (
	"errors"
	"fmt"
	"time"
)

type cloudMapMetadata struct {
	// AWS Auth (handled by common auth)
	AccessKey    string `json:"accessKey" mapstructure:"accessKey" mdignore:"true"`
	SecretKey    string `json:"secretKey" mapstructure:"secretKey" mdignore:"true"`
	SessionToken string `json:"sessionToken" mapstructure:"sessionToken" mdignore:"true"`
	Region       string `json:"region" mapstructure:"region" mapstructurealiases:"awsRegion" mdignore:"true"`

	// CloudMap Specific
	Endpoint      string `json:"endpoint"`      // Optional endpoint for testing
	NamespaceName string `json:"namespaceName"` // CloudMap namespace name
	NamespaceID   string `json:"namespaceId"`   // Optional: CloudMap namespace ID (if name not provided)

	// Optional Features
	EnableCaching bool   `json:"enableCaching"` // Enable/disable address caching
	CacheTTL      string `json:"cacheTTL"`      // Cache TTL duration (default: "60s")
	MaxCacheSize  int    `json:"maxCacheSize"`  // Maximum number of cached entries (default: 1000)
}

const (
	defaultCacheTTL     = time.Second * 60
	defaultMaxCacheSize = 1000
)

func (m *cloudMapMetadata) Validate() error {
	if m.NamespaceName == "" && m.NamespaceID == "" {
		return errors.New("either namespaceName or namespaceId must be provided")
	}

	if m.EnableCaching {
		if m.CacheTTL == "" {
			m.CacheTTL = "60s"
		}
		if _, err := time.ParseDuration(m.CacheTTL); err != nil {
			return fmt.Errorf("invalid cacheTTL format: %w", err)
		}
		if m.MaxCacheSize <= 0 {
			m.MaxCacheSize = defaultMaxCacheSize
		}
	}

	return nil
}
