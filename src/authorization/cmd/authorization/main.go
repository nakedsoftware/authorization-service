// Copyright 2025 Naked Software, LLC
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package main implements Authorization Service. Authorization Service
// provides identity and authorization services to websites and resource
// services. Authorization Service implements the OAuth 2, OpenID Connect,
// and WebAuthn protocols to help client applications to identify users and
// obtain authorization to use resource services on behalf of a user or
// other client application.
package main

import (
	"errors"
	"github.com/nakedsoftware/authorization-service/src/authorization/internal/commands"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"path"
)

const environmentPrefix = "NAKEDSOFTWARE_AUTHORIZATION"

func main() {
	cobra.OnInitialize(configureViper)

	if err := commands.Execute(); err != nil {
		slog.Error("failed to execute command", "error", err.Error())
		os.Exit(1)
	}
}

func configureViper() {
	viper.SetEnvPrefix(environmentPrefix)
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/nakedsoftware/authorization")
	addHomeDirectoryConfigurationPath()
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			slog.Debug("the configuration file was not found")
		} else {
			slog.Error(
				"failed to read configuration file",
				"error",
				err.Error(),
			)
			os.Exit(1)
		}
	}
}

func addHomeDirectoryConfigurationPath() {
	homePath, err := os.UserHomeDir()
	if err != nil {
		slog.Warn(
			"failed to find the user home directory",
			"error",
			err.Error(),
		)
		return
	}

	configPath := path.Join(
		homePath,
		".nakedsoftware/authorization",
	)
	viper.AddConfigPath(configPath)
}
