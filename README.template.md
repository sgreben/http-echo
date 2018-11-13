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
curl -L https://github.com/sgreben/${APP}/releases/download/${VERSION}/${APP}_${VERSION}_linux_x86_64.tar.gz | tar xz

# OS X
curl -L https://github.com/sgreben/${APP}/releases/download/${VERSION}/${APP}_${VERSION}_osx_x86_64.tar.gz | tar xz

# Windows
curl -LO https://github.com/sgreben/${APP}/releases/download/${VERSION}/${APP}_${VERSION}_windows_x86_64.zip
unzip versions_${VERSION}_windows_x86_64.zip
```

## Use it

```text
${USAGE}
```
