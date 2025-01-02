package main

import (
	"context"
	"net/http"
	"net/url"
	"os"

	"github.com/yosebyte/webd/internal"
)

func coreDispatch(parsedURL *url.URL, stop chan os.Signal) {
	switch parsedURL.Scheme {
	case "http":
		runServer(parsedURL, stop)
	default:
		logger.Fatal("Invalid scheme: %v", parsedURL.Scheme)
		getExitInfo()
	}
}

func runServer(parsedURL *url.URL, stop chan os.Signal) {
	server := internal.NewServer(parsedURL, logger)
	go func() {
		logger.Info("Server started: %v", parsedURL.Host)
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
