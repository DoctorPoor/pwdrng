# pwdrng

[![Go Reference](https://pkg.go.dev/badge/github.com/doctorpoor/pwdrng.svg)](https://pkg.go.dev/github.com/doctorpoor/pwdrng)
[![Go version](https://img.shields.io/github/go-mod/go-version/doctorpoor/pwdrng)](https://github.com/DoctorPoor/pwdrng/blob/master/go.mod#L3)
[![Go Report Card](https://goreportcard.com/badge/github.com/doctorpoor/pwdrng)](https://goreportcard.com/report/github.com/doctorpoor/pwdrng)
[![GitHub repo size](https://img.shields.io/github/repo-size/doctorpoor/pwdrng)](https://github.com/DoctorPoor/pwdrng/releases)
[![GitHub repo size](https://img.shields.io/github/license/doctorpoor/pwdrng)](https://github.com/doctorpoor/pwdrng/blob/master/LICENSE)

Simple command-line tool to generate passwords

```
$ pwdrng
```

```
Copied password to clipboard:
bfx861[X<26-b^UT
```

## Installation

- With [Homebrew](https://docs.brew.sh/Installation)

```
$ brew tap doctorpoor/pwdrng
$ brew install pwdrng
```

- With [Go](https://go.dev/doc/install)

```
$ go install github.com/doctorpoor/pwdrng@latest
```

## Usage

```
$ pwdrng -h
```

```
Example usage:
  pwdrng
  pwdrng -l 32

Flags:
  -l, -length, --length int
    length of the generated password (default 16)

Further help:
  https://github.com/doctorpoor/pwdrng
```

## License

[MIT](https://github.com/DoctorPoor/pwdrng/blob/master/LICENSE)
