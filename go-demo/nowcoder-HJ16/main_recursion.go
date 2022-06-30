package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unsafe"
)

var reader *bufio.Reader

func scan() (int, int, int) {
	if reader == nil {
		reader = bufio.NewReader(os.Stdin)
	}
	line, _, _ := reader.ReadLine()
	lineStr := *(*string)(unsafe.Pointer(&line))
	lineArr := strings.Split(lineStr, " ")
	a, _ := strconv.Atoi(lineArr[0])
	b, _ := strconv.Atoi(lineArr[1])
	c := 0
	if len(lineArr) > 2 {
		c, _ = strconv.Atoi(lineArr[2])
	}
	return a, b, c
}

func getMaxValue(allProducts [][2]int, primaryIds []int, primaryChildren [][]int, i, money, extraValue int) int {
	if money < 0 {
		return 0
	}
	if i < 0 || money == 0 {
		return extraValue
	}
	maxResult := getMaxValue(allProducts, primaryIds, primaryChildren, i-1, money, 0)
	primaryId := primaryIds[i]
	tempCost := allProducts[primaryId][0]
	tempValue := allProducts[primaryId][1]
	tempResult := getMaxValue(allProducts, primaryIds, primaryChildren, i-1, money-tempCost, tempValue)
	if maxResult < tempResult {
		maxResult = tempResult
	}
	if primaryChildren[primaryId] != nil {
		id1 := primaryChildren[primaryId][0]
		tempCost = allProducts[primaryId][0] + allProducts[id1][0]
		tempValue = allProducts[primaryId][1] + allProducts[id1][1]
		tempResult = getMaxValue(allProducts, primaryIds, primaryChildren, i-1, money-tempCost, tempValue)
		if maxResult < tempResult {
			maxResult = tempResult
		}
		if len(primaryChildren[primaryId]) > 1 {
			id2 := primaryChildren[primaryId][1]
			tempCost = allProducts[primaryId][0] + allProducts[id2][0]
			tempValue = allProducts[primaryId][1] + allProducts[id2][1]
			tempResult = getMaxValue(allProducts, primaryIds, primaryChildren, i-1, money-tempCost, tempValue)
			if maxResult < tempResult {
				maxResult = tempResult
			}
			tempCost = allProducts[primaryId][0] + allProducts[id1][0] + allProducts[id2][0]
			tempValue = allProducts[primaryId][1] + allProducts[id1][1] + allProducts[id2][1]
			tempResult = getMaxValue(allProducts, primaryIds, primaryChildren, i-1, money-tempCost, tempValue)
			if maxResult < tempResult {
				maxResult = tempResult
			}
		}
	}
	return maxResult + extraValue
}

func main() {
	myMoney, maxId, _ := scan()
	allProducts := make([][2]int, maxId+1)
	primaryIds := make([]int, 0)
	primaryChildren := make([][]int, maxId+1)
	for i := 1; i <= maxId; i++ {
		price, level, parentId := scan()
		if parentId == 0 {
			allProducts[i][0] = price
			allProducts[i][1] = price * level
			primaryIds = append(primaryIds, i)
		} else {
			allProducts[i][0] = price
			allProducts[i][1] = price * level
			primaryChildren[parentId] = append(primaryChildren[parentId], i)
		}
	}
	fmt.Println(getMaxValue(allProducts, primaryIds, primaryChildren, len(primaryIds)-1, myMoney, 0))
}
