package main

import (
	"fmt"
	"sort"
)

func test1() {
	a := []int{3, 9, 1, 6, 4, 2, 8, 2, 4, 5, 3, 0}
	sort.Ints(a) //对 []int进行排序
	fmt.Println(a)
	fmt.Println(sort.IntsAreSorted(a)) // true

	i := sort.SearchInts(a, 7) //搜索 a 中值为 x 的索引，如果找不到，则返回最接近且大于 x 的值的索引，  可能是 len(a)
	fmt.Println(i, a[i])       // 8
}

func test2() {
	a := sort.IntSlice{3, 7, 1, 3, 6, 9, 4, 1, 8, 5, 2, 0}
	a.Sort()
	fmt.Println(a)                // [0 1 1 2 3 3 4 5 6 7 8 9]
	fmt.Println(sort.IsSorted(a)) // true

	i := a.Search(6)
	fmt.Println(i, a[i]) // 8 6
}

func test3() {
	i := []int{3, 7, 1, 3, 6, 9, 4, 1, 8, 5, 2, 0}
	a := sort.IntSlice(i)
	sort.Sort(a)
	fmt.Println("排序结果:", a)
	fmt.Println("是否排序:", sort.IsSorted(a)) // true
}

func test4() {
	i := []int{3, 7, 1, 3, 6, 9, 4, 1, 8, 5, 2, 0}
	a := sort.IntSlice(i)
	c := sort.Reverse(a)                   //反序
	fmt.Println("是否排序:", sort.IsSorted(c)) // false
	fmt.Println("Reverse结果:", c)
	sort.Sort(c)
	fmt.Println("排序结果:", c)
	fmt.Println("是否排序:", sort.IsSorted(c)) // true
}

func main() {
	//test1()		//直接对数据进行排序
	//test2()
	//test3()
	test4() //反序
}
