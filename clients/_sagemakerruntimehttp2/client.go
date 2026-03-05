// AUTO-GENERATED CODE - DO NOT EDIT
// See instructions under /codegen/README.md
// GENERATED ON 2026-03-04 19:41:02

// Package _sagemakerruntimehttp2 provides AWS client management functions for the sagemakerruntimehttp2
// AWS service.
//
// The Client() is a wrapper on sagemakerruntimehttp2.NewFromConfig(), which creates & caches
// the client.
//
// The Delete() clears the cached client.
package _sagemakerruntimehttp2

import (
	"sync"

	"github.com/TouchBistro/awesome/providers"
	"github.com/aws/aws-sdk-go-v2/service/sagemakerruntimehttp2"
)

var cmap sync.Map

// Client builds or returns the singleton sagemakerruntimehttp2 client for the supplied provider
// If functional options are supplied, they are passed as-is to the underlying NewFromConfig(...)
// for the corresponding client
func Client(provider providers.CredsProvider, optFns ...func(*sagemakerruntimehttp2.Options)) (*sagemakerruntimehttp2.Client, error) {

	if provider == nil {
		return nil, providers.ErrNilProvider
	}
	if _, ok := cmap.Load(provider.Key()); !ok {
		client := sagemakerruntimehttp2.NewFromConfig(provider.Config(), optFns...)
		cmap.Store(provider.Key(), client)
	}
	client, _ := cmap.Load(provider.Key())
	return client.(*sagemakerruntimehttp2.Client), nil
}

// Must wraps the _sagemakerruntimehttp2.Client( ) function & panics if a non-nil error is returned.
func Must(provider providers.CredsProvider, optFns ...func(*sagemakerruntimehttp2.Options)) *sagemakerruntimehttp2.Client {

	client, err := Client(provider, optFns...)
	if err != nil {
		panic(err)
	}
	return client
}

// Delete removes the cached sagemakerruntimehttp2 client for the supplied provider; This foreces the subsequent
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

// Refresh discards the cached sagemakerruntimehttp2 client if it exists, builds & returns a new singleton instance
func Refresh(provider providers.CredsProvider, optFns ...func(*sagemakerruntimehttp2.Options)) (*sagemakerruntimehttp2.Client, error) {

	err := Delete(provider)
	if err != nil {
		return nil, err
	}
	return Client(provider, optFns...)
}
