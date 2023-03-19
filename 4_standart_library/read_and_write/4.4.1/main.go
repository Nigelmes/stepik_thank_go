package main

import (
	"fmt"
	"os"
)

// начало решения

// readLines возвращает все строки из указанного файла
func readLines(name string) ([]string, error) {
	res := []string{}
	text, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	var temp string
	for i := 0; i < len(text); i++ {
		if text[i] == '\n' {
			res = append(res, temp)
			temp = ""
			continue
		}
		temp += string(text[i])
	}

	return res, nil
}

// конец решения

func main() {
	lines, err := readLines("./test.txt")
	if err != nil {
		panic(err)
	}
	for idx, line := range lines {
		fmt.Printf("%d: %s\n", idx+1, line)
	}
}
