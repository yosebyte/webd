package main

import (
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/yosebyte/x/log"
	"golang.org/x/net/webdav"
)

var version = "dev"

func main() {
	if len(os.Args) < 2 {
		helpInfo()
		os.Exit(1)
	}
	rawURL := os.Args[1]
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		log.Fatal("Unable to parse raw URL: %v", err)
	}
	handler := &webdav.Handler{
		FileSystem: webdav.Dir(parsedURL.Fragment),
		Prefix:     parsedURL.Path,
		LockSystem: webdav.NewMemLS(),
	}
	log.Info("Server started: %v", parsedURL.String())
	for {
		if err := http.ListenAndServe(parsedURL.Host, handler); err != nil {
			log.Error("Server error: %v", err)
			time.Sleep(1 * time.Second)
			log.Info("Server restarted")
		}
	}
}
