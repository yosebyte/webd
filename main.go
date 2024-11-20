package main

import (
    "log"
    "net/url"
    "net/http"
    "os"

    "golang.org/x/net/webdav"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatalf("[ERRO] Usage: http://hostname:port/path#directory")
    }
    rawURL := os.Args[1]
    parsedURL, err := url.Parse(rawURL)
    if err != nil {
        log.Printf("[WARN] %v", err)
    }
    webdavHandler := &webdav.Handler{
        FileSystem: webdav.Dir(parsedURL.Fragment),
        Prefix:     parsedURL.Path,
        LockSystem: webdav.NewMemLS(),
    }
    log.Printf("[INFO] %v", rawURL)
    if err := http.ListenAndServe(parsedURL.Host, webdavHandler); err != nil {
        log.Fatalf("[ERRO] %v", err)
    }
}
