// AUTO-GENERATED CODE - DO NOT EDIT
// See instructions under /codegen/README.md
// GENERATED ON 2026-03-04 19:41:02

// Package _computeoptimizerautomation provides AWS client management functions for the computeoptimizerautomation
// AWS service.
//
// The Client() is a wrapper on computeoptimizerautomation.NewFromConfig(), which creates & caches
// the client.
//
// The Delete() clears the cached client.
package _computeoptimizerautomation

import (
	"sync"

	"github.com/TouchBistro/awesome/providers"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizerautomation"
)

var cmap sync.Map

// Client builds or returns the singleton computeoptimizerautomation client for the supplied provider
// If functional options are supplied, they are passed as-is to the underlying NewFromConfig(...)
// for the corresponding client
func Client(provider providers.CredsProvider, optFns ...func(*computeoptimizerautomation.Options)) (*computeoptimizerautomation.Client, error) {

	if provider == nil {
		return nil, providers.ErrNilProvider
	}
	if _, ok := cmap.Load(provider.Key()); !ok {
		client := computeoptimizerautomation.NewFromConfig(provider.Config(), optFns...)
		cmap.Store(provider.Key(), client)
	}
	client, _ := cmap.Load(provider.Key())
	return client.(*computeoptimizerautomation.Client), nil
}

// Must wraps the _computeoptimizerautomation.Client( ) function & panics if a non-nil error is returned.
func Must(provider providers.CredsProvider, optFns ...func(*computeoptimizerautomation.Options)) *computeoptimizerautomation.Client {

	client, err := Client(provider, optFns...)
	if err != nil {
		panic(err)
	}
	return client
}

// Delete removes the cached computeoptimizerautomation client for the supplied provider; This foreces the subsequent
// calls to Client() for the same provider to recreate & return a new instnce.
func Delete(provider providers.CredsProvider) error {

	if provider == nil {
		return providers.ErrNilProvider
	}
	if _, ok := cmap.Load(provider.Key()); ok {
		cmap.Delete(provider.Key())
	}
	return nil
}

// Refresh discards the cached computeoptimizerautomation client if it exists, builds & returns a new singleton instance
func Refresh(provider providers.CredsProvider, optFns ...func(*computeoptimizerautomation.Options)) (*computeoptimizerautomation.Client, error) {

	err := Delete(provider)
	if err != nil {
		return nil, err
	}
	return Client(provider, optFns...)
}
