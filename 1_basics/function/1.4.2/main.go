package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

// shuffle перемешивает элементы nums in-place.
func shuffle(nums []int) {
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
}

// init - это специальная функция, которую Go вызывает до main()
// обычно используется для инициализации начального состояния приложения
func init() {
	rand.Seed(42)
}

func main() {
	nums := readInput()
	shuffle(nums)
	fmt.Println(nums)
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
