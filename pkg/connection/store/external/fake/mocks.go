/*
 Copyright 2023 The Crossplane Authors.

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

// Package fake is a fake ExternalSecretStoreServiceClient.
package fake

import (
	"context"

	"google.golang.org/grpc"

	ess "github.com/crossplane/crossplane-runtime/apis/proto/v1alpha1"
)

// ExternalSecretStoreServiceClient is a fake ExternalSecretStoreServiceClient.
type ExternalSecretStoreServiceClient struct {
	GetSecretFn   func(context.Context, *ess.GetSecretRequest, ...grpc.CallOption) (*ess.GetSecretResponse, error)
	ApplySecretFn func(context.Context, *ess.ApplySecretRequest, ...grpc.CallOption) (*ess.ApplySecretResponse, error)
	DeleteKeysFn  func(context.Context, *ess.DeleteKeysRequest, ...grpc.CallOption) (*ess.DeleteKeysResponse, error)
	*ess.UnimplementedExternalSecretStoreServiceServer
}

// GetSecret returns the secret.
func (e *ExternalSecretStoreServiceClient) GetSecret(ctx context.Context, req *ess.GetSecretRequest, opts ...grpc.CallOption) (*ess.GetSecretResponse, error) {
	return e.GetSecretFn(ctx, req)
}

// ApplySecret applies the secret.
func (e *ExternalSecretStoreServiceClient) ApplySecret(ctx context.Context, req *ess.ApplySecretRequest, opts ...grpc.CallOption) (*ess.ApplySecretResponse, error) {
	return e.ApplySecretFn(ctx, req)
}

// DeleteKeys deletes the secret keys.
func (e *ExternalSecretStoreServiceClient) DeleteKeys(ctx context.Context, req *ess.DeleteKeysRequest, opts ...grpc.CallOption) (*ess.DeleteKeysResponse, error) {
	return e.DeleteKeysFn(ctx, req)
}
