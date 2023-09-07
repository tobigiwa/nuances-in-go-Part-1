package main

import "fmt"

func main() {

	arrayOfInt := []int{0, 1, 2}
	for range arrayOfInt {
		arrayOfInt = append(arrayOfInt, 10)
	}

	array2OfInt := []int{0, 1, 2}
	for i := 0; i < len(array2OfInt); i++ {
		array2OfInt = append(array2OfInt, 10)
	}
	fmt.Println(array2OfInt)
}
