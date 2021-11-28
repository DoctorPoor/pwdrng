package main

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/doctorpoor/pwdrng/internal/generator"
)

func main() {
	generatedPassword := generator.GeneratePassword()
	clipboard.WriteAll(generatedPassword)
	fmt.Printf("Copied password to clipboard: \n%v\n", generatedPassword)
}
