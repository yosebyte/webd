# 📂 Webd

[![GitHub release](https://img.shields.io/github/v/release/yosebyte/webd)](https://github.com/yosebyte/webd/releases)
[![Go Version](https://img.shields.io/github/go-mod/go-version/yosebyte/webd)](https://golang.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/yosebyte/webd)](https://goreportcard.com/report/github.com/yosebyte/webd)

Webd is a lightweight, efficient WebDAV server designed for simplicity and security. Configured through a single URL command, it supports basic authentication, TLS encryption, and adjustable logging levels. Perfect for file sharing, remote development, and content management scenarios, webd offers a zero-configuration approach with its intuitive URL-style syntax, making it an ideal solution for developers and system administrators who need quick, secure WebDAV capabilities.

## 📋 Table of Contents

- [Features](#-features)
- [Requirements](#-requirements)
- [Installation](#-installation)
  - [Option 1: Pre-built Binaries](#-option-1-pre-built-binaries)
  - [Option 2: Using Go Install](#-option-2-using-go-install)
  - [Option 3: Building from Source](#-option-3-building-from-source)
  - [Option 4: Using Container Image](#-option-4-using-container-image)
- [Usage](#-usage)
- [Configuration](#-configuration)
  - [Log Levels](#-log-levels)
- [Examples](#-examples)
  - [Running as an HTTP WebDAV server](#-running-as-an-http-webdav-server)
  - [Setting up a secure HTTPS WebDAV server](#-setting-up-a-secure-https-webdav-server)
- [How It Works](#-how-it-works)
- [Common Use Cases](#-common-use-cases)
- [Troubleshooting](#-troubleshooting)
- [Contributing](#-contributing)
- [License](#-license)

## ✨ Features

- **🔑 Basic Authentication**: Secure access to your WebDAV server with username/password
- **🔒 TLS Support**: Enable HTTPS for encrypted data transfer
- **📊 Flexible Logging System**: Configurable verbosity with multiple logging levels
- **🌐 HTTP/HTTPS Protocol Support**: Full support for both protocols
- **📦 Single-Binary Deployment**: Easy to distribute and install with no dependencies
- **⚙️ Zero Configuration Files**: Everything is specified via a single URL command
- **🚀 Low Resource Footprint**: Minimal CPU and memory usage

## 📋 Requirements

- Go 1.24 or higher (for building from source)
- Network access for serving WebDAV content
- Admin privileges may be required for binding to ports below 1024

## 📥 Installation

### 💾 Option 1: Pre-built Binaries

Download the latest release for your platform from our [releases page](https://github.com/yosebyte/webd/releases).

### 🔧 Option 2: Using Go Install

```bash
go install github.com/yosebyte/webd/cmd/webd@latest
```

### 🛠️ Option 3: Building from Source

```bash
# Clone the repository
git clone https://github.com/yosebyte/webd.git

# Build the binary
cd webd
go build -o webd ./cmd/webd

# Optional: Install to your GOPATH/bin
go install ./cmd/webd
```

### 🐳 Option 4: Using Container Image

Webd is available as a container image:

```bash
# Pull the container image
docker pull ghcr.io/yosebyte/webd:latest

# Run a WebDAV server
docker run -d --rm -p 10101:10101 -v /path/to/share:/data ghcr.io/yosebyte/webd http://qwer:1234@0.0.0.0:10101/secret#/data
```

## 🚀 Usage

Webd operates using a single, intuitive URL-style command:

```bash
webd http://<username>:<password>@<server_addr>/<prefix>?<log=level>#<root_dir>
```

- `username:password`: Basic authentication credentials (optional)
- `server_addr`: Server address to bind to (hostname:port)
- `prefix`: URL prefix for the WebDAV endpoint
- `log`: Log level (debug, info, warn, error, fatal)
- `root_dir`: Local directory to share via WebDAV

## ⚙️ Configuration

Webd uses a minimalist approach with all configuration specified in the URL command:

### 📝 Log Levels

- `debug`: Verbose debugging information showing all operations
- `info`: General operational information (default)
- `warn`: Warning conditions
- `error`: Error conditions
- `fatal`: Critical conditions

## 📚 Examples

### 🌍 Running as an HTTP WebDAV server

```bash
webd http://qwer:1234@10.1.0.1:10101/secret?log=debug#/root
```

This will:
1. Start a WebDAV server on 10.1.0.1:10101
2. Require username "qwer" and password "1234" for access
3. Mount the local "/root" directory to the "/secret" WebDAV path
4. Log at the DEBUG level

### 🔐 Setting up a secure HTTPS WebDAV server

```bash
webd https://qwer:1234@10.1.0.1:10101/secret?log=warn#/root
```

This will:
1. Start a WebDAV server with TLS encryption on 10.1.0.1:10101
2. Generate a self-signed certificate for secure communication
3. Require username "qwer" and password "1234" for access
4. Mount the local "/root" directory to the "/secret" WebDAV path
5. Log at the WARN level

## 🔍 How It Works

Webd creates a WebDAV server using the Go standard library and golang.org/x/net/webdav:

1. Parses the provided URL to extract configuration parameters
2. Sets up authentication if username/password are provided
3. Configures TLS for HTTPS connections when requested
4. Mounts the specified directory as a WebDAV resource
5. Handles incoming connections with proper logging

## 💡 Common Use Cases

- **📁 Remote File Access**: Access your files from anywhere securely
- **💻 Development Environments**: Mount remote directories for development
- **🔄 Content Management**: Simple way to manage and update website content
- **📱 Mobile Device Access**: Connect to your files from mobile devices
- **🔒 Secure File Sharing**: Share files with teams using authentication

## 🔧 Troubleshooting

- Ensure the target directory exists and has appropriate permissions
- Verify firewall settings allow traffic on the specified port
- For HTTPS, check that clients properly handle self-signed certificates
- Increase log level to debug for more detailed information

## 👥 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Open a Pull Request

## 📄 License

This project is licensed under the [MIT License](LICENSE).
