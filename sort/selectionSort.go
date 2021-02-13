package sort

func SelectionSort(arr []int){
	if len(arr) <= 1{
		return
	}

	for i :=1 ;i < len(arr);i++{
		min:= i-1
		for j:=i;j<len(arr);j++{
			if arr[j] < arr[min]{
				min=j
			}
		}

		if min != i-1{
			arr[min],arr[i-1] = arr[i-1],arr[min]
		}
	}
}
