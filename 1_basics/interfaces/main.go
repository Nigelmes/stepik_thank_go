package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// element - интерфейс элемента последовательности
type element interface{}

// weightFunc - функция, которая возвращает вес элемента
type weightFunc func(element) int

// iterator - интерфейс, который умеет
// поэлементно перебирать последовательность
type iterator interface {
	next() bool
	val() element
}

// intIterator - итератор по целым числам
// (реализует интерфейс iterator)
type intIterator struct {
	next1 *intIterator
	val1  element
}

func (c *intIterator) next() bool {
	if c != nil && c.next1 != nil {
		*c = *c.next1
		return true
	}
	return false
}

func (c *intIterator) val() element {
	return c.val1
}

// методы intIterator, которые реализуют интерфейс iterator
// конструктор intIterator

func newIntIterator(src []int) *intIterator {
	var masstruct *intIterator
	first := masstruct
	for _, a := range src {
		var tmp intIterator
		tmp.next1 = nil
		tmp.val1 = a
		if masstruct == nil {
			masstruct = &tmp
			first = masstruct
		} else {
			masstruct.next1 = &tmp
			masstruct = masstruct.next1
		}
	}
	return first
}

// main находит максимальное число из переданных на вход программы.
func main() {
	nums := readInput()
	it := newIntIterator(nums)
	weight := func(el element) int {
		return el.(int)
	}
	m := max(it, weight)
	fmt.Println(m)
}

// max возвращает максимальный элемент в последовательности.
// Для сравнения элементов используется вес, который возвращает
// функция weight.
func max(it iterator, weight weightFunc) element {
	var maxEl = it.val()
	for it.next() {
		curr := it.val()
		if maxEl == nil || weight(curr) > weight(maxEl) {
			maxEl = curr
		}
	}
	return maxEl
}

// readInput считывает последовательность целых чисел из os.Stdin.
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	return nums
}
