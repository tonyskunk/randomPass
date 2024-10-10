package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var passwordLenght int
	var charSet string
	fmt.Print("Введите длину пароля: ")
	fmt.Scan(&passwordLenght)
	fmt.Print("Введите набор символов:")
	fmt.Scan(&charSet)

	charSetRune := []rune(charSet)
	rand.Seed(time.Now().UnixNano())

	password := make([]rune, passwordLenght)
	for i := range password {
		password[i] = charSetRune[rand.Intn(len(password))]

	}
	fmt.Println("Сгенерированный пароль:", string(password))

}
