package generator

import (
	"math/rand"
	"strings"
	"time"
)

func getRandomFromArray(input []string) string {
	return input[rand.Intn(len(input))]
}

func getRandomLowerCase() string {
	return getRandomFromArray([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"})
}

func getRandomUpperCase() string {
	return getRandomFromArray([]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"})
}

func getRandomInteger() string {
	return getRandomFromArray([]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"})
}

func getRandomSpecialCharater() string {
	return getRandomFromArray([]string{" ", "!", "#", "$", "%", "&", "'", "(", ")", "*", "+", ",", "-", ".", "/", ":", ";", "<", "=", ">", "?", "@", "[", "\\", "]", "^", "_", "`", "{", "|", "}", "~ "})
}

func getRandomCharacter(getSingleCharacterFunctions []func() string, arrayLength int) string {
	return getSingleCharacterFunctions[rand.Intn(arrayLength)]()
}

func shuffleString(input string) string {
	stringSlice := strings.Split(input, "")
	rand.Shuffle(len(stringSlice), func(i, j int) {
		stringSlice[i], stringSlice[j] = stringSlice[j], stringSlice[i]
	})
	return strings.Join(stringSlice, "")
}

func GeneratePassword() string {
	rand.Seed(time.Now().Unix())

	pwd := ""
	pwd += getRandomLowerCase()
	pwd += getRandomUpperCase()
	pwd += getRandomInteger()
	pwd += getRandomSpecialCharater()

	getSingleCharacterFunctions := make([]func() string, 4)
	getSingleCharacterFunctions[0] = getRandomLowerCase
	getSingleCharacterFunctions[1] = getRandomUpperCase
	getSingleCharacterFunctions[2] = getRandomInteger
	getSingleCharacterFunctions[3] = getRandomSpecialCharater
	functionArrayLength := len(getSingleCharacterFunctions)

	for i := 0; i < 12; i++ {
		pwd += getRandomCharacter(getSingleCharacterFunctions, functionArrayLength)
	}

	return shuffleString(pwd)
}
