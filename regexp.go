package main

import (
	"fmt"
	"regexp"
)

// 示例：MatchString、QuoteMeta
func regexpTest1() {
	pat := `(((abc.)def.)ghi)`
	src := `abc-def-ghi abc+def+ghi`

	fmt.Println(regexp.MatchString(pat, src))
	// true <nil>

	fmt.Println(regexp.QuoteMeta(pat))
	// \(\(\(abc\.\)def\.\)ghi\)
}

// 示例：第一匹配和最长匹配
func regexpTest2() {
	b := []byte("abc1def1")
	pat := `abc1|abc1def1`
	reg1 := regexp.MustCompile(pat)      // 第一匹配
	reg2 := regexp.MustCompilePOSIX(pat) // 最长匹配
	fmt.Printf("%s\n", reg1.Find(b))     // abc1
	fmt.Printf("%s\n", reg2.Find(b))     // abc1def1

	b = []byte("abc1def1")
	pat = `(abc|abc1def)*1`
	reg1 = regexp.MustCompile(pat)      // 第一匹配
	reg2 = regexp.MustCompilePOSIX(pat) // 最长匹配
	fmt.Printf("%s\n", reg1.Find(b))    // abc1
	fmt.Printf("%s\n", reg2.Find(b))    // abc1def1
}

// 示例：正则信息
func regexpTest3() {
	pat := `(abc)(def)(ghi)`
	reg := regexp.MustCompile(pat)

	// 获取正则表达式字符串
	fmt.Println(reg.String()) // (abc)(def)(ghi)

	// 获取分组数量
	fmt.Println(reg.NumSubexp()) // 3

	fmt.Println()

	// 获取分组名称
	pat = `(?P<Name1>abc)(def)(?P<Name3>ghi)`
	reg = regexp.MustCompile(pat)

	for i := 0; i <= reg.NumSubexp(); i++ {
		fmt.Printf("%d: %q\n", i, reg.SubexpNames()[i])
	}
	// 0: ""
	// 1: "Name1"
	// 2: ""
	// 3: "Name3"

	fmt.Println()

	// 获取字面前缀
	pat = `(abc1)(abc2)(abc3)`
	reg = regexp.MustCompile(pat)
	fmt.Println(reg.LiteralPrefix()) // abc1abc2abc3 true

	pat = `(abc1)|(abc2)|(abc3)`
	reg = regexp.MustCompile(pat)
	fmt.Println(reg.LiteralPrefix()) //  false

	pat = `abc1|abc2|abc3`
	reg = regexp.MustCompile(pat)
	fmt.Println(reg.LiteralPrefix()) // abc false
}

// 示例：Find、FindSubmatch
func regexpTest4() {
	pat := `(((abc.)def.)ghi)`
	reg := regexp.MustCompile(pat)

	src := []byte(`abc-def-ghi abc+def+ghi`)

	// 查找第一个匹配结果
	fmt.Printf("%s\n", reg.Find(src)) // abc-def-ghi

	fmt.Println()

	// 查找第一个匹配结果及其分组字符串
	first := reg.FindSubmatch(src)
	for i := 0; i < len(first); i++ {
		fmt.Printf("%d: %s\n", i, first[i])
	}
	// 0: abc-def-ghi
	// 1: abc-def-ghi
	// 2: abc-def-
	// 3: abc-
}

// 示例：FindIndex、FindSubmatchIndex
func regexpTest5() {
	pat := `(((abc.)def.)ghi)`
	reg := regexp.MustCompile(pat)

	src := []byte(`abc-def-ghi abc+def+ghi`)

	// 查找第一个匹配结果
	matched := reg.FindIndex(src)
	fmt.Printf("%v\n", matched) // [0 11]
	m := matched[0]
	n := matched[1]
	fmt.Printf("%s\n\n", src[m:n]) // abc-def-ghi

	// 查找第一个匹配结果及其分组字符串
	matched = reg.FindSubmatchIndex(src)
	fmt.Printf("%v\n", matched) // [0 11 0 11 0 8 0 4]
	for i := 0; i < len(matched)/2; i++ {
		m := matched[i*2]
		n := matched[i*2+1]
		fmt.Printf("%s\n", src[m:n])
	}
	// abc-def-ghi
	// abc-def-ghi
	// abc-def-
	// abc-
}

// 示例：FindAll、FindAllSubmatch
func regexpTest6() {
	pat := `(((abc.)def.)ghi)`
	reg := regexp.MustCompile(pat)

	s := []byte(`abc-def-ghi abc+def+ghi`)

	// 查找所有匹配结果
	for _, one := range reg.FindAll(s, -1) {
		fmt.Printf("%s\n", one)
	}
	// abc-def-ghi
	// abc+def+ghi

	// 查找所有匹配结果及其分组字符串
	all := reg.FindAllSubmatch(s, -1)
	for i := 0; i < len(all); i++ {
		fmt.Println()
		one := all[i]
		for i := 0; i < len(one); i++ {
			fmt.Printf("%d: %s\n", i, one[i])
		}
	}
	// 0: abc-def-ghi
	// 1: abc-def-ghi
	// 2: abc-def-
	// 3: abc-

	// 0: abc+def+ghi
	// 1: abc+def+ghi
	// 2: abc+def+
	// 3: abc+
}

// 示例：Expand
func main() {
	pat := `(((abc.)def.)ghi)`
	reg := regexp.MustCompile(pat)

	src := []byte(`abc-def-ghi abc+def+ghi`)
	template := []byte(`$0   $1   $2   $3`)

	// 替换第一次匹配结果
	match := reg.FindSubmatchIndex(src)
	fmt.Printf("%v\n", match) // [0 11 0 11 0 8 0 4]
	dst := reg.Expand(nil, template, src, match)
	fmt.Printf("%s\n\n", dst)
	// abc-def-ghi   abc-def-ghi   abc-def-   abc-

	// 替换所有匹配结果
	for _, match := range reg.FindAllSubmatchIndex(src, -1) {
		fmt.Printf("%v\n", match)
		dst := reg.Expand(nil, template, src, match)
		fmt.Printf("%s\n", dst)
	}
	// [0 11 0 11 0 8 0 4]
	// abc-def-ghi   abc-def-ghi   abc-def-   abc-
	// [12 23 12 23 12 20 12 16]
	// abc+def+ghi   abc+def+ghi   abc+def+   abc+
}
