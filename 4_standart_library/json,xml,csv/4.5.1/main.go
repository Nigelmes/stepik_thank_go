package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// начало решения

// Duration описывает продолжительность фильма
type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	temp := time.Duration(d)
	var str string = "\""
	if temp.Hours() >= 1 {
		str += fmt.Sprintf("%dh", int(temp.Hours())%24)
	}
	if int(temp.Minutes())%60 > 0 {
		str += fmt.Sprintf("%dm", int(temp.Minutes())%60)
	}
	str += "\""
	res := []byte(str)
	return res, nil
}

// Rating описывает рейтинг фильма
type Rating int

func (r Rating) MarshalJSON() ([]byte, error) {
	if r < 0 {
		return nil, fmt.Errorf("Рейтинг не может быть отрицательным")
	}
	count := int(r)
	var str string = "\""
	for i := 0; i < 5; i++ {
		if count <= 0 {
			str += "☆"
			continue
		}
		str += "★"
		count--
	}
	str += "\""
	res := []byte(str)
	return res, nil
}

// Movie описывает фильм
type Movie struct {
	Title    string
	Year     int
	Director string
	Genres   []string
	Duration Duration
	Rating   Rating
}

// MarshalMovies кодирует фильмы в JSON.
//   - если indent = 0 - использует json.Marshal
//   - если indent > 0 - использует json.MarshalIndent
//     с отступом в указанное количество пробелов.
func MarshalMovies(indent int, movies ...Movie) (string, error) {
	otstup := strings.Repeat(" ", indent)
	if indent < 0 {
		return "", fmt.Errorf("ident не может быть отрицательным")
	}
	if indent > 0 {
		res, err := json.MarshalIndent(movies, "", otstup)
		if err != nil {
			return "", err
		}
		return string(res), nil
	}
	res, err := json.Marshal(movies)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// конец решения

func main() {
	m1 := Movie{
		Title:    "Interstellar",
		Year:     2014,
		Director: "Christopher Nolan",
		Genres:   []string{"Adventure", "Drama", "Science Fiction"},
		Duration: Duration(2*time.Hour + 49*time.Minute),
		Rating:   5,
	}
	m2 := Movie{
		Title:    "Sully",
		Year:     2016,
		Director: "Clint Eastwood",
		Genres:   []string{"Drama", "History"},
		Duration: Duration(6540000000000),
		Rating:   4,
	}

	b, err := MarshalMovies(4, m1, m2)
	fmt.Println(err)
	// nil
	fmt.Println(string(b))
	/*
		[
		    {
		        "Title": "Interstellar",
		        "Year": 2014,
		        "Director": "Christopher Nolan",
		        "Genres": [
		            "Adventure",
		            "Drama",
		            "Science Fiction"
		        ],
		        "Duration": "2h49m",
		        "Rating": "★★★★★"
		    },
		    {
		        "Title": "Sully",
		        "Year": 2016,
		        "Director": "Clint Eastwood",
		        "Genres": [
		            "Drama",
		            "History"
		        ],
		        "Duration": "1h36m",
		        "Rating": "★★★★☆"
		    }
		]
	*/
}
