package main

import "fmt"

func main() {
	//Массивы. Основы
	//Создание массива под хранение 5-ти целочисленных элементов
	var arr [5]int //при инициализации массива важно передать информацию о количестве элементов
	fmt.Println("This is my array:", arr)
	//Определение элементов массива (после предварительной инициализации)
	//Необходимо обратиться к элементу массива через синтаксис arr[i] = elem
	arr[0] = 10
	arr[1] = 20
	arr[3] = -400
	arr[4] = 10000000
	fmt.Println("This is init array:", arr)
	//определение массива с указанием элементов на месте
	newArr := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Short declar:", newArr)
	//создание массива через инициализацию переменных
	ArrValues := [...]int{10, 20, 30, 40}
	fmt.Println("Arr decl values:", ArrValues, "Lenght:", len(ArrValues))
	// массив это набор ЗНАЧЕНИЙ, то есть при всех манипуляциях - массив
	//копируется (жестко, на уровне компиляции)
	first := [...]int{1, 2, 3}
	second := first
	second[0] = 10000
	fmt.Println("First", first)
	fmt.Println("Second", second)
	//массив и его размер - это две составляющие одного типа (Размер массива - часть типа)
	/* var aArr [5]int
	var bArr [6]int
	aArr[0] = 100
	bArr = aArr */
	//Итерирование по массиву
	floatArr := [...]float64{12.5, 13.5, 15, 16.2, 17.8}
	for i := 0; i < len(floatArr); i++ {
		fmt.Printf("%d element of arr is %.2f\n", i, floatArr[i])
	}
	//иетерирование по массиву через оператор range
	var sum float64
	for id, val := range floatArr {
		sum += val
		fmt.Printf("%d elemt of arr is %.2f. Sum = %.2f\n", id, val, sum)
	}
	// игнорирование id в range based for цикле
	for _, val := range floatArr {
		fmt.Printf("elemt of arr is %.2f\n", val)
	}
	//многомерные массивы
	words := [2][2]string{
		{"Bob", "Pop"},
		{"Victor", "Vika"},
	}
	fmt.Println("Multidimensional:", words)
	//итерирование по многомерному массиву
	for _, val1 := range words {
		for _, val2 := range val1 {
			fmt.Printf("%s ", val2)
		}
		fmt.Println()
	}
}
