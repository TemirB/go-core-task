package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
)

func main() {
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	fmt.Println("Типы переменных:")
	fmt.Printf("numDecimal: %s\n", reflect.TypeOf(numDecimal))
	fmt.Printf("numOctal: %s\n", reflect.TypeOf(numOctal))
	fmt.Printf("numHexadecimal: %s\n", reflect.TypeOf(numHexadecimal))
	fmt.Printf("pi: %s\n", reflect.TypeOf(pi))
	fmt.Printf("name: %s\n", reflect.TypeOf(name))
	fmt.Printf("isActive: %s\n", reflect.TypeOf(isActive))
	fmt.Printf("complexNum: %s\n", reflect.TypeOf(complexNum))

	combinedString := fmt.Sprintf("%d %O %X %.2f %s %t %g",
		numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println("\nОбъединенная строка:", combinedString)

	runeSlice := []rune(combinedString)

	tmp := append([]rune("go-2024"), runeSlice[len(runeSlice)/2:]...)
	saltedRunes := append(runeSlice[:len(runeSlice)/2], tmp...)

	hash := sha256.Sum256([]byte(string(saltedRunes)))

	fmt.Printf("\nSHA256 hash: %x\n", hash)
}
