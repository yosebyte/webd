package main

import (
	"context"
	"net/http"
	"net/url"
	"os"

	"github.com/yosebyte/webd/internal"
	"github.com/yosebyte/x/tls"
)

func coreDispatch(parsedURL *url.URL, stop chan os.Signal) {
	switch parsedURL.Scheme {
	case "http":
		serveHTTP(parsedURL, stop)
	case "https":
		serveHTTPS(parsedURL, stop)
	default:
		logger.Fatal("Invalid scheme: %v", parsedURL.Scheme)
		getExitInfo()
	}
}

func serveHTTP(parsedURL *url.URL, stop chan os.Signal) {
	server := internal.NewServer(parsedURL, nil, logger)
	go func() {
		logger.Info("Server started: %v", parsedURL.String())
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Server error: %v", err)
		}
	}()
	<-stop
	logger.Info("Server stopping")
	if err := server.Shutdown(context.Background()); err != nil {
		logger.Error("Server shutdown error: %v", err)
	}
	logger.Info("Server stopped")
}

func serveHTTPS(parsedURL *url.URL, stop chan os.Signal) {
	tlsConfig, err := tls.NewTLSconfig("yosebyte/webd:" + version)
	if err != nil {
		logger.Error("Unable to generate TLS config: %v", err)
	}
	server := internal.NewServer(parsedURL, tlsConfig, logger)
	go func() {
		logger.Info("Server started: %v", parsedURL.String())
		if err := server.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
			logger.Error("Server error: %v", err)
		}
	}()
	<-stop
	logger.Info("Server stopping")
	if err := server.Shutdown(context.Background()); err != nil {
		logger.Error("Server shutdown error: %v", err)
	}
	logger.Info("Server stopped")
}
