🤔 Фильм из JSON

Есть JSON-объект с информацией о фильме:

`
{
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
Требуется извлечь его в структуру:

type Genre string


type Movie struct {

Title  string

Year   int

Genres []Genre

}