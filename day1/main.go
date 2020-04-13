package main

import (
	"ch01/lib"
	"fmt"
)

func main() {
	arr := [4]int{3, 2, 5, 4}
	long = arr

	z := []int{}

	z = append(z, 20)
	fmt.Println(z[0], len(z))

	y := make([]int, 3)

	fmt.Println(y[1])
	fmt.Println("****** Variadic Functions ******")

	x := []int{5, 10, 15, 20, 25}

	for i := 0; i < 5; i++ {
		fmt.Println(x[i])
	}

	sum := Sum(x...)

	fmt.Printf("Sum: %d\n", sum)

	fmt.Println("****** Variadic Functions ******")

	fmt.Printf("Testing %d", Add(1, 2))
	fmt.Println("****** Default favorite packages ******")
	lib.PrintFavorites()
	// Add couple of favorites
	lib.Add("github.com/dgrijalva/jwt-go")
	lib.Add("github.com/onsi/ginkgo")
	fmt.Println("****** All favorite packages ******")
	lib.PrintFavorites()
	count := len(lib.GetAll())
	fmt.Printf("Total packages in the favorite list:%d", count)
}

//Add the two number
func Add(x int, y int) int {
	return x + y
}

//Sum should given list of numebr
func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
