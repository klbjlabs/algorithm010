// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

// 示例 1:

// 输入: s = "anagram", t = "nagaram"
// 输出: true
// 示例 2:

// 输入: s = "rat", t = "car"
// 输出: false
// 说明:
// 你可以假设字符串只包含小写字母。

// 进阶:
// 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/valid-anagram
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "anagram"
	t := "nagaram"

	fmt.Println(isAnagram_3(s, t))
}

// TIME:  O(n)
// SPACE: O(n)
func isAnagram(s, t string) bool {
	m1 := map[string]int{}
	m2 := map[string]int{}

	for _, c := range s {
		m1[string(c)]++
	}
	for _, c := range s {
		m2[string(c)]++
	}

	return reflect.DeepEqual(m1, m2)
}

func isAnagram_1(s, t string) bool {
	var sA, tA [26]int

	for _, c := range s {
		sA[c-'a']++
	}
	for _, c := range s {
		tA[c-'a']++
	}

	return sA == tA
}

func isAnagram_2(s, t string) bool {
	var a [26]int

	for _, c := range s {
		a[c-'a']++
	}
	for _, c := range s {
		a[c-'a']--
	}
	for _, c := range a {
		if c != 0 {
			return false
		}
	}
	return true
}

func isAnagram_3(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var a [26]int
	for i := 0; i < len(s); i++ {
		a[s[i]-'a']++
		a[t[i]-'a']--
	}
	for _, v := range a {
		if v != 0 {
			return false
		}
	}
	return true
}
