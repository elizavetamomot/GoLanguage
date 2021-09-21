package main

import "fmt"

func main() {
	//Классический условный оператор
	/*	if condition {
			body
		} else {
			body
		}
	*/
	var value int
	fmt.Scan(&value)

	if value%2 == 0 {
		fmt.Println("Число", value, "четное")
	} else {
		fmt.Println("Число", value, "не четное")
	}

	//Инициализация в блоке условного оператора
	//блок присваивания только :=
	//инициалируемая переменная видна тольковнутри области
	//видимости условного оператора (в телах if, else if, else)
	if num := 10; num%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("ODD")
	}

	//Правило номер 1: В Go стараются избегать блоков else
	if height := 101; height > 100 {
		fmt.Println("height > 100")
		return
	}
	fmt.Println("height <= 100")
}
