package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
)

// начало решения

// Task описывает задачу, выполненную в определенный день
type Task struct {
	Date  time.Time
	Dur   time.Duration
	Title string
}

// ParsePage разбирает страницу журнала
// и возвращает задачи, выполненные за день
func ParsePage(src string) ([]Task, error) {
	lines := strings.Split(src, "\n")
	date, err := parseDate(lines[0])
	if err != nil {
		return nil, err
	}
	tasks, err := parseTasks(date, lines[1:])
	if err != nil {
		return nil, err
	}
	sortTasks(tasks)
	return tasks, nil
}

// parseDate разбирает дату в формате дд.мм.гггг
func parseDate(src string) (time.Time, error) {
	t, err := time.Parse("02.01.2006", src)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

var timez = "15:04"

// parseTasks разбирает задачи из записей журнала
func parseTasks(date time.Time, lines []string) ([]Task, error) {
	res := make([]Task, 0)
	titlemap := make(map[string]int)
	re := regexp.MustCompile(`(\d+:\d+) - (\d+:\d+) (.+)`)

	for i, line := range lines {
		if re.MatchString(string(line)) {
			fz := re.FindAllString(line, -1)
			find := strings.Fields(fz[0])
			diff, err := difftime(timez, find[0], find[2])
			if err != nil {
				return nil, err
			}
			ttitle := strings.Join(find[3:], " ")
			if zz, ok := titlemap[ttitle]; ok {
				res[zz].Dur += diff
			} else {
				titlemap[ttitle] = i
				res = append(res, Task{date, diff, ttitle})
			}
		} else {
			return nil, errors.New("Error input page")
		}
	}
	for _, j := range res {
		if j.Dur == 0 {
			return nil, errors.New("Error input page")
		}
	}

	return res, nil
}

func difftime(mask, time1, time2 string) (time.Duration, error) {
	before, err := time.Parse(timez, time1)
	if err != nil {
		return 0, err
	}
	after, err := time.Parse(timez, time2)
	if err != nil {
		return 0, err
	}
	diff := after.Sub(before)
	if diff < 0 {
		return 0, errors.New("incorrect time")
	}
	return diff, nil
}

// sortTasks упорядочивает задачи по убыванию длительности
func sortTasks(tasks []Task) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Dur > tasks[j].Dur
	})
}

// конец решения

func main() {
	page := `15.04.2022
	10:00 - 10:00 мурик
	8:00 - 8:30 Завтрак
	8:30 - 9:30 Оглаживание кота
	9:30 - 10:00 Интернеты
	10:00 - 14:00 Напряженная работа
	14:00 - 14:45 Обед
	14:45 - 15:00 Оглаживание кота
	15:00 - 19:00 Напряженная работа
	19:00 - 19:30 Интернеты
	19:30 - 22:30 Безудержное веселье
	22:30 - 23:00 Оглаживание кота`

	entries, err := ParsePage(page)
	if err != nil {
		panic(err)
	}
	fmt.Println("Мои достижения за", entries[0].Date.Format("2006-01-02"))
	for _, entry := range entries {
		fmt.Printf("- %v: %v\n", entry.Title, entry.Dur)
	}

	// ожидаемый результат
	/*
		Мои достижения за 2022-04-15
		- Напряженная работа: 8h0m0s
		- Безудержное веселье: 3h0m0s
		- Оглаживание кота: 1h45m0s
		- Интернеты: 1h0m0s
		- Обед: 45m0s
		- Завтрак: 30m0s
	*/
}
