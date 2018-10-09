# Introduction to Go

> The code in this repository is based on Go 1.11.*. It should work on other versions as well.

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
    - [4.2.1. Exploring variables, constants and types](#421-exploring-variables-constants-and-types)
      - [4.2.1.1. Implicit vs Explicit](#4211-implicit-vs-explicit)
      - [4.2.1.2. Constants](#4212-constants)
      - [4.2.1.3. Data Types](#4213-data-types)
    - [4.2.2. Working with Strings](#422-working-with-strings)
    - [4.2.3. Using Math Operator and Math Package](#423-using-math-operator-and-math-package)
    - [4.2.4. Date and Time](#424-date-and-time)
  - [4.3. Managing Complex Types and collections](#43-managing-complex-types-and-collections)
    - [4.3.1. Pointers](#431-pointers)
    - [4.3.2. Arrays](#432-arrays)
    - [4.3.3. Slices](#433-slices)
    - [4.3.4. Memory](#434-memory)
      - [4.3.4.1. The `new()` Function](#4341-the-new-function)
      - [4.3.4.2. The `make()` Function](#4342-the-make-function)
      - [4.3.4.3. Creating `nil` Objects](#4343-creating-nil-objects)
      - [4.3.4.4. Memory Deallocation](#4344-memory-deallocation)
    - [4.3.5. Maps](#435-maps)
    - [4.3.6. Data Structure](#436-data-structure)
  - [4.4. Managing Program Flow](#44-managing-program-flow)
    - [4.4.1. Conditional Logic](#441-conditional-logic)
    - [4.4.2. Switch Statement](#442-switch-statement)
    - [4.4.3. Loops](#443-loops)
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

[Source - helloworld_1.go](https://github.com/akshaybabloo/Go-Notes/blob/master/src/helloworld_1.go)

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

[Source - stringformats_2.go](https://github.com/akshaybabloo/Go-Notes/blob/master/src/stringformats_2.go)

This section looks into [pkg](https://golang.org/pkg/fmt/) package.

A function in Go can return `n` number of objects, which includes `errors`.

For example `fmt.Println` can have multiple returns:

```go
stringLength, err := fmt.Println(str1, str2)
```

If there is no error, the content in the `stringLength` will be returned and the `err` will be equal to `nil`.

#### 4.1.3. Getting Input from Console

[Source - consoleinput_3.go](https://github.com/akshaybabloo/Go-Notes/blob/master/src/consoleinput_3.go)

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

To take the input from the user as is, you will have to use different packages, such as `bufio` and `os`.

### 4.2. Managing Values

In this section we will look into formatting string, date time, date types.

#### 4.2.1. Exploring variables, constants and types

##### 4.2.1.1. Implicit vs Explicit

Using `var` keyword and an assignment operator with a data type you can declare a value.

```go
var anInteger int = 30
var aString string = "Hello!"
```

Go allows you to implicitly declare a variable without using a data type and the `var` keyword and replace them by using `:=`.

```go
anInteger := 30
aString := "Hello!"
```

Above code is still static and cannot be changed.

##### 4.2.1.2. Constants

Explicit:

```go
const anInteger int = 30
const aString string = "Hello!"
```

Implicit:

```go
const anInteger = 30
const aString = "Hello!"
```

##### 4.2.1.3. Data Types

All data types are spelled in lower case.

1. `bool` - Booleans values
2. `string` - String type
3. Integers:
  1. Fixed Integers:
  ```
  uint8 uint16 uint32 uint64
  int8 int16 int32 int64
  ```
  2. Aliases: They could be a 32 bit or a 64 bit depending upon the operating system.
  ```
  byte uint int uintptr
  ```
4. Floating values:
  ```
  float32 float64
  ```
5. Complex numbers:
  ```
  complex64 complex128
  ```
6. Data collections:
  ```
  Arrays Slices Maps Structs
  ```
7. Language organization:
  ```
  Functions Interfaces Channels
  ```
8. Data Management:
  ```
  Pointers
  ```

#### 4.2.2. Working with Strings

[Source - workingwithstrings_4.go](https://github.com/akshaybabloo/Go-Notes/blob/master/src/workingwithstrings_4.go)

Go language is case sensitive as in `hello != Hello`. For more information on `strings` see [http://golang.com/pkg/strings/](https://golang.com/pkg/strings/)

#### 4.2.3. Using Math Operator and Math Package

[Source - math_5.go](https://github.com/akshaybabloo/Go-Notes/blob/master/src/math_5.go)

Math Operation

| Operator | Definition |
|:--------:|:----------:|
|     +    |     Add    |
|     -    |  Subtract  |
|     *    |  Multiply  |
|     /    |   Divide   |
|     %    |  Remainder |

Bitwise Operation

| Operator |  Definition |
|:--------:|:-----------:|
|     &    | Bitwise AND |
|     \|   |  Bitwise OR |
|     ^    | Bitwise XOR |
|    &^    |   Bitclear  |
|    >>    | Right Shift |
|    <<    |  Left Shift |

In Go, numeric types don't implicitly convert, i.e you cannot add a float to int, the only way you can add ghem is by converting an int to float. For example:

```go
int1 := 30
float1 :=30.1
fmt.Println(float64(int1) + float1)
```

For more complex math problems see [http://golang.com/pkg/math/](https://golang.com/pkg/math/).

#### 4.2.4. Date and Time

[Source - datetime_6.go](https://github.com/akshaybabloo/Go-Notes/blob/master/src/datetime_6.go)

When trying to format date and time make sure you format based on the constant numbers given in the package.

```
1 - Month
2 - Day
3 - Hour
4 - Minutes
5 - Seconds
6 - Year
7 - Timezone
```

For example:

```go
shortFormat := "2/1/06"
fmt.Println("Today is", time.Now().Format(shortFormat))
```

### 4.3. Managing Complex Types and collections

In this section we will look into some complex types such as pointers and collections such as arrays, maps and Structs.

#### 4.3.1. Pointers

[Source - pointers_7.go](https://github.com/akshaybabloo/Go-Notes/blob/master/src/pointers_7.go)

Like C/C++, Go also support pointers natively.

#### 4.3.2. Arrays

[Source - arrays_8.go](https://github.com/akshaybabloo/Go-Notes/blob/master/src/arrays_8.go)

All the contents in an array are of same type. Like Java, the first index always starts with `0`.

There are some memory implication when using arrays in Go, unlike C/C++ (where arrays are pointers) arrays in Go are values i.e. when assigning one array to another copies all the contents of it. Passing an array to the function will pass a copy of the array and not the pointer.

#### 4.3.3. Slices

[Source - slices_9.go](https://github.com/akshaybabloo/Go-Notes/blob/master/src/slices_9.go)

A slice is an abstraction layer that sits on top of an Array. Like arrays all the contents in the slices are of the same type, but unlike it, slices can are resizeable and can be sorted easily.

#### 4.3.4. Memory

Go's runtime is statistically linked into each compiled application. Go's memory is managed my the runtime on a different thread, that means memory allocation and deallocation is completely done automatically in the background.

To understand how the memory allocation works there two function that you'll need to understand - `make()` and `new()`, these function can be used to initialise `maps`, `slices` and `channels`.

In most cases you will like be using `make()` instead of `new()`.

##### 4.3.4.1. The `new()` Function

1. Allocates but does not initialise memory
2. Results in zeroed storage but returns a memory address

##### 4.3.4.2. The `make()` Function

1. Allocates and initialise memory
2. Results in non-zeroed storage but returns a memory address

##### 4.3.4.3. Creating `nil` Objects

You must always initialise complex objects before adding values, declaring them without `make()` will cause a _panic_.

For example:

```go
var m map[string]int // key value pair of string and integer.
m["key"] = 10
fmt.Println(m)
```

With the above code you will get a _panic_ error at the second line as you cannot add a value to zeroed memory. To fix this you can do;

```go
m := make(map[string]int) // key value pair of string and integer.
m["key"] = 10
fmt.Println(m)
// map[key:42]
```

##### 4.3.4.4. Memory Deallocation

Memory is deallocated by using garbage collector (GC), only the objects that are out of scope or set to `nil` are eligible.

For more information on GC you can go to [https://golang.org/pkg/runtime/](https://golang.org/pkg/runtime/).

#### 4.3.5. Maps

[Source - maps_10.go](https://github.com/akshaybabloo/Go-Notes/blob/master/src/maps_10.go)

Maps are unordered collections of key-value pairs aka hash table.

#### 4.3.6. Data Structure

[Source - structs_11.go](https://github.com/akshaybabloo/Go-Notes/blob/master/src/structs_11.go)

Structs in Go have similar function to C's struct and Java's classes. They both contain data and methods, but unlike Java, Go doesn't have an inheritance model.

In this section we will look into how to use struct to store data.

### 4.4. Managing Program Flow

#### 4.4.1. Conditional Logic

[Source - conditions_12.go](https://github.com/akshaybabloo/Go-Notes/blob/master/src/conditions_12.go)

In this section we will look at how the conditions in Go differs from C.

You will have to remember that the `else` keyword should come after the braces. Example:

```go

// This is wrong. you will get an error
if ...{
  ...
}
else {
  ...
}

// Instead use this
if ... {
  ...
} else {
  ...
}
```

This is because the way the Go lexer see the line feeds.

#### 4.4.2. Switch Statement

[Source - switch_13.go](https://github.com/akshaybabloo/Go-Notes/blob/master/src/switch_13.go)

Go's switch statement improvises C's and Java's. Like the recent versions of Java, Go can evaluate simple types, not just constants and strings. The code flow jumps from one case to another once the case if found, this means you don't have to add the redundant `break` statement.

#### 4.4.3. Loops

[Source - loops_14.go](https://github.com/akshaybabloo/Go-Notes/blob/master/src/loops_14.go)

Unlike Java there is no `while` statement in Go, but there are extended version of the `for` loop, which will be seen in this section.

## 5. What's Next

In the next tutorial (coming soon) we will look into using Go to develop a RESTful web application.
