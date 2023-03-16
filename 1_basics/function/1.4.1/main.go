package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func filter(predicate func(int) bool, iterable []int) []int {
	result := make([]int, 0, 10)
	for z, a := range iterable {
		if predicate(z) {
			result = append(result, a)
		}
	}
	return result
}

func main() {
	src := readInput()
	res := filter(func(i int) bool { return (src[i] % 2) == 0 }, src)
	fmt.Println(res)
}

// readInput считывает целые числа из `os.Stdin`
// и возвращает в виде среза
// разделителем чисел считается пробел
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}
