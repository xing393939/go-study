package sort

import "testing"

var sourceArr = []int{77, 198, 5, 8, 6, 52, 4, 3, 1, 63, 9, 44, 20}

// 冒泡排序
func TestBubbleSort(t *testing.T) {
	targetLen := len(sourceArr)
	targetArr := make([]int, targetLen)
	copy(targetArr, sourceArr)
	for i := 0; i < targetLen; i++ {
		for j := i + 1; j < targetLen; j++ {
			if targetArr[i] > targetArr[j] {
				temp := targetArr[i]
				targetArr[i] = targetArr[j]
				targetArr[j] = temp
			}
		}
	}
	t.Log(targetArr)
}

// 插入排序
func TestInsertionSort(t *testing.T) {
	targetLen := len(sourceArr)
	targetArr := make([]int, targetLen)
	copy(targetArr, sourceArr)
	for i := 1; i < targetLen; i++ {
		for j := i; j > 0; j-- {
			if targetArr[j] > targetArr[j-1] {
				break
			}
			temp := targetArr[j-1]
			targetArr[j-1] = targetArr[j]
			targetArr[j] = temp
		}
	}
	t.Log(targetArr)
}

func quickPartition(targetArr []int, p, r int) {
	if p >= r {
		return
	}
	i, j, pivot := 0, 0, targetArr[r]
	for ; j < r; j++ {
		if targetArr[j] < pivot {
			temp := targetArr[i]
			targetArr[i] = targetArr[j]
			targetArr[j] = temp
			i++
		}
	}
	temp := targetArr[i]
	targetArr[i] = targetArr[j]
	targetArr[j] = temp
	quickPartition(targetArr, p, i-1)
	quickPartition(targetArr, i+1, r)
}

// 快速排序
func TestQuickSort(t *testing.T) {
	targetLen := len(sourceArr)
	targetArr := make([]int, targetLen)
	copy(targetArr, sourceArr)
	quickPartition(targetArr, 0, targetLen-1)
	t.Log(targetArr)
}

func mergePartition(targetArr []int, p, r, q int) {
	temp := make([]int, 0, r-p+1)
	i, j := p, q+1
	for i <= q && j <= r {
		if targetArr[i] < targetArr[j] {
			temp = append(temp, targetArr[i])
			i++
		} else {
			temp = append(temp, targetArr[j])
			j++
		}
	}
	for ; i <= q; i++ {
		temp = append(temp, targetArr[i])
	}
	for ; j <= r; j++ {
		temp = append(temp, targetArr[j])
	}
	for k := range temp {
		targetArr[p+k] = temp[k]
	}
}

func mergeSort(targetArr []int, p, r int) {
	if p >= r {
		return
	}
	q := (r + p) / 2
	mergeSort(targetArr, p, q)
	mergeSort(targetArr, q+1, r)
	mergePartition(targetArr, p, r, q)
}

// 归并排序
func TestMergeSort(t *testing.T) {
	targetLen := len(sourceArr)
	targetArr := make([]int, targetLen)
	copy(targetArr, sourceArr)
	mergeSort(targetArr, 0, targetLen-1)
	t.Log(targetArr)
}
