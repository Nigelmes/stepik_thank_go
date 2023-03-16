package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	var abbr []rune
	phrase := readString()
	v := strings.Fields(phrase)
	for _, a := range v {
		r, _ := utf8.DecodeRuneInString(a)
		if unicode.IsLetter(r) {
			abbr = append(abbr, unicode.ToUpper(r))
		}

	}

	fmt.Println(string(abbr))
}

// readString читает строку из `os.Stdin` и возвращает ее
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
