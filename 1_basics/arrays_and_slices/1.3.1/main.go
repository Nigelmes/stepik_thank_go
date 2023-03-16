package main

import (
	"fmt"
)

func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width)
	if len(text)-1 < width {
		fmt.Print(text)
		return
	}

	res := []rune(text[:width])
	// Возьмите первые `width` байт строки `text`, // допишите в конце `...` и сохраните результат
	// в переменную `res`  ...

	fmt.Println(string(res) + "...")
}
