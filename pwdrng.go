package main

import (
	"fmt"

	"github.com/atotto/clipboard"

	"github.com/doctorpoor/pwdrng/internal/arguments"
	"github.com/doctorpoor/pwdrng/internal/generator"
)

func main() {
	passwordLength, err := arguments.GetPasswordLength()
	if err != nil {
		fmt.Println(err)
		return
	}
	generatedPassword := generator.GeneratePassword(passwordLength)
	clipboard.WriteAll(generatedPassword)
	fmt.Printf("Copied password to clipboard: \n%v\n", generatedPassword)
}
