package main

import (
	"fmt"
	"math/rand"
)

func main() {
	passwordGen(30)
}

func passwordGen(qtd int) {
	charactersLower := "abcdefghijklmnopqrstuvwxyz"
	charactersUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	characters := "!@#$%¨&*()_-{}^~:?;>]"

	list := []string{charactersLower, charactersUpper, numbers, characters}

	for i := 0; i < qtd; i++ {
		randomList := rand.Intn(len(list))
		character := list[randomList]

		randomCharacter := rand.Intn(len(character))

		fmt.Print(string(character[randomCharacter]))
	}
}