package main

import (
	"fmt"
	"sort"
)

func test1() {
	a := sort.StringSlice{"hello", "world", "golang", "sort", "nice"}
	a.Sort() // 二分法必须先排序
	fmt.Println(a)

	// 获取首字母大于 n 的元素中最小的
	i := sort.Search(len(a), func(i int) bool {
		return len(a[i]) > 0 && a[i][0] > 'n'
	})
	// 显示找到的元素
	fmt.Println(a[i]) // sort
}

func main() {
	test1()
}
