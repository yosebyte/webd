package main

import (
	"net/url"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/yosebyte/x/log"
)

var (
	logger  = log.NewLogger(log.Info, true)
	version = "dev"
)

func main() {
	parsedURL := getParsedURL(os.Args)
	initLogLevel(parsedURL.Query().Get("log"))
	coreDispatch(parsedURL, getStopSignal())
}

func getStopSignal() chan os.Signal {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	return stop
}

func getParsedURL(args []string) *url.URL {
	if len(args) < 2 {
		getExitInfo()
	}
	parsedURL, err := url.Parse(args[1])
	if err != nil {
		logger.Fatal("URL parse: %v", err)
		getExitInfo()
	}
	return parsedURL
}

func initLogLevel(level string) {
	switch level {
	case "debug":
		logger.SetLogLevel(log.Debug)
		logger.Debug("Init log level: DEBUG")
	case "warn":
		logger.SetLogLevel(log.Warn)
		logger.Warn("Init log level: WARN")
	case "error":
		logger.SetLogLevel(log.Error)
		logger.Error("Init log level: ERROR")
	case "fatal":
		logger.SetLogLevel(log.Fatal)
		logger.Fatal("Init log level: FATAL")
	default:
		logger.SetLogLevel(log.Info)
		logger.Info("Init log level: INFO")
	}
}

func getExitInfo() {
	logger.SetLogLevel(log.Info)
	logger.Info(`Version: %v %v/%v

Usage:
	webd http://<addr>/<pre>#<dir>

Examples:
	webd http://10.1.0.1:10101/secret/#/root

Arguments:
	<addr>  Server ip:port to be exposed
	<pre>   Optional prefix, default "/"
	<dir>   Root directory, default "./"
`, version, runtime.GOOS, runtime.GOARCH)
	os.Exit(1)
}
