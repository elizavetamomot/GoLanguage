package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//Строка - это байтовый слайс со своими особенностями при обращении к массиву
	wordEN := "Simple"
	wordRU := "Просто"
	fmt.Println(wordEN)
	fmt.Println(wordRU)
	//Какие значения байт сейчас находятся в слайсе wordEN, wordRU?
	fmt.Println("Bytes:")
	for i := 0; i < len(wordEN); i++ {
		fmt.Printf("%x ", wordEN[i]) //%x - формат представления 16-тиричного байта
	}
	fmt.Println()
	fmt.Println("Characters:")
	for i := 0; i < len(wordEN); i++ {
		fmt.Printf("%c ", wordEN[i]) //%c - формат представления символа
	}
	fmt.Println()
	fmt.Println("Bytes:")
	for i := 0; i < len(wordRU); i++ {
		fmt.Printf("%x ", wordRU[i])
	}
	fmt.Println()
	fmt.Println("Characters:")
	for i := 0; i < len(wordRU); i++ {
		fmt.Printf("%c ", wordRU[i])
	}
	fmt.Println()

	//Строки в Go - хранятся как наборы UTF-8 символов.
	//Каждый символ может занимать больше 1 байт

	//Руна (Rune) - это стандартный встроенный тип в Go (alias над int32),
	//позволяющий хранить единый неделимый UTF символ вне зависимости
	//от того только байт он занимает
	fmt.Println("Runes:")
	runeSlice := []rune(wordRU) // преобразование слайса байт к слайсу рун []byte{sliceRune}
	for i := 0; i < len(runeSlice); i++ {
		fmt.Printf("%c ", runeSlice[i])
	}
	fmt.Println()
	//Итерирование по строке с использованием рун
	for id, runeVal := range wordRU {
		fmt.Printf("%c starts at position %d\n", runeVal, id)
	}
	//id - это индекс байта, с которого начинается руна runeVal

	//Создание строки из слайса байт
	myByteSlice := []byte{0x40, 0x41, 0x42, 0x43}
	myStr := string(myByteSlice)
	fmt.Println(myStr)

	//можно использовать 10ое представление байтов
	myDecimalByteSlice := []byte{100, 101, 102, 103}
	myDecimalStr := string(myDecimalByteSlice)
	fmt.Println(myDecimalStr)

	//Создание строки из слайса рун
	//Руны как hex
	runeHexSlice := []rune{0x45, 0x46, 0x47, 0x48}
	myStrFromRume := string(runeHexSlice)
	fmt.Println("From Runes(hex):", myStrFromRume)
	//Руны как литералы
	runeLiteralSlice := []rune{'v', 'a', 's', 'y', 'a'} // '' - в таких символах обозначаются руны
	myStrFromLiteralSlice := string(runeLiteralSlice)
	fmt.Println("From Runes(literals):", myStrFromLiteralSlice)

	// Длина и ёмкость строки
	//Длина - количество байт в слайсе
	fmt.Println("Lenth of Вася", len("Вася"), "bytes")
	//Длина RuneCounter - количество рун
	fmt.Println("Length of Вася:", utf8.RuneCountInString("Вася"), "runes")
	//Вычисление ёмкости строки - бессмыслено, т.к строка базовый тип
	fmt.Println(cap([]rune("Вася")))

	//Сравнение строк == и !=. Начиная с go 1.6
	word1, word2 := "Вася", "Петя"
	if word1 == word2 {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}

	//Конкатенация
	//создается новая строка
	word3 := word1 + word2
	fmt.Println(word3)

	//Строитель строк
	firstName := "Alex"
	secondName := "Doe"
	fullName := fmt.Sprintf("%s ##### %s", firstName, secondName)
	fmt.Println("Fullname:", fullName)

	//Строки неизменяемые
	//fullName[0] = 'Q' - ошибка

	//А слайсы изменяемые
	fullNameSlice := []rune(fullName)
	fullNameSlice[0] = 'F'
	fullName = string(fullNameSlice)
	fmt.Println("String mutation:", fullName)

	//Сравнение рун
	if 'Q' == 'q' {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}

	//Где живут полезные методы работы со строками?
	//import "strings"

}
