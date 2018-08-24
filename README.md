# Introduction to Go

> The code in this repository is based on Go 1.10.*. It should work on other versions as well.

<!-- TOC -->

- [1. Installation](#1-installation)
  - [1.1. Windows](#11-windows)
  - [1.2. macOS](#12-macos)
  - [1.3. Linux](#13-linux)
- [2. Setup](#2-setup)
  - [2.1. Setting up IDE](#21-setting-up-ide)
- [3. Getting Stated and Conventions](#3-getting-stated-and-conventions)
  - [3.1. Using the Playground](#31-using-the-playground)
  - [3.2. Coding and Naming Conventions](#32-coding-and-naming-conventions)
    - [3.2.1. Naming Conventions](#321-naming-conventions)
      - [3.2.1.1. File Names](#3211-file-names)
      - [3.2.1.2. Function Names](#3212-function-names)
  - [3.3. Go Command Line](#33-go-command-line)
    - [3.3.1. go](#331-go)
    - [3.3.2. godoc](#332-godoc)
    - [3.3.3. gofmt](#333-gofmt)
  - [3.4. Workspace](#34-workspace)
- [4. Let's Start Coding](#4-lets-start-coding)
  - [4.1. Go Essential](#41-go-essential)
    - [4.1.1. Hello World](#411-hello-world)
    - [4.1.2. String Formatting](#412-string-formatting)
    - [4.1.3. Getting Input from Console](#413-getting-input-from-console)
  - [4.2. Managing Values](#42-managing-values)
- [5. What's Next](#5-whats-next)

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

I am using JetBrains GoLand IDE but you can use any IDE you wish, I would recommend using VSCode.

## 3. Getting Stated and Conventions

### 3.1. Using the Playground

Go website has a wonderful way to test your code without installing Go libraries on your system, its called [The Go Playground](https://play.golang.org/).

![The Go Playground](https://github.com/akshaybabloo/Go-Notes/raw/master/assets/images/playground.JPG)

Go also provides a way to run playground (you will need internet though) on your system. Type in:

```bash
> godoc -http=:6060 -play -index
```

Now open your browser and type in `localhost:6060`, this should open playground.

There are few limitations when using the Go playground as its being run on a secured server that has no access to external world:

1. You cannot run HTTP requests or host your web server.
2. The playground fakes the files system so that your IO works correctly.
3. The date and time are always constant - November 10, 2009, at 11:00 p.m.
4. You can work with only one source code at a time; you cannot have multiple go file and import them locally unlike you do on your system.
5. Moreover, it is free.

### 3.2. Coding and Naming Conventions

Before we dive into the coding part of Go language, let's look into some coding and naming conventions of this programming language.

#### 3.2.1. Naming Conventions

##### 3.2.1.1. File Names

According to the Go docs, there is no right or wrong way to name your file - you can either use underscores or camelCase or Pascal Case or small letters, but most Go developers use underscore to name their files. The underscore contradicts the standard library files, for example:

```
go/src/mime/encodedword.go
go/src/mime/encodedword_test.go
```

`encodedword.go` is the name of the main file where other methods are available, and their tests can be found in `encodedword_test.go`. Simple as.

Sometimes you wish to run the file on Windows, Linux or both. In that case, you can name your files something like this

```
filename.go - compiles universally
filename_windows.go - compiles on windows only
filename_unix.go - compiles on posix.
filename_amd64.go - compiles on 64-bit processors only
filename_arm.go - compiles on arm processors only
```

Go build takes care of what to use depending on your operating system, for more information see [https://golang.org/pkg/go/build/](https://golang.org/pkg/go/build/)

##### 3.2.1.2. Function Names

A function always starts with the keyword `func`, the name given to it are of two types; one that begins with a capital letter and the other that begins with a small letter.

For example:

```go
package main

import (
  "strconv"
  "fmt"
  "reflect"
)

func main() {
  sum := Add(1, 2)

  fmt.Println(sum)
  fmt.Println(reflect.TypeOf(sum))
}

func Add(a int, b int) string {
  return toString(a + b)
}

func toString(c int) string {
  return strconv.Itoa(c)
}
```

Let's consider the above example, the `Add()` function takes in two numbers, adds it and converts it to a string by using `toString()` function. The `Add()` function is available throughout the application, as in it can be imported to other go files. Whereas the `toString()` can only be used locally.

### 3.3. Go Command Line

Under `c:\go\bin` there are three commands:

![Go CLI](https://github.com/akshaybabloo/Go-Notes/raw/master/assets/images/go_cli.JPG)

#### 3.3.1. go

This is the primary command you will be using throughout the tutorial. It incorporates more than a dozen commands in it. We will look into most of them in this tutorial. If you type in `go` you will see all the available commands and its description.

#### 3.3.2. godoc

`godoc` is similar to `javadoc` but for Go packages. If you enter `godoc fmt`, it will display all the documentation related to `fmt` package. Typing `godoc` itself will print out all the command flags it can take.

#### 3.3.3. gofmt

This command can be used to format your source code. You can use this by typing `gofmt filename.go` this will only print out the formatted text from the file but does not write it, to make it write the file you should add `-w` flag - `gofmt -w filename.go`.

### 3.4. Workspace

Go follows some strict norms on how you create your workspace structure. The root directory can be any name but should not contain any space or special characters.

The structure should look something like this:

```
workspacename
|
+- src
+- bin
`- pkg
```

`src` is for your source files, `bin` for executables and `pkg` for external libraries.

Once, the folder has been set; the next thing is to create a `GOPATH` to do so type in:

```
set GOTPATH=%USERPROFILE%\gowork
```

> Note: I would create this directory in my `home` folder. It is up to you where to create it.

You can check all the Go environments by typing in `go env`. When you change directory to the `src` folder with a go file in it and type in `go install` it creates an executable file in the `bin` folder.

## 4. Let's Start Coding

> Note: Ignore `_#` (where `#` is the number) as the filename below, the number represents the flow of things to start and end with.

**Code Structure**

```
src
|
`- helloworld_1.go
```

### 4.1. Go Essential

Some helpful Go CLI commands:

1. `godoc` - Used to get the documentation for a given package. Example `godoc fmt`.
2. `gofmt` - To format a Go file to its specifications. Example `gofmt -w filename.go` `-w` will write the file.
3. `go` - It has a lot of sub-commands, in this notes we will be using mostly `go run`
  1. `go run` - Builds and runs a Go file. Example `go run filename.go`
  2. `go build` - Build an executable file. Example `go build filename.go`

#### 4.1.1. Hello World

[Source](https://github.com/akshaybabloo/Go-Notes/blob/master/src/helloworld_1.go)

Say we have a file called `helloworld_1.go`, which has the following content.

```go
package main // The main method should always be in "main" package

import "fmt" // Standard Go library for formatted I/O

func main() { // Go always starts at "main" method
  fmt.Println("hello world") // outputs to I/O with new line.
}
```

To run the file you should use - `go build filename.go` this compiles the file to `filename.exe` (in windows). In this case it will look something like this:

![Go build helloworld_1.go](https://github.com/akshaybabloo/Go-Notes/raw/master/assets/images/go_build.JPG)

Or you can also use `go run` to build the file in `temp` folder and run it. Something like:

![Go build helloworld_1.go](https://github.com/akshaybabloo/Go-Notes/raw/master/assets/images/go_run.JPG)

#### 4.1.2. String Formatting

[source](https://github.com/akshaybabloo/Go-Notes/blob/master/src/stringformats_2.go)

This section looks into [pkg](https://golang.org/pkg/fmt/) package.

A function in Go can return `n` number of objects, which includes `errors`.

For example `fmt.Println` can have multiple returns:

```go
stringLength, err := fmt.Println(str1, str2)
```

If there is no error, the content in the `stringLength` will be returned and the `err` will be equal to `nil`.

#### 4.1.3. Getting Input from Console

[source](https://github.com/akshaybabloo/Go-Notes/blob/master/src/consoleinput_3.go)

For user input you can use `fmt.Scanln(&ref)` and reference to a string variable, but the problem with the `Scanln` is that whenever it encounters a space, the word after the space is given to a new variable.

for example:

```go
var s string
var s2 string

// Send in the reference of the string s
fmt.Scanln(&s, &s2)  // hello hi
// Print the value of it.
fmt.Println(s, s2)  // hello hi
```

To take the input from the user as is, you will have to use different packages, such as `buffio` and `os`.

### 4.2. Managing Values
## 5. What's Next

In the next tutorial (coming soon) we will look into using Go to develop a RESTful web application.
