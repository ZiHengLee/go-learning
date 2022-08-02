package dp

import (
	"fmt"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	fmt.Println(longestCommonSubsequence("abc","bc"))
}

func TestMaxSubArray(t *testing.T) {
	a := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(a))
}

func TestMinDistance(t *testing.T) {
	a ,b := "horse", "ros"
	fmt.Println(minDistance(a,b))
}

func TestLongestValidParentheses(t *testing.T) {
	a := "(()())"
	fmt.Println(longestValidParentheses(a))
}