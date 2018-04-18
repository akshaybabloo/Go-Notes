# Introduction to Go

> The code in this repository is based on Go 1.10.*. It should work on other versions as well.

<!-- TOC -->

- [1. Installation](#1-installation)
  - [1.1. Windows](#11-windows)
  - [1.2. macOS](#12-macos)
  - [1.3. Linux](#13-linux)
- [2. Setup](#2-setup)
  - [2.1. Setting up IDE](#21-setting-up-ide)
- [3. Let's Start Coding](#3-lets-start-coding)
  - [3.1. Using the Playground](#31-using-the-playground)
- [4. What's Next](#4-whats-next)

<!-- /TOC -->

## 1. Installation

Installation is straightforward, just download the executable file and follow the instruction.

### 1.1. Windows

For windows, you will have to download the MSI executable file. Remember that MSI is a 64-bit so make sure you have a 64-bit processor.

Goto [https://golang.org/dl/](https://golang.org/dl/) and download the latest version, for example for Go 1.10.1 it is `go1.10.1.windows-amd64.msi`. Once installed, you check it by opening your command prompt and type in `go version`.

### 1.2. macOS

A pre-built package is available to download and install. Just follow the installation window. Download the latest version at [https://golang.org/dl/](https://golang.org/dl/), for example, for Go 1.10.1 the file name is `go1.10.1.darwin-amd64.pkg`

### 1.3. Linux

For Linux distribution, there is no pre-built package, but it is easy to install. Just download the latest *.tar.gz file from [https://golang.org/dl/](https://golang.org/dl/), for example, for Go 1.10.1 the file is `go1.10.1.linux-amd64.tar.gz`.

Once downloaded, follow the steps:

1. Open your terminal and change the directory to where you download the tar.gz file, then type in:

  ```bash
  tar -c /usr/local -xzf go1.10.1.linux-amd64.tar.gz
  ```

2. Then add the `/usr/local/go/bin` to the path by doing:

```bash
export PATH=$PATH:/usr/local/go/bin
```

You can test it by typing in `go version`.

> Note: You can also add file to your home file by just replacing the `/usr/local/` with `/home/username`, `username` being your user folder.

## 2. Setup

Windows and macOS should automatically update the path to Go, if not you will have to add the path - `installed-folder/go/bin`, where installed-path is where you have installed the Go language. For Linux, see above.

### 2.1. Setting up IDE

I am using JetBrains GoLand IDE but you can use any IDE you wish, I would recommend using VS Code.

## 3. Let's Start Coding

### 3.1. Using the Playground

Go website has a wonderful way to test your code without installing Go libraries on your system, its called [The Go Playground](https://play.golang.org/).

## 4. What's Next

In the next tutorial (coming soon) we will look into using Go to develop a RESTful web service.
