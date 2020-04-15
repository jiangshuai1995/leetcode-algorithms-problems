package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(010))
}

/*  1. Two Sum  */
func twoSum(nums []int, target int) []int {
	v := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		dif := target - nums[i]
		c, ok := v[dif]
		if ok != false {
			return []int{c, i}
		}
		v[nums[i]] = i
	}
	return []int{-1, -1}
}

/* 2. Add Two Numbers */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	plus := 0
	array := []ListNode{}
	for l1 != nil || l2 != nil {
		a := 0
		b := 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		sum := a + b + plus
		if sum > 9 {
			value := sum - 10
			temp := ListNode{value, nil}
			array = append(array, temp)
			plus = 1
		} else {
			value := sum
			temp := ListNode{value, nil}
			array = append(array, temp)
			plus = 0
		}
	}
	if plus == 1 {
		temp := ListNode{plus, nil}
		array = append(array, temp)
	}
	for index := 0; index < len(array)-1; index++ {
		array[index].Next = &array[index+1]
	}
	return &(array[0])
}

/* 3. Longest Substring Without Repeating Characters */
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	total := 1
	for i := 0; i < len(s); i++ {
		result := make(map[byte]int)
		result[s[i]] = 1
		cal := 1
		for j := i + 1; j < len(s); j++ {
			_, exists := result[s[j]]
			if exists {
				break
			}
			result[s[j]] = j
			cal++
		}
		if cal > total {
			total = cal
		}
	}
	return total
}

/* 5. Longest Palindromic Substring */
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	longest := ""
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if isPalindromicString(s[i:j]) && ((j - i + 1) > len(longest)) {
				longest = s[i:j]
			}
		}
	}
	return longest
}
func isPalindromicString(s string) bool {
	if s == "" {
		return true
	}
	mid := len(s) / 2
	total := len(s) - 1
	for i := 0; i < mid; i++ {
		if s[i] != s[total-i] {
			return false
		}
	}
	return true
}

func reverse(x int) int {
	y := 0
	a := -(1 << 31)
	b := (1 << 31) - 1
	for x != 0 {
		y = y*10 + x%10
		x = x / 10
	}
	if y < a || y > b {
		return 0
	}
	return y
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	return isPalindromicString(s)
}

func romanToInt(s string) int {
	// 由于末位数字是I,X,C会导致越界，加一个字符Q
	s = s + "Q"
	sum := 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			if s[i+1] == 'X' {
				sum += 9
				i++
			} else if s[i+1] == 'V' {
				sum += 4
				i++
			} else {
				sum += 1
			}
			break
		case 'X':
			if s[i+1] == 'C' {
				sum += 90
				i++
			} else if s[i+1] == 'L' {
				sum += 40
				i++
			} else {
				sum += 10
			}
			break
		case 'C':
			if s[i+1] == 'M' {
				sum += 900
				i++
			} else if s[i+1] == 'D' {
				sum += 400
				i++
			} else {
				sum += 100
			}
			break
		case 'V':
			sum += 5
			break
		case 'L':
			sum += 50
			break
		case 'D':
			sum += 500
			break
		case 'M':
			sum += 1000
			break
		case 'Q':
			break
		default:
			break
		}
	}
	return sum
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	longest := ""
	s := strs[0]
	for i := 0; i < len(s); i++ {
		for _, value := range strs {
			if i >= len(value) {
				return longest
			}
			if s[i] != value[i] {
				return longest
			}
		}
		longest = s[:i+1]
	}
	return longest
}

/* Definition for singly-linked list. */
type ListNode struct {
	Val  int
	Next *ListNode
}
