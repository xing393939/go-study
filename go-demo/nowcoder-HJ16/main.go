package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unsafe"
)

var reader2 *bufio.Reader

func scan2() (int, int, int) {
	if reader2 == nil {
		reader2 = bufio.NewReader(os.Stdin)
	}
	line, _, _ := reader2.ReadLine()
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

func getMaxResult(dp [][]int, i, j, tempCost, tempValue, maxResult int) int {
	if j >= tempCost {
		tempResult := dp[i-1][j-tempCost] + tempValue
		if tempResult > maxResult {
			return tempResult
		}
	}
	return maxResult
}

func main() {
	myMoney, maxId, _ := scan2()
	allProducts := make([][2]int, maxId+1)
	primaryIds := make([]int, 0)
	primaryChildren := make([][]int, maxId+1)
	for i := 1; i <= maxId; i++ {
		price, level, parentId := scan2()
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
	primaryLen := len(primaryIds)
	dp := make([][]int, primaryLen+1)
	for i := 0; i <= primaryLen; i++ {
		dp[i] = make([]int, myMoney+1)
		if i == 0 {
			continue
		}
		for j := 0; j <= myMoney; j++ {
			maxResult := dp[i-1][j]
			primaryId := primaryIds[i-1]
			tempCost := allProducts[primaryId][0]
			tempValue := allProducts[primaryId][1]
			maxResult = getMaxResult(dp, i, j, tempCost, tempValue, maxResult)
			if primaryChildren[primaryId] != nil {
				id1 := primaryChildren[primaryId][0]
				tempCost = allProducts[primaryId][0] + allProducts[id1][0]
				tempValue = allProducts[primaryId][1] + allProducts[id1][1]
				maxResult = getMaxResult(dp, i, j, tempCost, tempValue, maxResult)
				if len(primaryChildren[primaryId]) > 1 {
					id2 := primaryChildren[primaryId][1]
					tempCost = allProducts[primaryId][0] + allProducts[id2][0]
					tempValue = allProducts[primaryId][1] + allProducts[id2][1]
					maxResult = getMaxResult(dp, i, j, tempCost, tempValue, maxResult)
					tempCost = allProducts[primaryId][0] + allProducts[id1][0] + allProducts[id2][0]
					tempValue = allProducts[primaryId][1] + allProducts[id1][1] + allProducts[id2][1]
					maxResult = getMaxResult(dp, i, j, tempCost, tempValue, maxResult)
				}
			}
			dp[i][j] = maxResult
		}
	}
	fmt.Println(dp[primaryLen][myMoney])
}
