package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const connStr = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

type SQLMap struct {
	database *sql.DB
}

// NewSQLMap создает новую SQL-карту в указанной базе
func NewSQLMap(db *sql.DB) (*SQLMap, error) {
	query := `create table if not exists map(key text primary key, val int)`
	_, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	return &SQLMap{database: db}, nil
}

// Get возвращает значение для указанного ключа.
// Если такого ключа нет - возвращает ошибку sql.ErrNoRows.
func (m *SQLMap) Get(key string) (any, error) {
	var res any
	row := m.database.QueryRow(`select val from map where key = $1`, key) //выполняем select-запрос и получаем указатель на результат sql.Rows
	err := row.Scan(&res)                                                 //cчитываем поля текущей строки в поле переменной.
	if err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	} else if err != nil {
		return nil, err
	}
	return res, nil
}

// Set устанавливает значение для указанного ключа.
// Если такой ключ уже есть - затирает старое значение (это не считается ошибкой).
func (m *SQLMap) Set(key string, val any) error {
	_, err := m.database.Exec(`insert into map(key, val) values ($1, $2)
on conflict (key) do update set val = excluded.val`, key, val) //принимает текст запроса и значения параметров
	if err != nil {
		return err
	}
	return nil
}

func (m *SQLMap) Delete(key string) error {
	_, err := m.database.Exec(`delete from map where key = $1`, key)
	if err != nil {
		return err
	}
	return nil
}

func main() {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Open db succesfull")
	defer db.Close()

	bd, _ := NewSQLMap(db)
	fmt.Println(bd.Get("name"))
	fmt.Println(bd.Set("name", 77))
	fmt.Println(bd.Delete("name"))
	fmt.Println(bd.Get("name"))
}

/*
Через DB.Query() выполняем select-запрос и получаем указатель на результат sql.Rows
Через Rows.Close() гарантируем, что освободим ресурсы, занятые результатами.
Через Rows.Next() проходим по строкам результата. Если строк не осталось, он вернет false.
Через Rows.Scan() cчитываем поля текущей строки в поля структуры.
Через Rows.Err() проверяем наличие ошибок в результате.
Rows.Scan() автоматически преобразует типы данных СУБД в типы Go. Поддерживаются int, string, float64, bool, []byte и некоторые другие.
*/
