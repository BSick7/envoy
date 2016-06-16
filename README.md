# envoy

Inspired by [https://github.com/kailunshi/consul-backup](https://github.com/kailunshi/consul-backup).

This utility helps manage consul.

## Install

### Using `go get`

If you have golang installed, you can use issue the following:

```bash
$ go get github.com/BSick7/envoy
```

### Download

Alternatively, download the binary for your OS/Architecture from [releases](https://github.com/BSick7/envoy/releases).

### Build from source

```bash
$ mkdir -p $GOPATH/src/github.com/BSick7
$ cd $GOPATH/src/github.com/BSick7
$ git clone git@github.com:BSick7/envoy.git
$ cd envoy
$ make deps
$ make install
```

## Usage

```bash
$ go run main.go
usage: envoy [--version] [--help] <command> [<args>]

Available commands are:
    backup     Backup consul k/v store
    restore    Restores consul k/v store
```

### Backup

```bash
$ envoy backup -http-address=127.0.0.1:8500 output.tar.gz
```

```bash
$ CONSUL_HTTP_ADDR=127.0.0.1:8500
$ envoy backup > output.tar.gz
```

### Restore

```bash
$ envoy restore -http-address=127.0.0.1:8500 output.tar.gz
```

```bash
$ CONSUL_HTTP_ADDR=127.0.0.1:8500
$ cat output.tar.gz | envoy restore
```

## Deploying to GitHub Releases

`GITHUB_TOKEN` environment variable must exist in CircleCI with `public_repo` and `repo` scope.
Travis automates this, but we are using CircleCI.
