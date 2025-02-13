# WEBD

- Simple webdav server using 1-URL command to start.
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
