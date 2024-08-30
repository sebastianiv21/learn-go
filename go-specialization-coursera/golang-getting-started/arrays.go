package main

import "fmt"

func main() {
	// initialize an array
	// var name [array_length]type
	var myArray [5]int = [5]int{1, 2, 3, 4, 5}
	// array values are initialized with the zero value of the type

	myArray[0] = 2

	// initialize an array with values (array literal)
	myArrayLiteral := [5]int{1, 2, 3, 4, 5}

	fmt.Println(myArrayLiteral[2])

	// can also declare array length with [...],
	// this way, the length is infered by the numbers inside {}
	inferredLengthArray := [...]string{"abc", "def", "ghi"}

	fmt.Println(inferredLengthArray)

	// iterating through arrays
	arrayToIterate := [3]int{1, 2, 3}

	for i, v := range arrayToIterate {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}

  // slices
  // slices are like arrays, but they are dynamic
  // slices are like references to arrays
  arr := [...]{"a", "b", "c", "d", "e"}

  // slice syntax
  s1 := arr[1:3] // [b c]
  s2 := arr[:4] // [a b c d]
  // len() and cap() functions
  // len() returns the length of the slice
  // cap() returns the capacity of the slice
  // capacity is the number of elements in the underlying array starting from the first element in the slice
  s1Len := len(s1) // 2
  s1Cap := cap(s1) // 4

  // make() function
  // make() function is used to create slices, maps and channels
  slice := make([]int, 5, 10) // []int{0, 0, 0, 0, 0} len: 5, cap: 10

  // append() function
  // append() function is used to append elements to slices
  // if the capacity of the slice is exceeded, the underlying array is copied to a new array with double the capacity
  slice = append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10) // len: 15, cap: 20

  // use make() function to create a map
  // map[key_type]value_type
  myMap := make(map[string]int)
}
