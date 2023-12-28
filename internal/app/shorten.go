package app

import (
	"strings"
)

var (
	n           uint32 = 1
	alphabet    string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	alphabetLen        = uint32(len(alphabet))
)

func Short(str string) string {
	var (
		nums    []uint32
		num     = n
		builder strings.Builder
	)

	for num > 0 {
		nums = append(nums, num%alphabetLen)
		num /= alphabetLen
	}

	Reverse(nums)

	for _, num := range nums {
		builder.WriteString(string(alphabet[num]))
	}
	n++
	return builder.String()
}

func Reverse(s []uint32) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
