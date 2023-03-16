package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

// result представляет результат матча
type result byte

// возможные результаты матча
const (
	win  result = 'W'
	draw result = 'D'
	loss result = 'L'
)

// team представляет команду
type team byte

type match struct {
	first, second team
	result        result
}

// rating представляет турнирный рейтинг команд -
// количество набранных очков по каждой команде
type rating map[team]int

// tournament представляет турнир
type tournament []match

// calcRating считает и возвращает рейтинг турнира
func (trn *tournament) calcRating() rating {
	rating1 := make(rating)
	for _, onematching := range *trn {
		if onematching.result == win {
			rating1[onematching.first] += 3
			rating1[onematching.second] += 0
		} else if onematching.result == loss {
			rating1[onematching.second] += 3
			rating1[onematching.first] += 0
		} else if onematching.result == draw {
			rating1[onematching.first] += 1
			rating1[onematching.second] += 1
		}

	}
	return rating1

}

func main() {
	src := readString()
	trn := parseTournament(src)
	rt := trn.calcRating()
	rt.print()
}

// readString считывает и возвращает строку из os.Stdin
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, err := rdr.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	return str
}

// parseTournament парсит турнир из исходной строки
func parseTournament(s string) tournament {
	descriptions := strings.Split(s, " ")
	trn := tournament{}
	for _, descr := range descriptions {
		m := parseMatch(descr)
		trn.addMatch(m)
	}
	return trn
}

// parseMatch парсит матч из фрагмента исходной строки
func parseMatch(s string) match {
	return match{
		first:  team(s[0]),
		second: team(s[1]),
		result: result(s[2]),
	}
}

// addMatch добавляет матч к турниру
func (t *tournament) addMatch(m match) {
	*t = append(*t, m)
}

// print печатает результаты турнира
// в формате Aw Bx Cy Dz
func (r *rating) print() {
	var parts []string
	for team, score := range *r {
		part := fmt.Sprintf("%c%d", team, score)
		parts = append(parts, part)
	}
	sort.Strings(parts)
	fmt.Println(strings.Join(parts, " "))
}
