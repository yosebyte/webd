## Feature

- A WebDAV server written in Go, using 1-URL command to start.

- Container image also provided at `ghcr.io/yosebyte/webd`.

## Usage

- Basic usage with binary file: `./webd http://hostname:port/path#directory`

- `hostname`: Domain name or IP address, `0.0.0.0` if left blank.

- `port`: Any customizable port, `80` if CDN is applied.

- `path`: Any customizable prefix path, `/` if left blank.

- `directory`: With which directory to share, `./` if left blank.
