package main

import "fmt"

func main() {
	// Call to Variadic functions
	fmt.Println(sum("Alex", 12, 45, 78, 56))
	fmt.Println(sum("Franco", 10, 20, 30, 40, 50))
	printData("ads", 22, true, 3.11)
	// Call to Recursive functions
	fmt.Println(factorial(3))
	// Call to Anonymous functions
	r1 := apply(double, 3)
	r2 := apply(triple, 3)
	fmt.Println(r1, r2)
	// Call to Higher-order functions
	r := doubleValue(addOne, 3)
	fmt.Println("Result: ", r)
	// Call to Closure functions
	nextInt := incrementer()
	fmt.Println(nextInt())
}

// Variadic function
func sum(name string, nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	fmt.Printf("Hello %s, the sum is: %d\n", name, total)
	return total
}

func printData(data ...interface{}) {
	for _, d := range data {
		fmt.Printf("%T - %v\n", d, d)
	}
}

// Recursive function
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// Anonymous function
func double(n int) int {
	return n * 2
}

func triple(n int) int {
	return n * 3
}

func apply(f func(int) int, n int) int {
	return f(n)
}

// Higher-order function
func doubleValue(f func(int) int, x int) int {
	return f(x * 2)
}

func addOne(x int) int {
	return x + 1
}

// Closure function
func incrementer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
