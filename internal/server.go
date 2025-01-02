package internal

import (
	"net/http"
	"net/url"

	"github.com/yosebyte/x/log"
	"golang.org/x/net/webdav"
)

func NewServer(parsedURL *url.URL, logger *log.Logger) *http.Server {
	handler := &webdav.Handler{
		Prefix:     parsedURL.Path,
		FileSystem: webdav.Dir(parsedURL.Fragment),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err != nil {
				logger.Warn("Internal: %v", err)
			}
		},
	}
	return &http.Server{
		Addr:     parsedURL.Host,
		ErrorLog: logger.StdLogger(),
		Handler:  handler,
	}
}
