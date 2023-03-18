package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// начало решения

// asLegacyDate преобразует время в легаси-дату
func asLegacyDate(t time.Time) string {
	usec := t.Unix()
	unano := t.UnixNano()
	usecstr := strconv.FormatInt(usec, 10)
	unanostr := strconv.FormatInt(unano%1e+9, 10)
	unanostr = strings.TrimRight(unanostr, "0")
	if len(unanostr) == 0 {
		return usecstr + ".0"
	}
	return usecstr + "." + unanostr
}

// parseLegacyDate преобразует легаси-дату во время.
// Возвращает ошибку, если легаси-дата некорректная.
func parseLegacyDate(d string) (time.Time, error) {
	str := strings.Split(d, ".")
	if len(str) < 2 || str[1] == "" {
		return time.Time{}, fmt.Errorf("incorrect input string")
	}
	p_second, err := strconv.Atoi(str[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("incorrect input string")
	}
	p := str[1]
	if len(str[1]) < 9 {
		for i := 0; i < 9-len(str[1]); i++ {
			p += "0"
		}
	}
	p_nanosecond, err := strconv.Atoi(p)
	if err != nil {
		return time.Time{}, fmt.Errorf("incorrect input string")
	}
	if p_second < 0 {
		return time.Time{}, fmt.Errorf("cannot be negative")
	}
	return time.Unix(int64(p_second), int64(p_nanosecond)), nil
}
