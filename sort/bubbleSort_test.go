package sort

import "fmt"

func Example_BubbleSort(){
	arr := []int{3,2,1,4,5}
	BubbleSort(arr)
	fmt.Println(arr)
	// Output:
	// [1 2 3 4 5]
}
