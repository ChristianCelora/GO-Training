package main

/**
* import(
*   "fmt"
*   "errors"
* )
*/
import "fmt"
import "errors"

func sum(a int, b int) int {
	return a + b
}

func sumFloat(a float32, b float32) float32 {
	return a + b
}

func sumMultiple(nums...int) (int) {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	return sum
}

/** 
* In golang there is no operator overloading. You can't define your own +.
*/
// func sumGeneric[T any](a T, b T) T {
// 	return a + b
// }

func basicOps(a int, b int) (int, int, int, int) {
	sum := a + b
	diff := a - b
	mul := a * b
	div := a / b
	return sum, diff, mul, div
}

func getArrayCursor() func(array []int) (int, error) {
	i := -1
	return func(array []int) (int, error) {
		i++
		if (i >= len(array)) {
			return -1, errors.New("array out of bounds")
		}
		return array[i], nil
	}
}

func main() {
	var a = 4
	var b = 5
	var c int
	c = sum(a, b)
	fmt.Println("sum: ", a, "+", b, "=", c)

	var af float32 = 4.5
	var bf float32 = 5.3
	// c = sum(af, bf) // cannot use af (variable of type float32) as type int in argument to sum
	var cf float32
	cf = sumFloat(af, bf)
	fmt.Println("sumFloat: ", af, "+", bf, "=", cf)

	// c = sumGeneric(a, b)
	// cf = sumGeneric(af, bf)
	// fmt.Println("sumGeneric (int): ", a, "+", bf, "=", c)
	// fmt.Println("sumGeneric (float32): ", a, "+", bf, "=", c)

	sum, diff, mul, div := basicOps(a, b)
	fmt.Println("values: ", a, b)
	fmt.Println("sum: ", sum, "| diff:", diff, "| mul:", mul, "| div:", div)

	// If you only want a subset of the returned values, use the blank identifier _.
	// there can be multiple _
	_, diff, mul, _ = basicOps(a, b)
	fmt.Println("diff:", diff, "| mul:", mul)

	var sum2 int
	sum2 = sumMultiple(1, 2, 3, 4)
	fmt.Println("1+2+3+4", ", sumMultiple:", sum2)
	values := [] int{1, 2, 3, 4}
	sum2 = sumMultiple(values...)
	fmt.Println("vals", values, ", sumMultiple:", sum2)

	// a function can return an anonymus function (closure)
	nextElem := getArrayCursor()
	fmt.Println(nextElem(values))
	fmt.Println(nextElem(values))
	fmt.Println(nextElem(values))
	fmt.Println(nextElem(values))
	fmt.Println(nextElem(values))
	fmt.Println(nextElem(values))
}