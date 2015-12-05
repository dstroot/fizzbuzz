// Fizzbuzz writes "fizz" for multiples of three, "buzz" for multiples
// of five, and "fizzbuzz" for multiples of both three and five. it requires
// configuration and takes no input.
package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(fizzbuzz(i))
	}
}

// fizzbuzz function takes an int and returns a string. The string will
// contain the int passed in, or "fizz" for multiples of three, "buzz"
// for multiples of five, or "fizzbuzz" for multiples of three *and* five.
func fizzbuzz(n int) string {
	if n%5 == 0 && n%3 == 0 {
		return "fizzbuzz"
	}
	if n%5 == 0 {
		return "buzz"
	}
	if n%3 == 0 {
		return "fizz"
	}
	// two different ways below to illustrate the power of benchmarking
	// return strconv.Itoa(n)
	return fmt.Sprintf("%v", n)
}
