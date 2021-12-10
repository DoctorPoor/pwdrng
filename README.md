# pwdrng

Simple command-line tool to generate passwords

```
$ pwdrng
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
