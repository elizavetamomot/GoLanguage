package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	//Boolean => default false
	var firstBoolean bool
	fmt.Println(firstBoolean)
	//Boolean operands
	aBool, bBool := true, false
	fmt.Println("AND", aBool && bBool)
	fmt.Println("OR", aBool || bBool)
	fmt.Println("NOT", !aBool)

	//Numerics. Integers
	//int8, int16, int32, int64, int - числа это биты
	//uint8, uint16, uint32, uint64, uint - беззнаковые
	//+ - * / %
	var a int = 32
	b := 92
	fmt.Println("Value a:", a, "Value b", b, "Value a+b:", a+b)
	//Выбод типа через %Т форматирование
	fmt.Printf("Type is %T\n", a)
	//Сколько байт занимает переменная типа int
	fmt.Printf("Type %T size of %d bytes\n", a, unsafe.Sizeof(a))
	//При использовании короткого объявления - тип для целого числа int платформо-зависимый
	fmt.Printf("Type %T size of %d bytes\n", b, unsafe.Sizeof(b))

	//Использование явного приведение типа при необходимости
	var first32 int32 = 12
	var second64 int64 = 13
	fmt.Println(int64(first32) + second64)

	//int != int64, даже если int занимает 8 байт
	var third64 int64 = 12121212
	var fourthInt int = 111111
	fmt.Println(third64 + int64(fourthInt))

	//Numerics. Float
	//float32, float64
	//+ - * /
	floatFirst, floatSecond := 23.43, 55.555555555
	fmt.Printf("Type of a %T and type of b %T\n", floatFirst, floatSecond)
	fmt.Printf("%.3f \n", floatSecond)

	//Numeric. Complex
	c1 := complex(5, 7)
	c2 := 12 + 32i
	fmt.Println(c1 + c2)
	fmt.Println(c1 * c2)

	//Strings Строка - набор БАЙТ
	name := "Fedor"
	kiril := "Федя"
	lastname := "Ivanov"
	concatinacia := name + " " + lastname
	fmt.Println("Full name:", concatinacia)
	fmt.Println("Length:", name, len(name), kiril, len(kiril))
	fmt.Println("Amount od chars:", kiril, utf8.RuneCountInString(kiril))
	//Rune - руна, один utf символ
	//Поиск подстроки в строке
	totalString, subString := "ABCDEFG", "FG"
	fmt.Println(strings.Contains(totalString, subString))
	//rune -> alias int32
	var sampleRune rune
	var anotherRune rune = 'Q' //руны в ''
	fmt.Println(sampleRune)
	fmt.Printf("Rune as char %c \n", sampleRune)
	fmt.Printf("Rune as char %c \n", anotherRune)
	// "A" < "abcd"
	fmt.Println(strings.Compare("abcd", "abcd"))   // 0 if a ==b
	fmt.Println(strings.Compare("abcd", "ab"))     // 1 if a > b
	fmt.Println(strings.Compare("abcd", "abcdef")) //-1 if a < b

	//alias uint8
	var by byte
	fmt.Println(by)
}
