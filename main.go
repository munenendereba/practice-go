package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("we here")

	n1 := []int{10, 20, 30, 40, 50, 60, 7, 8, 9, 6, 7, 8, 99, 909, 56, 55, 44, 34}
	var n2 = []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 2, 3, 11, 222, 1, 21, 222, 33}
	var n3 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1}
	var myarr [9]string
	myarr[0] = "123"
	myarr[8] = "hall"

	n4 := make([]int, 6)
	n4 = append(n4, 2, 3, 4, 5, 8)

	fmt.Println("this is it", n4)

	fmt.Println(len(n1), cap(n1), len(n2), len(n3), cap(n3))
	fmt.Println(myarr)

	n1 = append(n1, 100)
	fmt.Println(len(n1), cap(n1))
	fmt.Println(n1)

	lang := "golang"

	fmt.Printf(`the value of lang is %s and the type is %T, and the pointer is %p`, lang, lang, &lang)

	myMap := make(map[string]string)

	myMap["one"] = "two"
	myMap["three"] = "four"
	myMap["izzo"] = "nesh"
	myMap["jb"] = "tugi"
	fmt.Println()

	fmt.Println(myMap)

	myMap["jb"] = "jabali"
	fmt.Println(myMap)

	myMap["k"] = "pp"

	fmt.Println(myMap)

	for index, n := range myMap {
		fmt.Println(index, n)
	}

	var j interface{} = 56.7

	str, ok := j.(string)

	if ok {
		fmt.Println("this is a string", str)
	} else {
		fmt.Printf("it's not a string, but a %T\n", j)
	}

	fmt.Println(time.Now())
	fmt.Println(time.Now())
	fmt.Println(time.Now())

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	n := 9
	fmt.Printf("factorial of %d is %d\n", n, factorial(n))

	p := 14
	fmt.Printf("is number %d a prime? %t\n", p, isPrime(p))

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
