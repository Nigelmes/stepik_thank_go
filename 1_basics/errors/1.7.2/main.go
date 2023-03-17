package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// account представляет счет
type account struct {
	balance   int
	overdraft int
}

func main() {
	var acc account
	var trans []int
	var erno error
	acc, trans, erno = parseInput()
	if erno != nil {
		fmt.Println(erno)
	} else {
		fmt.Println("->", acc, trans)
	}
}

// parseInput считывает счет и список транзакций из os.Stdin.
func parseInput() (account, []int, error) {
	accSrc, transSrc := readInput()
	acc, err := parseAccount(accSrc)
	if err != nil {
		return account{}, nil, err
	}
	trans, err := parseTransactions(transSrc)
	if err != nil {
		return account{}, nil, err
	}
	return acc, trans, nil
}

// readInput возвращает строку, которая описывает счет
// и срез строк, который описывает список транзакций.
// эту функцию можно не менять
func readInput() (string, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	accSrc := scanner.Text()
	var transSrc []string
	for scanner.Scan() {
		transSrc = append(transSrc, scanner.Text())
	}
	return accSrc, transSrc
}

// parseAccount парсит счет из строки
// в формате balance/overdraft.
func parseAccount(src string) (account, error) {
	parts := strings.Split(src, "/")
	balance, err := strconv.Atoi(parts[0])
	if err != nil {
		return account{0, 0}, err
	}
	overdraft, err := strconv.Atoi(parts[1])
	if err != nil {
		return account{0, 0}, err
	}
	if overdraft < 0 {
		return account{0, 0}, errors.New("expect overdraft >= 0")
	}
	if balance < -overdraft {
		return account{0, 0}, errors.New("balance cannot exceed overdraft")
	}
	return account{balance, overdraft}, nil
}

// parseTransactions парсит список транзакций из строки
// в формате [t1 t2 t3 ... tn].
func parseTransactions(src []string) ([]int, error) {
	trans := make([]int, len(src))
	for idx, s := range src {
		if t, err := strconv.Atoi(s); err != nil {
			return nil, err
		} else {
			trans[idx] = t
		}
	}
	return trans, nil
}
