![WEBD](https://img.shields.io/badge/Yosebyte-WEBD-blue)
![GitHub License](https://img.shields.io/github/license/yosebyte/webd)
[![Go Report Card](https://goreportcard.com/badge/github.com/yosebyte/webd)](https://goreportcard.com/report/github.com/yosebyte/webd)

# WEBD

- Simple webdav server from the Yosebyte Collections.
- Written in Golang, using 1-URL command to start.
- Container image provided at [ghcr.io/yosebyte/webd](https://ghcr.io/yosebyte/webd).

## Usage

```
Usage:
    webd http://<addr>/<pre>#<dir>

Examples:
    webd http://10.1.0.1:10101/secret/#/root

Arguments:
    <addr>  Server ip:port to be exposed
    <pre>   Optional prefix, default "/"
    <dir>   Root directory, default "./"
```
