package main

import (
	"runtime"

	"github.com/yosebyte/x/log"
)

func helpInfo() {
	log.Info(`Version: %v %v/%v

Usage:
	webd http://<addr>/<pre>#<dir>

Examples:
	webd http://10.0.0.1:10101/secret/#/root

Arguments:
	<addr>  Server ip:port to be exposed
	<pre>   Optional prefix, default "/"
	<dir>   Root directory, default "./"
`, version, runtime.GOOS, runtime.GOARCH)
}
