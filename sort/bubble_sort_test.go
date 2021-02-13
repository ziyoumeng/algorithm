package sort

import "fmt"

func Example_BubbleSort(){
	arr := []int{3,2,1,4,5}
	BubbleSort(arr)
	fmt.Println(arr)
	// Output:
	// [1 2 3 4 5]
}

func Example_SelectionSort(){
	arr := []int{3,2,1,4,5}
	SelectionSort(arr)
	fmt.Println(arr)
	// Output:
	// [1 2 3 4 5]
}

func Example_InsertionSort(){
	arr := []int{3,2,1,4,5}
	InsertionSort(arr)
	fmt.Println(arr)
	// Output:
	// [1 2 3 4 5]
}
