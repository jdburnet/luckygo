# luckygo

## Command Line Interface for Google's "I'm Feeling Lucky" search option.

If you're like me and live in bash shell, luckygo is a great way to ask quick questions from the command line and get quick answers in your web browser. Luckygo uses the google search engine to open the top search results to any query.

## Installation

### From Source

Make sure you have a working Go environment.  Go version 1.2+ is supported.  [See
the install instructions for Go](http://golang.org/doc/install.html).

To install luckygo:

```bash
go get github.com/jdb098/luckygo
go install luckygo
```

Make sure your `PATH` includes the `$GOPATH/bin` directory so your commands can
be easily used:

```bash
export PATH=$PATH:$GOPATH/bin
```

## Use

Example,

```bash
lucky --limit 10 "my search"
```
