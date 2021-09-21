package main

import "fmt"

func main() {
	fmt.Println("Wello", "Hello121212")
	fmt.Println("Helllo")
	//Форматированный вывод
	fmt.Printf("Hello, %s. Current year is %d \n", "User", 2021)
	//Декларирование переменных
	var age int
	fmt.Println("My age is:", age)
	age = 32
	fmt.Println("Age after assignment:", age)
	//Декларирование и инициализация пользовательским значением
	var height int = 186
	fmt.Println("My height is:", height)
	//Полустрогость типизации
	var uid = 12345
	fmt.Printf("My uid: %d. Type = %T\n", uid, uid)
	//Декларирование и инициализация переменных одного типа(множественный случай)
	var first, second int = 1, 2
	fmt.Printf("first = %d, second = %d\n", first, second)
	//Декларирование блока переменных
	var (
		personName string = "Bob"
		personAge  int    = 42
		personUID  int
	)

	fmt.Printf("Name: %s\nAge: %d\nUID:%d\n", personName, personAge, personUID)

	//особенности
	var a, b = 10, "volvo"
	fmt.Println(a, b)

	//короткое объявление
	count := 10
	strr := "Count:"
	fmt.Println(strr, count)
	//множественное короткое объявление
	aArg, bArg := 10, "strin"
	fmt.Println(aArg, bArg)
	//нельзя снова aArg, bArg := 30, "newStr", но
	//особенности
	bArg, cArg := "newstr", 40
	fmt.Println(aArg, bArg, cArg)
}
