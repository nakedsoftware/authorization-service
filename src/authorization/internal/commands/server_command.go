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
	"github.com/nakedsoftware/authorization-service/src/authorization/internal/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const defaultPort int16 = 80
const portFlag = "port"
const portKey = "server.port"

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Starts Authorization Service to listen for requests",
	Long: `
The server command starts Authorization Service to listen for incoming
authorization and token requests from client applications.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return server.Serve(cmd.Context(), viper.GetInt(portKey))
	},
}

func init() {
	rootCommand.AddCommand(serverCommand)

	_ = viper.BindEnv(portKey, "APP_PORT")
	viper.SetDefault(portKey, defaultPort)
	serverCommand.PersistentFlags().Int16P(
		portFlag,
		"p",
		defaultPort,
		"The TCP/IP port to listen for requests on",
	)
	_ = viper.BindPFlag(
		portKey,
		serverCommand.PersistentFlags().Lookup(portFlag),
	)
}
