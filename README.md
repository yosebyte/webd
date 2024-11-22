## Feature

- A WebDAV server written in Go, doing more works with a smaller size.

- Container image also provided at `ghcr.io/yosebyte/webd`, container size less than 6MB.

## Usage

- Basic usage with binary file: `./webd http://hostname:port/path#directory`

- `hostname`: Domain name or IP address, `0.0.0.0` if left blank.

- `port`: Any customizable port, `80` if CDN applied.

- `path`: Any customizable prefix path, `/` if left blank.

- `directory`: With which directory to share, `./` if leave blank.
