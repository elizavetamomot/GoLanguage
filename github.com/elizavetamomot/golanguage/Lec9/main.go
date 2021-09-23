package main

import "fmt"

func main() {
	//слайсы (срезы) - динамическая обвязка над массивом
	//это набор ссылок на элементы
	//созданик слайса, основываясь на уже существующем массиве
	startArr := [4]int{10, 20, 30, 40}
	var startSlice []int = startArr[0:2] //слайс инициализируется пустыми квадратными скобками
	fmt.Println("Slice[0:2]:", startSlice)

	//создание слайса без явной инициализации массива
	secondSlice := []int{15, 20, 30, 45}
	fmt.Println("SecondSlice:", secondSlice)

	//изменение элемента среза
	originArr := [...]int{11, 22, 33, 44, 55, 66}
	fSlice := originArr[1:4]
	for i, _ := range fSlice {
		fSlice[i]++
	}
	fmt.Println("OriginArr:", originArr)
	fmt.Println("fSlice:", fSlice)

	//Один массив и два производных среза
	firstOrigin := originArr[:]
	seconOrigin := originArr[2:5]

	fmt.Println("Before modifications: Arr:", originArr, "\nfirstOrigin:",
		firstOrigin, "seconOrigin:", seconOrigin)
	firstOrigin[3]++
	seconOrigin[1]++
	fmt.Println("After modifications: Arr:", originArr, "\nfirstOrigin:",
		firstOrigin, "seconOrigin:", seconOrigin)

	//срез как встроенный тип
	/* 	type slice struct {
		Length int
		Capacity int
		ZeroElement *byte
	} */

	//длина и ёмкость слайса
	wordsSlice := []string{"one", "two", "three"}
	fmt.Println("slice:", wordsSlice, "length:", len(wordsSlice),
		"Capacity:", cap(wordsSlice))
	wordsSlice = append(wordsSlice, "four")
	fmt.Println("slice:", wordsSlice, "length:", len(wordsSlice),
		"Capacity:", cap(wordsSlice))

	//Capacity (или ёмкость слайса) - это значение, показывающее,
	//сколько элементов впринципе можно добавить в срез
	//без выделения дополнительной памяти под ниже лежащий массив

	/*
		Допустим есть срез на 3 элемента (инициализировали
		без явного указания массива).
		Компилятор при создании это среза создал массив ровно на 3 элемента.
		После этого компилятор вернул адрес, где этот массив живет.
		Срез запомнил этот адрес и теперь ссылается на него.
		Начинаем деформировать слайс (увеличим количество элементов).
		Проблема - в нижележащем массиве (на котором основан слайс) всего 3 ячейки
		Что делать?
		Компилятор ищет в памяти место для массива размера 3*2
		(в общем случае n*2, где n - первоначальный размер)
		После того, как место найдено (в нашем случае
			найдено место для 6 элементов),
			в это место копируются старые 3 элемента на свои позиции.
		На 4-ую позицию мы добавляем новый элемент.
		После этого компилятор возвращает нашему слайсу новый адрес в памяти,
		где находится массив под 6 элементов
	*/

	//Ёмкость всегда будет изменяться как n*2
	numerics := []int{1, 2}
	for i := 0; i < 200; i++ {
		if i%10 == 0 {
			fmt.Println("Current len:", len(numerics),
				"Current cap:", cap(numerics))
		}
		numerics = append(numerics, i)
	}

	//После выделение памяти под новый массив, ссылки со старым будут перенесены в новый
	//Пример:
	numArr := [2]int{1, 2}
	numSlice := numArr[:]
	numSlice = append(numSlice, 3) //в этот момент numSlice больше не ссылается на numArr
	numSlice[0] = 10
	fmt.Println("numArr:", numArr, "numSlice:", numSlice)

	//Как создавать слайсы наиболее эффективно?
	// make() - функция позволяющая более детально создавать срезы
	sl := make([]int, 10, 15)
	/*
		[]int - это тип коллекции
		10 - это длина
		15 - это ёмкость
		Сначала инициализируется arr = [15]int
		Затем по нему делается срез arr[0:10]
		После чего он возвращается
	*/
	fmt.Println(sl)

	//Добавление элементов в срез
	//Срез это интерфейс, который позволяет наиболее удобно работать с массивом
	myWords := []string{"one", "two", "three"}
	fmt.Println("myWords:", myWords)
	anotherWords := []string{"four", "five", "six"}
	myWords = append(myWords, anotherWords...)
	fmt.Println("myWords:", myWords)

	//Многомерный срез
	mSlice := [][]int{
		{1, 2},
		{3, 4, 5, 6},
		{10, 20, 30},
		{},
	}
	fmt.Println(mSlice)
}
