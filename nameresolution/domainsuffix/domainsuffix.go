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
	"fmt"
	"reflect"

	"github.com/dapr/components-contrib/metadata"
	nr "github.com/dapr/components-contrib/nameresolution"
	kitmd "github.com/dapr/kit/metadata"
)

type DomainSuffixResolver struct {
	domainSuffix string
}

type domainSuffixMetadata struct {
	DomainSuffix string `mapstructure:"domainSuffix"`
}

func NewDomainSuffixResolver() nr.Resolver {
	return &DomainSuffixResolver{}
}

func (r *DomainSuffixResolver) Init(ctx context.Context, metadata nr.Metadata) error {
	var meta domainSuffixMetadata
	err := kitmd.DecodeMetadata(metadata.Configuration, &meta)
	if err != nil {
		return fmt.Errorf("failed to decode metadata: %w", err)
	}

	if meta.DomainSuffix == "" {
		return fmt.Errorf("domainSuffix is required in metadata")
	}

	// Store the domain suffix
	r.domainSuffix = meta.DomainSuffix

	return nil
}

func (r *DomainSuffixResolver) ResolveID(ctx context.Context, req nr.ResolveRequest) (string, error) {
	if req.ID == "" {
		return "", fmt.Errorf("empty ID not allowed")
	}

	return fmt.Sprintf("%s%s", req.ID, r.domainSuffix), nil
}

// Close implements io.Closer
func (r *DomainSuffixResolver) Close() error {
	return nil
}

// GetComponentMetadata returns the metadata of the component
func (r *DomainSuffixResolver) GetComponentMetadata() (metadataInfo metadata.MetadataMap) {
	metadataStruct := domainSuffixMetadata{}
	metadata.GetMetadataInfoFromStructType(reflect.TypeOf(metadataStruct), &metadataInfo, metadata.NameResolutionType)
	return
} 