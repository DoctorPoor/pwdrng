package main

import (
	"fmt"

	"github.com/atotto/clipboard"

	"github.com/doctorpoor/pwdrng/internal/flags"
	"github.com/doctorpoor/pwdrng/internal/generator"
)

func main() {
	flagValues, err := flags.GetFlagValues()
	if err != nil {
		fmt.Println(err)
		return
	}
	generatedPassword := generator.GeneratePassword(flagValues.Length)
	clipboard.WriteAll(generatedPassword)
	fmt.Printf("Copied password to clipboard: \n%v\n", generatedPassword)
}
