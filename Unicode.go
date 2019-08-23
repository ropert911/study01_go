package main

import (
	"fmt"
	"unicode"
)

func test1() {
	for _, r := range "Hello 世界！" {
		// 判断字符是否为汉字
		if unicode.Is(unicode.Scripts["Han"], r) {
			fmt.Printf("%c", r) // 世界
		}
	}
	fmt.Println("")
}

func test2() {
	for _, r := range "Hello ＡＢＣ！" {
		// 判断字符是否为大写
		if unicode.IsUpper(r) {
			fmt.Printf("%c", r) // HＡＢＣ
		}
	}
	fmt.Println("")

	for _, r := range "Hello ａｂｃ！" {
		// 判断字符是否为小写
		if unicode.IsLower(r) {
			fmt.Printf("%c", r) // elloａｂｃ
		}
	}
	fmt.Println("")

	for _, r := range "Hello ᾏᾟᾯ！" {
		// 判断字符是否为标题
		if unicode.IsTitle(r) {
			fmt.Printf("%c", r) // ᾏᾟᾯ
		}
	}
	fmt.Println("")
}

// 示例：输出 Unicode 规定的标题字符
func test3() {
	for _, cr := range unicode.Lt.R16 {
		for i := cr.Lo; i <= cr.Hi; i += cr.Stride {
			fmt.Printf("%c", i)
		}
	}
	fmt.Println("")
}

// 示例：转换大小写
func test4() {
	s := "Hello 世界！"

	for _, r := range s {
		fmt.Printf("%c", unicode.ToUpper(r))
	} // HELLO 世界！
	fmt.Println("")

	for _, r := range s {
		fmt.Printf("%c", unicode.ToLower(r))
	} // hello 世界！
	fmt.Println("")

	for _, r := range s {
		fmt.Printf("%c", unicode.ToTitle(r))
	} // HELLO 世界！
	fmt.Println("")

	for _, r := range s {
		fmt.Printf("%c", unicode.To(unicode.UpperCase, r))
	} // HELLO 世界！
	fmt.Println("")

	for _, r := range s {
		fmt.Printf("%c", unicode.To(unicode.LowerCase, r))
	} // hello 世界！
	fmt.Println("")

	for _, r := range s {
		fmt.Printf("%c", unicode.To(unicode.TitleCase, r))
	} // HELLO 世界！
}

// 示例
func test5() {
	s := "Hello 世界！"
	for _, r := range s {
		fmt.Printf("%c", unicode.SpecialCase(unicode.CaseRanges).ToUpper(r))
	} // HELLO 世界！
	fmt.Println("")

	for _, r := range s {
		fmt.Printf("%c", unicode.SpecialCase(unicode.CaseRanges).ToLower(r))
	} // hello 世界！
	fmt.Println("")

	for _, r := range s {
		fmt.Printf("%c", unicode.SpecialCase(unicode.CaseRanges).ToTitle(r))
	} // HELLO 世界！
	fmt.Println("")
}

// 示例：SimpleFold
func test6() {
	s := "ΦφϕkKK"
	// 看看 s 里面是什么
	for _, c := range s {
		fmt.Printf("%x  ", c)
	}
	fmt.Println()
	// 大写，小写，标题 | 当前字符 -> 下一个匹配字符
	for _, v := range s {
		fmt.Printf("%c, %c, %c | %c -> %c\n",
			unicode.ToUpper(v),
			unicode.ToLower(v),
			unicode.ToTitle(v),
			v,
			unicode.SimpleFold(v),
		)
	}
}

func test7() {
	fmt.Println() // 数字
	for _, r := range "Hello 123１２３一二三！" {
		if unicode.IsDigit(r) {
			fmt.Printf("%c", r)
		}
	} // 123１２３

	fmt.Println() // 数字
	for _, r := range "Hello 123１２３一二三！" {
		if unicode.IsNumber(r) {
			fmt.Printf("%c", r)
		}
	} // 123１２３

	fmt.Println() // 字母
	for _, r := range "Hello\n\t世界！" {
		if unicode.IsLetter(r) {
			fmt.Printf("%c", r)
		}
	} // Hello世界

	fmt.Println() // 空白
	for _, r := range "Hello \t世　界！\n" {
		if unicode.IsSpace(r) {
			fmt.Printf("%q", r)
		}
	} // ' ''\t''\u3000''\n'

	fmt.Println() // 控制字符
	for _, r := range "Hello\n\t世界！" {
		if unicode.IsControl(r) {
			fmt.Printf("%#q", r)
		}
	} // '\n''\t'

	fmt.Println() // 可打印
	for _, r := range "Hello　世界！\t" {
		if unicode.IsPrint(r) {
			fmt.Printf("%c", r)
		}
	} // Hello世界！

	fmt.Println() // 图形
	for _, r := range "Hello　世界！\t" {
		if unicode.IsGraphic(r) {
			fmt.Printf("%c", r)
		}
	} // Hello　世界！

	fmt.Println() // 掩码
	for _, r := range "Hello ៉៊់៌៍！" {
		if unicode.IsMark(r) {
			fmt.Printf("%c", r)
		}
	} // ៉៊់៌៍

	fmt.Println() // 标点
	for _, r := range "Hello 世界！" {
		if unicode.IsPunct(r) {
			fmt.Printf("%c", r)
		}
	} // ！

	fmt.Println() // 符号
	for _, r := range "Hello (<世=界>)" {
		if unicode.IsSymbol(r) {
			fmt.Printf("%c", r)
		}
	} // <=>
}

// 示例：判断汉字和标点
func test8() {
	// 将 set 设置为“汉字、标点符号”
	set := []*unicode.RangeTable{unicode.Han, unicode.P}
	for _, r := range "Hello 世界！" {
		if unicode.IsOneOf(set, r) {
			fmt.Printf("%c", r)
		}
	} // 世界！
}

// 示例：输出所有 mark 字符
func main() {
	fmt.Println("test1=========判断是不是汉字")
	test1()
	fmt.Println("test2=========判断大小写标题")
	test2()
	fmt.Println("test3=========输出 Unicode 规定的标题字符")
	test3()
	fmt.Println("test4=========转换大小写")
	test4()

	//for _, cr := range unicode.M.R16 {
	//	Lo, Hi, Stride := rune(cr.Lo), rune(cr.Hi), rune(cr.Stride)
	//	for i := Lo; i >= Lo && i <= Hi; i += Stride {
	//		if unicode.IsMark(i) {
	//			fmt.Printf("%c", i)
	//		}
	//	}
	//}
}
