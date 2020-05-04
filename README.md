# http-echo

`http-echo` is a dependency-free HTTP server that always responds with the request it receives.

## Contents

<!-- TOC -->

- [Contents](#contents)
- [Examples](#examples)
- [Get it](#get-it)
    - [Using `go get`](#using-go-get)
    - [Pre-built binary](#pre-built-binary)
- [Use it](#use-it)

<!-- /TOC -->


## Examples

```sh
$ http-echo
2018/11/13 20:58:13.542893  http-echo listening on ":8080", responding with HTTP 200 (OK)
```

```sh
$ http-echo -p 1234
2018/11/13 20:57:45.561703 http-echo listening on ":1234", responding with HTTP 200 (OK)
```

```sh
$ export PORT=9999
$ http-echo
2018/11/13 21:00:45.927216 http-echo listening on ":9999", responding with HTTP 200 (OK)
```

## Get it

### Using `go get`

```sh
go get -u github.com/sgreben/http-echo
```

### Pre-built binary

Or [download a binary](https://github.com/sgreben/http-echo/releases/latest) from the releases page, or from the shell:

```sh
# Linux
curl -L https://github.com/sgreben/http-echo/releases/download/1.0.1/http-echo_1.0.1_linux_x86_64.tar.gz | tar xz

# OS X
curl -L https://github.com/sgreben/http-echo/releases/download/1.0.1/http-echo_1.0.1_osx_x86_64.tar.gz | tar xz

# Windows
curl -LO https://github.com/sgreben/http-echo/releases/download/1.0.1/http-echo_1.0.1_windows_x86_64.zip
unzip versions_1.0.1_windows_x86_64.zip
```

## Use it

```text
Usage of http-echo:
  -a string
    	(alias for -addr) (default ":8080")
  -addr string
    	address to listen on (environment variable "ADDR") (default ":8080")
  -http-status int
    	use this HTTP status for responses (environment variable "HTTP_STATUS") (default 200)
  -p int
    	(alias for -port)
  -port int
    	port to listen on (overrides -addr port) (environment variable "PORT")
  -q	(alias for -quiet)
  -quiet
    	disable all log output (environment variable "QUIET")
```
