package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	correctstr := strings.Title(strings.ToLower(input))
	fmt.Println(correctstr)
}
