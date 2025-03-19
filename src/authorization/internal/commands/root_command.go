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

package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log/slog"
	"os"
)

const (
	jsonOutputKey = "log.jsonOutput"
	logLevelKey   = "log.level"
)

const (
	jsonOutputFlag = "json-output"
	logLevelFlag   = "log-level"
)

const (
	levelDebug = "debug"
	levelError = "error"
	levelInfo  = "info"
	levelWarn  = "warn"
)

var rootCommand = &cobra.Command{
	Use:     "authorization",
	Version: "1.0.0",
	Short:   "Authorization Service",
	Long: `
Authorization Service provides identity, authentication, and authorization
services for web applications, services, and client applications.
Authorization Service supports the OAuth 2, OpenID Connect, and WebAuthn
protocols to support authenticating users, providing user identity information
to applications and services, and authorizing access to resources in resource
services.
`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var level slog.Level
		logLevel := viper.GetString(logLevelKey)
		switch logLevel {
		case levelDebug:
			level = slog.LevelDebug
		case levelInfo:
			level = slog.LevelInfo
		case levelWarn:
			level = slog.LevelWarn
		case levelError:
			level = slog.LevelError
		default:
			return fmt.Errorf("unknown log level: %s", logLevel)
		}

		handlerOptions := &slog.HandlerOptions{
			Level: level,
		}

		var handler slog.Handler
		if viper.GetBool(jsonOutputKey) {
			handler = slog.NewJSONHandler(os.Stderr, handlerOptions)
		} else {
			handler = slog.NewTextHandler(os.Stderr, handlerOptions)
		}

		logger := slog.New(handler)
		slog.SetDefault(logger)

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	viper.SetDefault(jsonOutputKey, false)
	rootCommand.PersistentFlags().Bool(
		jsonOutputFlag,
		false,
		"output log events in JSON format",
	)
	_ = viper.BindPFlag(
		jsonOutputKey,
		rootCommand.PersistentFlags().Lookup(jsonOutputFlag),
	)

	viper.SetDefault(logLevelKey, levelInfo)
	rootCommand.PersistentFlags().String(
		logLevelFlag,
		levelInfo,
		"The level of logging output to see (debug, info, warn, wrror",
	)
	_ = viper.BindPFlag(
		logLevelKey,
		rootCommand.PersistentFlags().Lookup(logLevelFlag),
	)
}
