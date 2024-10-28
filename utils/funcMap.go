package utils

import "time"

func DateFormat(date time.Time, layout string) string {
	return date.Format(layout)
}

func Listtag(tagstr string) error {
	return nil
}

func Truncate(s string, n int) string {
	runes := []rune(s)
	if len(runes) > n {
		return string(runes[:n])
	}
	return s
}

func Minus(a1, a2 int) int {
	return a1 - a2
}

func Add(a1, a2 int) int {
	return a1 + a2
}

// 判断数字是否是奇数
func IsOdd(number int) bool {
	return !IsEven(number)
}

// 判断数字是否是偶数
func IsEven(number int) bool {
	return number%2 == 0
}