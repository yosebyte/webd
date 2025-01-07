package internal

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/yosebyte/x/log"
	"golang.org/x/net/webdav"
)

func NewServer(parsedURL *url.URL, tlsConfig *tls.Config, logger *log.Logger) *http.Server {
	handler := &webdav.Handler{
		Prefix:     parsedURL.Path,
		FileSystem: webdav.Dir(parsedURL.Fragment),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			logger.Debug("%v: %v -> %v", r.Method, r.RemoteAddr, r.URL)
			if err != nil {
				logger.Debug("%v: %v", r.Method, err)
			}
		},
	}
	return &http.Server{
		Addr:      parsedURL.Host,
		ErrorLog:  logger.StdLogger(),
		Handler:   updateHandler(handler, parsedURL, logger),
		TLSConfig: tlsConfig,
	}
}

func updateHandler(handler http.Handler, parsedURL *url.URL, logger *log.Logger) http.Handler {
	username := parsedURL.User.Username()
	password, hasPassword := parsedURL.User.Password()
	if username != "" && hasPassword {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if u, p, ok := r.BasicAuth(); !ok || u != username || p != password {
				w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				logger.Warn("Unauthorized access: %v", r.RemoteAddr)
				return
			}
			handler.ServeHTTP(w, r)
		})
	}
	return handler
}
