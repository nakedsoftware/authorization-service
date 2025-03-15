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

package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/nakedsoftware/authorization-service/src/authorization/internal/handlers/authorization"
	"github.com/nakedsoftware/authorization-service/src/authorization/internal/handlers/token"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const timeout = 5 * time.Second

func Serve(ctx context.Context, port int) error {
	mux := http.NewServeMux()
	mux.Handle("GET /authorize", authorization.NewHandler())
	mux.Handle("POST /token", token.NewHandler())

	serverCtx, stop := signal.NotifyContext(
		ctx,
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	listenAddress := fmt.Sprintf(":%d", port)
	server := http.Server{
		Addr:    listenAddress,
		Handler: mux,
		BaseContext: func(_ net.Listener) context.Context {
			return serverCtx
		},
	}

	_ = context.AfterFunc(serverCtx, func() {
		slog.Info("stopping authorization service")
		timeoutCtx, cancel := context.WithTimeout(
			context.Background(),
			timeout,
		)
		defer cancel()

		if err := server.Shutdown(timeoutCtx); err != nil {
			slog.Error("Failed to shut down server", "error", err.Error())
			os.Exit(0)
		}

		select {
		case <-timeoutCtx.Done():
			err := timeoutCtx.Err()
			if errors.Is(err, context.DeadlineExceeded) {
				slog.Error("timeout exceeded; forcing shutdown")
			} else if err != nil {
				slog.Error("failed to shutdown server", "error", err.Error())
			}

			os.Exit(0)
		}
	})

	slog.Info("listening for requests", "address", listenAddress)
	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	slog.Error("failed to start authorization service", "error", err.Error())
	return err
}
