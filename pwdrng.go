package main

import (
	"fmt"

	"github.com/doctorpoor/pwdrng/internal/generator"
)

func main() {
	fmt.Printf("Generated password: \n%v\n", generator.GeneratePassword())
}
