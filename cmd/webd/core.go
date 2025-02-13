package main

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/yosebyte/webd/internal"
	"github.com/yosebyte/x/tls"
)

func coreDispatch(parsedURL *url.URL, signalChan chan os.Signal) {
	switch parsedURL.Scheme {
	case "http":
		serveHTTP(parsedURL, signalChan)
	case "https":
		serveHTTPS(parsedURL, signalChan)
	default:
		logger.Fatal("Invalid scheme: %v", parsedURL.Scheme)
		getExitInfo()
	}
}

func serveHTTP(parsedURL *url.URL, signalChan chan os.Signal) {
	server := internal.NewServer(parsedURL, nil, logger)
	go func() {
		logger.Info("Server started: %v", parsedURL.String())
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Server error: %v", err)
		}
	}()
	<-signalChan
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	logger.Info("Server shutting down")
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown error: %v", err)
	}
	logger.Info("Server shutdown complete")
}

func serveHTTPS(parsedURL *url.URL, signalChan chan os.Signal) {
	logger.Info("Apply RAM cert: %v", version)
	tlsConfig, err := tls.GenerateTLSConfig("yosebyte/webd:" + version)
	if err != nil {
		logger.Fatal("Generate failed: %v", err)
		return
	}
	server := internal.NewServer(parsedURL, tlsConfig, logger)
	go func() {
		logger.Info("Server started: %v", parsedURL.String())
		if err := server.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
			logger.Error("Server error: %v", err)
		}
	}()
	<-signalChan
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	logger.Info("Server shutting down")
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown error: %v", err)
	}
	logger.Info("Server shutdown complete")
}
