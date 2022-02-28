package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n string
	var m string

	fmt.Scan(&n, &m)
	list := strings.Split(n, ",")
	var arr []int
	for _, v := range list {
		value, _ := strconv.Atoi(v)
		arr = append(arr, value)
	}

	mi, _ := strconv.Atoi(m)
	ans := subsets(arr, mi)
	fmt.Println(ans)
}

func subsets(nums []int, m int) (ans [][]int) {
	//回溯
	var set []int
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)

	//过滤m长度不够的
	var newAns = make([][]int, 0, 100)
	for _, v := range ans {
		if len(v) >= m {
			newAns = append(newAns, v)
		}
	}
	ans = newAns

	//重新排序
	length := len(ans)
	var j int
	for i := 0; i < length-1; i++ {
		for j = i + 1; j < length; j++ {
			if len(ans[i]) > len(ans[j]) {
				ans[i], ans[j] = ans[j], ans[i]
			} else if len(ans[i]) == len(ans[j]) {
				for k := 0; k < len(ans[i]); k++ {
					if ans[i][k] > ans[j][k] {
						ans[i], ans[j] = ans[j], ans[i]
						break
					}
				}
			}
		}
	}

	return
}
