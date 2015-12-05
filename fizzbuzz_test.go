package main

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// TestFizzbuzz will test for the correct functionality of our fizzbuzz
// function. It defines how the function should behave.
func TestFizzbuzz(t *testing.T) {

	Convey("if integer is multiple of three return \"fizz\"", t, func() {
		So(fizzbuzz(3), ShouldEqual, "fizz")
	})

	Convey("If integer is multiple of five return \"buzz\"", t, func() {
		So(fizzbuzz(5), ShouldEqual, "buzz")
	})

	Convey("If integert is multiple of three and five return \"fizzbuzz\"", t, func() {
		So(fizzbuzz(15), ShouldEqual, "fizzbuzz")
	})

	Convey("otherwise just return the number", t, func() {
		So(fizzbuzz(1), ShouldEqual, "1")
	})

}

// Godoc examples are snippets of Go code that are displayed as package
// documentation and that are verified by running them as tests. They
// can also be run by a user visiting the godoc web page for the package
// and clicking the associated "Run" button.
//
// Having executable documentation for a package guarantees that the
// information will not go out of date.
//
// As with typical tests, examples are functions that reside in a
// package's _test.go files. Unlike normal test functions, though,
// example functions take no arguments and begin with the word
// Example instead of Test.  See http://blog.golang.org/examples
func ExampleFizzbuzz() {
	fmt.Println(fizzbuzz(15))
	// Output: fizzbuzz
}

// For benchmarking, Go's testing package will take care of increasing
// the variable b.N for any function we include within the for loop, so
// that we'll benchmark our function fizzbuzz repeatedly for a minimum
// of 1 second. This is so that we can get statistically significant
// results with sufficient repetitions. All we need to type to run our
// tests is: `go test -bench=.`
func BenchmarkFizzbuzz(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fizzbuzz(n)
	}
}
