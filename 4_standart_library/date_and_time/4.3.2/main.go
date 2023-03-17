package main

import (
	"errors"
	"fmt"
	"time"
)

// начало решения

// TimeOfDay описывает время в пределах одного дня
type TimeOfDay struct {
	hour, min, sec int
	loc            *time.Location
}

// Hour возвращает часы в пределах дня
func (t TimeOfDay) Hour() int {
	return t.hour
}

// Minute возвращает минуты в пределах часа
func (t TimeOfDay) Minute() int {
	return t.min
}

// Second возвращает секунды в пределах минуты
func (t TimeOfDay) Second() int {
	return t.sec
}

// String возвращает строковое представление времени
// в формате чч:мм:сс TZ (например, 12:34:56 UTC)
func (t TimeOfDay) String() string {
	return fmt.Sprintf("%02d:%02d:%02d %v", t.hour, t.min, t.sec, t.loc)
}

// Equal сравнивает одно время с другим.
// Если у t и other разные локации - возвращает false.
func (t TimeOfDay) Equal(other TimeOfDay) bool {
	if t.loc.String() != other.loc.String() {
		return false
	}
	a, b := convert(t, other)
	return a.Equal(b)
}

func convert(t1, t2 TimeOfDay) (time.Time, time.Time) {
	a := time.Date(2020, 1, 2, t1.hour, t1.min, t1.sec, 0, t1.loc)
	b := time.Date(2020, 1, 2, t2.hour, t2.min, t2.sec, 0, t2.loc)
	return a, b
}

// Before возвращает true, если время t предшествует other.
// Если у t и other разные локации - возвращает ошибку.
func (t TimeOfDay) Before(other TimeOfDay) (bool, error) {
	if t.loc.String() != other.loc.String() {
		return false, errors.New("not implemented")
	}
	a, b := convert(t, other)
	if a.Before(b) {
		return true, nil
	}

	return false, nil
}

// After возвращает true, если время t идет после other.
// Если у t и other разные локации - возвращает ошибку.
func (t TimeOfDay) After(other TimeOfDay) (bool, error) {
	if t.loc.String() != other.loc.String() {
		return false, errors.New("not implemented")
	}
	a, b := convert(t, other)
	if a.After(b) {
		return true, nil
	}

	return false, nil
}

// MakeTimeOfDay создает время в пределах дня
func MakeTimeOfDay(hour, min, sec int, loc *time.Location) TimeOfDay {
	return TimeOfDay{hour, min, sec, loc}
}

// конец решения

func main() {
	t1 := MakeTimeOfDay(17, 1, 22, time.UTC)
	t2 := MakeTimeOfDay(20, 3, 4, time.UTC)

	if t1.Equal(t2) {
		fmt.Printf("%v should not be equal to %v", t1, t2)
	}

	before, _ := t1.Before(t2)
	if !before {
		fmt.Printf("%v should be before %v", t1, t2)
	}

	after, _ := t1.After(t2)
	if after {
		fmt.Printf("%v should NOT be after %v", t1, t2)
	}
}
