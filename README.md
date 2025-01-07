![WEBD](https://img.shields.io/badge/Yosebyte-WEBD-blue)
![GitHub License](https://img.shields.io/github/license/yosebyte/webd)
[![Go Report Card](https://goreportcard.com/badge/github.com/yosebyte/webd)](https://goreportcard.com/report/github.com/yosebyte/webd)

# WEBD

- Simple webdav server from the Yosebyte Collections.
- Written in Golang, using 1-URL command to start.
- Support basic-auth, tls and different log levels.
- Container image provided at [ghcr.io/yosebyte/webd](https://ghcr.io/yosebyte/webd).

## Usage

```
webd http://<username>:<password>@<server_addr>/<prefix>?<log=level>#<root_dir>

# Run as http webdav server
webd http://qwer:1234@10.1.0.1:10101/secret?log=debug#/root

# Run as https webdav server
webd https://qwer:1234@10.1.0.1:10101/secret?log=warn#/root
```
