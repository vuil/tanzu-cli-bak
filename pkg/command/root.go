// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package command provides commands
package command

import (
	"github.com/spf13/cobra"

	"github.com/vmware-tanzu/tanzu-plugin-runtime/plugin"

	"github.com/vmware-tanzu/tanzu-cli/pkg/catalog"
	"github.com/vmware-tanzu/tanzu-cli/pkg/cli"
)

type DummyPluginSupplier struct {
}

// GetInstalledPlugins returns plugins for the dummy supplier.
// TODO(vuil): delete and replace with actual implementation
func (s *DummyPluginSupplier) GetInstalledPlugins() ([]*cli.PluginInfo, error) {
	plugins := make([]*cli.PluginInfo, 0)
	pi := &cli.PluginInfo{
		Name:             "fakefoo",
		Description:      "Fake foo",
		Group:            plugin.SystemCmdGroup,
		Aliases:          []string{"ff"},
		InstallationPath: "/opt/tanzu/tmpfoo.info.cmds.goodbad.usage",
	}
	plugins = append(plugins, pi)
	return plugins, nil
}

// NewRootCmd creates a root command.
func NewRootCmd(ps catalog.PluginSupplier) (*cobra.Command, error) {
	var rootCmd = &cobra.Command{
		Use: "tanzu",
		// Don't have Cobra print the error message, the CLI will
		// print it itself in a nicer format.
		SilenceErrors: true,
	}

	uFunc := cli.NewMainUsage().UsageFunc()
	rootCmd.SetUsageFunc(uFunc)
	rootCmd.AddCommand(
		newVersionCmd(),
	)

	if ps == nil {
		ps = &DummyPluginSupplier{}
	}

	plugins, err := ps.GetInstalledPlugins()
	if err != nil {
		return nil, err
	}
	for _, plugin := range plugins {
		rootCmd.AddCommand(cli.GetCmd(plugin))
	}

	return rootCmd, nil
}

// Execute executes the CLI.
func Execute() error {
	root, err := NewRootCmd(nil)
	if err != nil {
		return err
	}
	return root.Execute()
}
