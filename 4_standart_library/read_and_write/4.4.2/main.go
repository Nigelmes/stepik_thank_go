package main

import (
	"bufio"
	"fmt"
	"os"
)

// начало решения

// readLines возвращает все строки из указанного файла
func readLines(name string) ([]string, error) {
	res := []string{}
	input, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res, scanner.Err()
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
