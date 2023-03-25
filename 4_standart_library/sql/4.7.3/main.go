package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

const connStr = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

type SQLMap struct {
	database *sql.DB
	state    map[string]*sql.Stmt
	timeout  time.Duration
}

func (m *SQLMap) SetTimeout(d time.Duration) {
	m.timeout = d
}

// NewSQLMap создает новую SQL-карту в указанной базе
func NewSQLMap(db *sql.DB) (*SQLMap, error) {
	query := `create table if not exists map(key text primary key, val int)`
	_, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	get, _ := db.Prepare(`select val from map where key = $1`)
	set, _ := db.Prepare(`insert into map(key, val) values ($1, $2)
on conflict (key) do update set val = excluded.val`)
	del, _ := db.Prepare(`delete from map where key = $1`)
	return &SQLMap{database: db,
		state: map[string]*sql.Stmt{
			"get": get,
			"set": set,
			"del": del,
		}, timeout: 10 * time.Second}, nil
}

// Get возвращает значение для указанного ключа.
// Если такого ключа нет - возвращает ошибку sql.ErrNoRows.
func (m *SQLMap) Get(key string) (any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()
	var res any
	stmt := m.state["get"]
	row := stmt.QueryRowContext(ctx, key)
	err := row.Scan(&res) //cчитываем поля текущей строки в поле переменной.
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
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()
	stmt := m.state["set"]
	_, err := stmt.ExecContext(ctx, key, val) //принимает текст запроса и значения параметров
	if err != nil {
		return err
	}
	return nil
}

func (m *SQLMap) SetItems(items map[string]any) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()
	stmt := m.state["set"]
	tx, err := m.database.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}
	txStmt := tx.StmtContext(ctx, stmt)
	for key, val := range items {
		_, err = txStmt.ExecContext(ctx, key, val)
		if err != nil {
			return err
		}
	}
	defer tx.Rollback()

	return tx.Commit()

}

func (m *SQLMap) Close() error {
	for _, stmt := range m.state {
		err := stmt.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *SQLMap) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()
	stmt := m.state["del"]
	_, err := stmt.ExecContext(ctx, key)
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

	m, err := NewSQLMap(db)
	if err != nil {
		panic(err)
	}
	defer m.Close()

	m.SetTimeout(10 * time.Millisecond)

	m.Set("name", 44)
	m.Get("name")
}

/*
Через DB.Query() выполняем select-запрос и получаем указатель на результат sql.Rows
Через Rows.Close() гарантируем, что освободим ресурсы, занятые результатами.
Через Rows.Next() проходим по строкам результата. Если строк не осталось, он вернет false.
Через Rows.Scan() cчитываем поля текущей строки в поля структуры.
Через Rows.Err() проверяем наличие ошибок в результате.
Rows.Scan() автоматически преобразует типы данных СУБД в типы Go. Поддерживаются int, string, float64, bool, []byte и некоторые другие.
*/
