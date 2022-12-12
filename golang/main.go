package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	const a float64 = 2.0
	const b float64 = 0.95
	var s int
	fmt.Scan(&s)
	var k int
	var n, m int
	var f int
	var phrase string
	var x, z int
	var FA []int
	var SOM [][]int
	fmt.Scan(&s)
	fmt.Scan(&k)
	fmt.Scan(&FA)
	fmt.Scan(&SOM)
	fmt.Scan(&n)
	fmt.Scan(&f)
	fmt.Scan(&m)
	fmt.Scan(&phrase)
	fmt.Scan(&x)
	fmt.Scan(&z)
	fmt.Println("Задание А")
	fmt.Println(TaskA(a, b))
	fmt.Println("Задание Б")
	fmt.Println(TaskB(a, b))
	fmt.Println(EvenOrOdd(s))
	fmt.Println(CountingSheep())
	fmt.Println(CountTheMonkeys(k))
	fmt.Println(SchoolPaperwork(n, m))
	fmt.Println(PolishAlphabet(phrase))
	fmt.Println(FindAll(FA, f))
	fmt.Println(SummOfMin(SOM, x, z))
}

func TaskA(a, b float64) []float64 {
	var x float64 = 1.25
	var taskA []float64
	for ; x <= 2.75; x += 0.3 {
		taskA = append(taskA, formula(x, a, b))
	}
	return taskA
}

func TaskB(a, b float64) []float64 {
	var taskB []float64
	var array = []float64{2.2, 3.78, 4.51, 6.58, 1.2}
	for i := 0; i < len(array); i++ {
		taskB = append(taskB, formula(array[i], a, b))
	}
	return taskB
}

func formula(x, a, b float64) float64 {
	var y float64 = (1 + math.Pow(math.Log10(a/b), 2.0)) / (b - math.Pow(2.718, x/a))
	return y
}

func EvenOrOdd(s int) string {
	var result string
	if s%2 == 0 {
		result = "Even"
	} else if s%2 != 0 {
		result = "Odd"
	}
	return result
}

func CountingSheep() int {
	var sheeps = []string{"true", "true", "true", "false", "true", "false", "true", "true"}
	var count int = 0
	for i := 0; i < len(sheeps); i++ {
		if sheeps[i] == "true" {
			count++
		}
	}
	return count
}

func CountTheMonkeys(k int) []int {
	var numbers []int
	for i := 1; i <= k; i++ {
		numbers = append(numbers, i)
	}
	return numbers
}

func SchoolPaperwork(n, m int) int {
	var paper int
	if n > 0 && m > 0 {
		paper = n * m
	} else {
		paper = 0
	}
	return paper
}

func PolishAlphabet(phrase string) string {
	var dictionary [][]string = [][]string{
		{"ą", "a"},
		{"ć", "c"},
		{"ę", "e"},
		{"ł", "l"},
		{"ń", "n"},
		{"ó", "o"},
		{"ś", "s"},
		{"ź", "z"},
		{"ż", "z"},
	}
	for i := 0; i < len(dictionary); i++ {
		phrase = strings.Replace(phrase, dictionary[i][0], dictionary[i][1], 1)
	}
	return phrase
}

func FindAll(FA []int, f int) []int {
	var res []int
	for i := 0; i < len(FA); i++ {
		if FA[i] == f {
			res = append(res, i)
		}
	}
	return res
}

func SummOfMin(SOM [][]int, x, z int) int {
	var res int = 0
	for i := 0; i < (x); i++ {
		min := math.Inf(1)
		for k := 0; k < (z); k++ {
			if float64(SOM[i][k]) < min {
				min = float64(SOM[i][k])
			}
		}
		res = res + int(min)
	}
	return res
}
