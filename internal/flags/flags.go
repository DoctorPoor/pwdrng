package flags

import (
	"errors"
	"flag"
	"fmt"
)

type FlagValues struct {
	Length int
}

func setFlagUsage() {
	usageString := `Example usage:
  pwdrng
  pwdrng -l 32

Flags:
  -l, -length, --length int
    length of the generated password (default 16)

Further help:
  https://github.com/doctorpoor/pwdrng`

	flag.Usage = func() {
		fmt.Println(usageString)
	}
}

func getPasswordLengthValue() (int, error) {
	defaultLength := 16
	usage := "length of the generated password"

	var passwordLength int
	flag.IntVar(&passwordLength, "length", defaultLength, usage)
	flag.IntVar(&passwordLength, "l", defaultLength, usage)
	flag.Parse()

	if passwordLength < 4 || passwordLength > 256 {
		return -1, errors.New("password length must be between 4 and 256")
	}

	return passwordLength, nil
}

func GetFlagValues() (FlagValues, error) {
	setFlagUsage()

	passwordLength, err := getPasswordLengthValue()
	if err != nil {
		return FlagValues{}, err
	}

	return FlagValues{
		Length: passwordLength,
	}, nil
}
