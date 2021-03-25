package main

import "fmt"

func task1(arr []int, sum int) int {
	for len(arr) != 0 {
		sum += arr[0]
		return task1(arr[1:], sum)
	}

	return sum
}

func task2(arr []string, m map[byte]int) map[byte]int {
	for len(arr) != 0 {
		if count, ok := m[arr[0][0]]; !ok {
			m[arr[0][0]] = len(arr[0])
		} else {
			if count < len(arr[0]) {
				m[arr[0][0]] = len(arr[0])
			}
		}

		return task2(arr[1:], m)
	}

	return m
}

func main() {
	var (
		sum int          = 0
		m   map[byte]int = make(map[byte]int)
	)

	println(task1([]int{1, 2, 3, 4, 5}, sum))
	m = task2([]string{"aa", "aaaaaaa", "b", "b", "bbbbbb", "cccc", "c", "ddd", "d"}, m)

	for k, v := range m {
		fmt.Printf("%c: %d\n", k, v)
	}
}
