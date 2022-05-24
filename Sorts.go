package main

import (
	"fmt"
	"math/rand"
)

/*
排序算法大合集
*/

//归并排序暂存的数字
var t [20]int

func main() {
	//声明两个数组
	num1 := []int{5, 1, 1, 2, 0, 0}
	num2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	//对数组1进行排序
	//mergeSort(num1, 0, len(num1)-1)
	HeapSort(num1)
	fmt.Println(num1)

	//对数组2进行排序
	//mergeSort(num2, 0, len(num2)-1)
	HeapSort(num2)
	fmt.Println(num2)
}

// HeapSort 堆排序
func HeapSort(nums []int) {
	// 1、构建堆(这里用大顶堆构建升序)
	// 2、调整堆，把堆顶元素和第i-1个元素交换，这样0....i-2就又成为一个堆，继续对这个堆进行构建，调整
	Hepify(nums, len(nums)) // 先构建n个元素的大顶堆
	for i := len(nums) - 1; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i] // 调整堆顶元素，把堆顶元素和最后一个元素交换
		Hepify(nums, i)
	}
}

// Hepify 构建堆，一般从最后一个非叶子节点开始构建，即从下往上调整，从下往上能让最大（小）值元素转移到堆顶
func Hepify(nums []int, unsortCapacity int) {
	for i := (unsortCapacity / 2) - 1; i >= 0; i-- { // 非叶子节点的i范围从0...(n/2-1)个
		// 调整左子树
		leftIndex := 2*i + 1
		if leftIndex < unsortCapacity && nums[i] < nums[leftIndex] {
			swap(nums, leftIndex, i)
		}
		// 调整右子树
		rightIndex := 2*i + 2
		if rightIndex < unsortCapacity && nums[i] < nums[rightIndex] {
			swap(nums, rightIndex, i)
		}
	}
}

//归并排序
func mergeSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	mergeSort(nums, left, mid)
	mergeSort(nums, mid+1, right)
	i, j, c := left, mid+1, 0
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			t[c] = nums[i]
			c++
			i++
		} else {
			t[c] = nums[j]
			c++
			j++
		}
	}
	for i <= mid {
		t[c] = nums[i]
		c++
		i++
	}
	for j <= right {
		t[c] = nums[j]
		c++
		j++
	}
	for k := 0; k < right-left+1; k++ {
		nums[k+left] = t[k]
	}

}

//快排
func quickSort(nums []int, left, right int) {
	if left < right {

		//q为排好序的位置，对该位置的左右两边进行排序
		q := randSort(nums, left, right)
		quickSort(nums, left, q-1)
		quickSort(nums, q+1, right)
	}
}

//选取一个随机数作为index进行排序，放置时间复杂度恶化至O(n²)
func randSort(nums []int, left, right int) int {
	q := rand.Intn(right-left+1) + left
	swap(nums, q, left)

	return partition(nums, left, right)
}

//选取index，保证左边的数都比他小，右边的数都比他大
func partition(nums []int, left int, right int) int {
	q := left
	for left < right {
		for left < right && nums[right] > nums[q] {
			right--
		}
		for left < right && nums[left] <= nums[q] {
			left++
		}

		if left < right {
			swap(nums, left, right)
		}
	}
	swap(nums, q, left)
	return left
}

//交换数组两个位置上的值
func swap(nums []int, left, right int) {
	t := nums[left]
	nums[left] = nums[right]
	nums[right] = t
}
