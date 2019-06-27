package 容器示例

import "fmt"

//map 初始化
func test1() {
	map1 := make(map[string]string, 5)
	map2 := make(map[string]string)
	map3 := map[string]string{}
	map4 := map[string]string{"a": "1", "b": "2", "c": "3"}
	fmt.Println(map1, map2, map3, map4)
}

//map 遍历
func test2() {
	map1 := make(map[string]string)
	map1["a"] = "1"
	map1["b"] = "2"
	map1["c"] = "3"
	for key, value := range map1 {
		fmt.Printf("%s->%-10s", key, value)
	}
}

//查找、修改和删除
func test3() {
	map4 := map[string]string{"a": "1", "b": "2", "c": "3"}
	val, exist := map4["a"]
	val2, exist2 := map4["d"]
	fmt.Printf("%v,%v\n", exist, val)
	fmt.Printf("%v,%v\n", exist2, val2)

	map4["a"] = "8" //修改映射和添加映射没什么区别
	fmt.Printf("%v\n", map4)

	fmt.Println("删除b：")
	delete(map4, "b")
	fmt.Printf("%v", map4)
}

func main() {
	test1()
	test2()
	test3()
}
