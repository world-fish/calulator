package main

import "fmt"

// 冒泡排序
func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				//arr[j], arr[j+1] = arr[j+1], arr[j]
				arr[j] ^= arr[j+1]
				arr[j+1] ^= arr[j]
				arr[j] ^= arr[j+1]
			}
		}
	}
}

// 插入排序
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
}

// 选择排序
func selectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// 快速排序
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := arr[0]
	var less, greater []int
	for _, num := range arr[1:] {
		if num <= mid {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	return append(append(quickSort(less), mid), quickSort(greater)...)
}

// 归并排序
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

// 二分
func erfen(nums []int, num int) int {
	left := 0
	right := len(nums)
	var mid int
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] < num {
			left = mid + 1
		} else if nums[mid] > num {
			right = mid - 1
		} else if nums[mid] == num {
			return mid
		}
	}
	return -1
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				result = append(result, left[0])
				left = left[1:]
			} else {
				result = append(result, right[0])
				right = right[1:]
			}
		} else if len(left) > 0 {
			result = append(result, left...)
		} else if len(right) > 0 {
			result = append(result, right...)
		}
	}
	return result
}

func main() {
	arr := []int{67, 12, 93, 5, 76, 41, 88, 72, 29, 34, 19, 95, 61, 50, 7, 83, 30, 91, 68, 53}
	arrr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(erfen(arrr, 20))

	// 冒泡排序
	bubbleSortArr := make([]int, len(arr))
	copy(bubbleSortArr, arr)
	bubbleSort(bubbleSortArr)
	fmt.Println("冒泡排序结果:", bubbleSortArr)

	// 插入排序
	insertionSortArr := make([]int, len(arr))
	copy(insertionSortArr, arr)
	insertionSort(insertionSortArr)
	fmt.Println("插入排序结果:", insertionSortArr)

	// 选择排序
	selectionSortArr := make([]int, len(arr))
	copy(selectionSortArr, arr)
	selectionSort(selectionSortArr)
	fmt.Println("选择排序结果:", selectionSortArr)

	// 快速排序
	quickSortArr := make([]int, len(arr))
	copy(quickSortArr, arr)
	quickSortArr = quickSort(quickSortArr)
	fmt.Println("快速排序结果:", quickSortArr)

	// 归并排序
	mergeSortArr := make([]int, len(arr))
	copy(mergeSortArr, arr)
	mergeSortArr = mergeSort(mergeSortArr)
	fmt.Println("归并排序结果:", mergeSortArr)
}
