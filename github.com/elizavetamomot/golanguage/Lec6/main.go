package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		for init; condition; post {
			init - блок инициализации переменных цикла
			condition - условие (если верно - то тело цикла выполняется)
			post - изменение переменной цикла
		}
	*/
	for i := 0; i <= 5; i++ {
		fmt.Printf("Current value: %d \n", i)
	}
	//в качестве init может быть использовано выражение присваения :=
	// break - команда, прирывающая текущее выполнение тела цикла и
	//передающее управление иструкциям, слующим за циклом
	for i := 0; i <= 15; i++ {
		if i > 12 {
			break
		}
		fmt.Printf("Current value: %d \n", i)
	}
	fmt.Println("After foor loop with BREAK")
	//continue - команда, прерывающая текущее выполение тела цикла и
	//передающая управление слудующей итерации цикла
	for i := 0; i <= 15; i++ {
		if i > 10 && i <= 13 {
			continue
		}
		fmt.Printf("Current value: %d \n", i)
	}
	fmt.Println("After foor loop with CONTINUE")

	//вложенные циклы и лейблы
	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("Треугольник")

	//лейбы
outer:
	for i := 0; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i:%d and j:%d and sum i+j=%d \n", i, j, i+j)
			if i == j {
				break outer //чтобы все циклы остановились
			}
		}
	}

	//модификации цикла for
	//while do
	var loopVar int = 0
	for loopVar < 10 {
		fmt.Println(loopVar)
		loopVar++
	}
	//бесконечный цикл
	var password string

	for {
		fmt.Print("Password: ")
		fmt.Scan(&password)
		if strings.Contains(password, "1234") {
			fmt.Println("Weak. Try again")
		} else {
			fmt.Println("Accepted")
			return
		}
	}

}
