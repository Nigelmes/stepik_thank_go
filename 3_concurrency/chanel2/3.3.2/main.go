package main

import (
	"fmt"
	"time"
)

// gather выполняет переданные функции одновременно
// и возвращает срез с результатами, когда они готовы
func gather(funcs []func() any) []any {
	lenfn := len(funcs)
	res := make([]any, lenfn)
	calcfn := make([]chan any, lenfn)
	for i := 0; i < lenfn; i++ {
		calcfn[i] = make(chan any)
		go func(fn func() any, c chan any) {
			c <- fn()
		}(funcs[i], calcfn[i])
	}
	for j := 0; j < lenfn; j++ {
		res[j] = <-calcfn[j]
	}
	return res
}

// squared возвращает функцию,
// которая считает квадрат n
func squared(n int) func() any {
	return func() any {
		time.Sleep(time.Duration(n) * 100 * time.Millisecond)
		return n * n
	}
}

func main() {
	funcs := []func() any{squared(2), squared(3), squared(4)}

	start := time.Now()
	nums := gather(funcs)
	elapsed := float64(time.Since(start)) / 1_000_000

	fmt.Println(nums)
	fmt.Printf("Took %.0f ms\n", elapsed)
}
