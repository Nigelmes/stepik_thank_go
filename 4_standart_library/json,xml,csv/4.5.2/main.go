package main

import (
	"encoding/json"
	"fmt"
)

// начало решения

// Genre описывает жанр фильма
type Genre string

// Movie описывает фильм
type Movie struct {
	Title  string  `json:"name"`
	Year   int     `json:"released_at"`
	Genres []Genre `json:"tags"`
}

func (a *Genre) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var name struct {
		Name string `json:"name"`
	}
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	*a = Genre(name.Name)
	return nil
}

// конец решения

func main() {
	const src = `{
		"name": "Interstellar",
		"released_at": 2014,
		"director": "Christopher Nolan",
		"tags": [
			{ "name": "Adventure" },
			{ "name": "Drama" },
			{ "name": "Science Fiction" }
		],
		"duration": "2h49m",
		"rating": "★★★★★"
	}`

	var m Movie
	err := json.Unmarshal([]byte(src), &m)
	fmt.Println(err)
	// nil
	fmt.Println(m)
	// {Interstellar 2014 [Adventure Drama Science Fiction]}
}
