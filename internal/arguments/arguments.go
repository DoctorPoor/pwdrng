package arguments

import (
	"errors"
	"os"
	"strconv"
)

const DEFAULT_LENGTH = 16

func GetPasswordLength() (int, error) {
	args := os.Args
	if len(args) < 2 {
		return DEFAULT_LENGTH, nil
	}

	argFirst := args[1]

	if argFirst == "-h" {
		return -1, errors.New(`Name:
  pwdrng - Simple command-line tool to generate passwords

Usage:
  pwdrng [length]

Arguments:
  length    length of the generated password (optional, default 16)

Example usage:
  pwdrng
  pwdrng 32

Further help:
  https://github.com/doctorpoor/pwdrng`)
	}

	passwordLength, err := strconv.Atoi(argFirst)
	if err != nil {
		return -1, errors.New("Password length must be an integer")

	}

	if passwordLength < 4 || passwordLength > 256 {
		return -1, errors.New("Password length must be between 4 and 256")
	}

	return passwordLength, nil
}
