package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("slice: ", slice)
	slice2 := slice[1:]
	slice2[0] = 10
	fmt.Println("slice: ", slice)
	fmt.Println("slice2: ", slice2)

	m := make(map[int]string, 10)
	m[0] = "00000"
	fmt.Println("m: ", m)
	m1 := m
	m1[0] = "111111"
	fmt.Println("m: ", m)
	fmt.Println("m1: ", m1)
	//m2[0] = "22222222"
	fmt.Println("m: ", m)
	fmt.Println("m2: ", m2)
}
