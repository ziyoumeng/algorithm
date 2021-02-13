package sort

func BubbleSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for j := 1; j < len(arr); j++ { //要进行len(arr)-1次冒泡
		isSwap:= false
		for i := 0; i < len(arr)-1-j; i++ { // 冒泡范围
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSwap = true
			}
		}

		if !isSwap{ //是否已经达到满有序度
			break
		}
	}
}
