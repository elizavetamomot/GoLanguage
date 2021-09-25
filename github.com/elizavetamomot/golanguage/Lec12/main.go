package main

import "fmt"

// Явная функция - локально-определенный блок кода имеющий имя (явное определение)
//Функцию необходимо определить + Функцию необходимо вызвать

//Сигнатеры функции и их определение
/*
func functionName(params type) typeReturnValue {
	body
}
*/

//Простейшая функция
//Определить функцию можно до или после или в любом месте пакета func main()
func add(a int, b int) int {
	result := a + b
	return result
}

func main() {
	fmt.Println("H")
	//Вызов простейшей функции
	res := add(10, 20)
	fmt.Println("Reasult a + b =", res)
	fmt.Println("Result 1*2*3*4 = ", mult(1, 2, 3, 4))

	per, area := rectangleParameters(10.5, 10.5)
	fmt.Println("Per 10.5 x 10.5 = ", per, "Area =", area)

	SecPer, SecArea := namedReturn(11, 10)
	fmt.Println("SecPer = ", SecPer, "SecArea = ", SecArea)

	helloVariadic(10, 20, 30, 40)
	helloVariadic(10)

	someStrings(2, 3, "Bob", "Alex")

	fmt.Println()
	sum1 := enterVariatic(10, 30, 40, 50, 60)
	sliceNumber := []int{10, 22, 30}
	sum2 := enterVariatic(sliceNumber...)
	fmt.Printf("Sum1 = %d Sum2 = %d\n", sum1, sum2)

	//	sum11 := enterSlice(10, 30, 40, 50, 60)
	sliceNumber11 := []int{10, 22, 30}
	sum22 := enterSlice(sliceNumber11)
	fmt.Printf("Sum1 = %d Sum2 = %d\n", sum1, sum22)

	//Анонимная функция
	anon := func(a, b int) int {
		return a + b
	}
	fmt.Println(anon(1, 2))
}

//Функция с однотипными параметрами
func mult(a, b, c, d int) int {
	return a * b * c * d
}

//Возврат больше чем одного значения
func rectangleParameters(length, width float64) (float64, float64) {
	var perimeter = 2 * (length + width)
	var area = length * width
	return perimeter, area
}

//Именнованный возврат значений
func namedReturn(a, b int) (perimeter, area int) {
	perimeter = 2 * (a + b)
	area = a * b
	return // не нужно в этом случае указывать переменные
}

//При вызове оператора return функция прекращает своё выполнение и возвращает что-то
func funcWithReturn(a, b int) int {
	if a > b {
		return a - b
	} else if a == b {
		return b
	} else {
		return b
	}
	// return b //недостижимый код
}

// Variadic parameters (континуальные параметры)
func helloVariadic(a ...int) {
	fmt.Println(a)
	fmt.Printf("value %v and type %T\n", a, a)
}

//Смешение параметр с variadic
//variadic всегда в конце
//variadic всегда один в функции
func someStrings(a, b int, words ...string) {
	fmt.Println("Parameter:", a)
	fmt.Println("Parameter:", b)
	var result string
	for _, word := range words {
		result += word
	}
	fmt.Println("Result concat:", result)
}

//Передача слайса или использование variatic? Variadic
func enterVariatic(nums ...int) int {
	var sum int
	for _, val := range nums {
		sum += val
	}
	return sum
}

func enterSlice(nums []int) int {
	var sum int
	for _, val := range nums {
		sum += val
	}
	return sum
}

//Анонимная функция внутри явной
func bigFunc(aArg, bArg int) int {
	return func(a, b int) int { return a + b }(aArg, bArg)
}
