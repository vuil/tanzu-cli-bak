// Copyright 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package command provides commands
package command

import (
	"github.com/spf13/cobra"

	"github.com/vmware-tanzu/tanzu-cli/pkg/cli"
)

// NewRootCmd creates a root command.
func NewRootCmd() (*cobra.Command, error) {
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

	plugins, err := getAvailablePlugins()
	if err != nil {
		return nil, err
	}
	for _, plugin := range plugins {
		rootCmd.AddCommand(cli.GetCmd(plugin))
	}

	return rootCmd, nil
}

func getAvailablePlugins() ([]*cli.PluginInfo, error) {
	/*
		plugins := make([]*cliapi.PluginDescriptor, 0)
		var err error

		if config.IsFeatureActivated(config.FeatureContextAwareCLIForPlugins) {
			currentServerName := ""

			server, err := config.GetCurrentServer()
			if err == nil && server != nil {
				currentServerName = server.Name
			}

			serverPlugin, standalonePlugins, err := pluginmanager.InstalledPlugins(currentServerName)
			if err != nil {
				return nil, fmt.Errorf("find installed plugins: %w", err)
			}

			allPlugins := serverPlugin
			allPlugins = append(allPlugins, standalonePlugins...)
			for i := range allPlugins {
				plugins = append(plugins, &allPlugins[i])
			}
		} else {
			// TODO: cli.ListPlugins is deprecated: Use pluginmanager.AvailablePluginsFromLocalSource or pluginmanager.AvailablePlugins instead
			plugins, err = cli.ListPlugins()
			if err != nil {
				return nil, fmt.Errorf("find available plugins: %w", err)
			}
		}
		return plugins, nil
	*/
	plugins := make([]*cli.PluginInfo, 0)
	return plugins, nil
}

// Execute executes the CLI.
func Execute() error {
	root, err := NewRootCmd()
	if err != nil {
		return err
	}
	return root.Execute()
}
