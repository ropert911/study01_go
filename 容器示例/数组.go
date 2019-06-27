package 容器示例

import "fmt"

func test5() {
	var iArray1 [5]int32
	var iArray2 = [5]int32{1, 2, 3, 4, 5}
	iArray3 := [5]int32{1, 2, 3, 4, 5}
	iArray4 := [5]int32{6, 7, 8, 9, 10}
	iArray5 := [...]int32{11, 12, 13, 14, 15}
	iArray6 := [4][4]int32{{1}, {1, 2}, {1, 2, 3}}
	fmt.Println(iArray1)
	fmt.Println(iArray2)
	fmt.Println(iArray3)
	fmt.Println(iArray4)
	fmt.Println(iArray5)
	fmt.Println(iArray6)
}

func test6() {
	iArray4 := [5]int32{6, 7, 8, 9, 10}
	fmt.Println(len(iArray4))
	fmt.Println(cap(iArray4))
	for i := range iArray4 {
		fmt.Println(iArray4[i])
	}
}

//切片与隐藏数组  切片是长度可变、容量固定的相同的元素序列。Go语言的切片本质是一个数组
func test9() {
	slice0 := []string{"a", "b", "c", "d", "e"}
	slice1 := slice0[2 : len(slice0)-1]
	slice2 := slice0[:3]
	fmt.Println(slice0, slice1, slice2)
	slice2[2] = "8"
	fmt.Println(slice0, slice1, slice2)
}

func test10() {
	slice0 := []string{"a", "b", "c", "d", "e"}
	fmt.Println("\n~~~~~~元素遍历~~~~~~")
	for _, ele := range slice0 { //第一个是索引号
		fmt.Print(ele, " ")
		ele = "7"
	}
	fmt.Println("\n~~~~~~索引遍历~~~~~~")
	for index := range slice0 {
		fmt.Print(slice0[index], " ")
	}
	fmt.Println("\n~~~~~~元素索引共同使用~~~~~~")
	for index, ele := range slice0 {
		fmt.Print(ele, slice0[index], " ")
	}
	fmt.Println("\n~~~~~~修改~~~~~~")
	for index := range slice0 {
		slice0[index] = "9"
	}
	fmt.Println(slice0)
}

//追加、复制切片
func test11() {
	slice := []int32{}
	fmt.Printf("slice的长度为：%d,slice为：%v\n", len(slice), slice)
	slice = append(slice, 12, 11, 10, 9)
	fmt.Printf("追加后，slice的长度为：%d,slice为：%v\n", len(slice), slice)
	slicecp := make([]int32, (len(slice)))
	fmt.Printf("slicecp的长度为：%d,slicecp为：%v\n", len(slicecp), slicecp)
	copy(slicecp, slice)
	fmt.Printf("复制赋值后，slicecp的长度为：%d,slicecp为：%v\n", len(slicecp), slicecp)
}

//内置函数append  内置函数append可以向一个切片后追加一个或多个同类型的其他值。如果追加的元素数量超过了原切片容量，那么最后返回的是一个全新数组中的全新切片。如果没有超过，那么最后返回的是原数组中的全新切片。无论如何，append对原切片无任何影响
func test12() {
	slice := []int32{1, 2, 3, 4, 5, 6}
	slice2 := slice[:2]
	slice3 := append(slice2, 50, 60, 70, 80, 90)
	fmt.Printf("slice为：%v\n", slice)
	fmt.Printf("操作的切片：%v\n", slice2)
	fmt.Printf("操作后新切片：%v\n", slice3)

	_ = append(slice2, 50, 60)
	fmt.Printf("slice为：%v\n", slice)
	fmt.Printf("操作的切片：%v\n", slice2)
}

func main() {
	test12()
}
