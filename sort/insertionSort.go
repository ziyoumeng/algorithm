package sort

func InsertionSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	for i := 1; i < len(arr); i++ {
		val := arr[i] //待插入元素
		j := i - 1
		for ; j >= 0; j-- { //有序部分
			if arr[j] > val {
				arr[j+1] = arr[j] //不满足则向右移动一位，空出下一个插入点
			} else {
				//这里是易错点，找到插入点后停止,不能再这里插入。因为容易遗漏了当val为最小值的情况
				break
			}
		}
		arr[j+1] = val //插入
	}
}
