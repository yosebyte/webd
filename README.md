## Feature

- A WebDAV server written in Go, doing more works with a smaller size.

- Container image also provided at `ghcr.io/yosebyte/webd`, built from scratch.

## Usage

- Basic usage with binary file: `./webd http://hostname:port/path#directory`

- `hostname`: Domain name or IP address, `0.0.0.0` if left blank.

- `port`: Any customizable port， except `80` if you use CDN.

- `path`: Any customizable prefix path, `/` if left blank.

- `directory`: With which directory to share, `.` if leave blank.
