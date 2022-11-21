// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"path/filepath"

	"github.com/adrg/xdg"
)

// optionsConfig is where the options are configured.
type optionsConfig struct {
	// pluginRoot is the directory that contains the plugins.
	pluginRoot string

	// gcpBucket is the bucket name for the gcp artifact repository.
	gcpBucket string

	// gcpRootPath is the root bucket path for the gcp artifact repository.
	gcpRootPath string

	// repoName is the repository name.
	repoName string

	// VersionSelector is the means to find versions of plugins in a repository.
	versionSelector VersionSelector
}

var (
	// DefaultPluginRoot is the default plugin root.
	DefaultPluginRoot = filepath.Join(xdg.DataHome, "tanzu-cli")
)

// makeDefaultOptions creates the default options for this namespace.
func makeDefaultOptions(list ...Option) optionsConfig {
	opts := optionsConfig{
		// by default, the plugin root is at $XDG_DATA_HOME/tanzu-cli
		pluginRoot: DefaultPluginRoot,
		/*
			gcpRootPath:     DefaultArtifactsDirectory,
		*/
		versionSelector: DefaultVersionSelector,
	}

	for _, o := range list {
		o(&opts)
	}

	return opts
}

// Option is a filesystem store option.
type Option func(o *optionsConfig)

// WithPluginRoot sets the root which directory plugins live in.
func WithPluginRoot(root string) Option {
	return func(o *optionsConfig) {
		o.pluginRoot = root
	}
}

// WithGCPBucket sets the gcp bucket to use for the artifact repository.
func WithGCPBucket(name string) Option {
	return func(o *optionsConfig) {
		o.gcpBucket = name
	}
}

// WithGCPRootPath sets the gcp bucket root path to use for the artifact repository.
func WithGCPRootPath(path string) Option {
	return func(o *optionsConfig) {
		o.gcpRootPath = path
	}
}

// WithName sets the name
func WithName(name string) Option {
	return func(o *optionsConfig) {
		o.repoName = name
	}
}

// WithVersionSelector sets the version finder.
func WithVersionSelector(finder VersionSelector) Option {
	return func(o *optionsConfig) {
		o.versionSelector = finder
	}
}
