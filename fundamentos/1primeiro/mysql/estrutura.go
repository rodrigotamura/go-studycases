package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // temos que colocar implicitamente para usar com o database/sql
)

type usuario struct {
	id   int
	nome string
}

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}

	return result
}

func main() {
	db, err := sql.Open("mysql", "admin:root@tcp(localhost:3306)/cursogo")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	exec(db, "Create database if not exists cursogo")
	exec(db, "use cursogo")
	exec(db, "drop table if exists usuarios")
	exec(db, `create table usuarios(
			id integer auto_increment,
			nome varchar(80),
			PRIMARY KEY (id)
		)`)

	stmt, _ := db.Prepare("insert into usuarios(nome) values (?)")
	stmt.Exec("Rodrigo")
	stmt.Exec("Debora")
	stmt.Exec("Levi")

	rows, _ := db.Query("select * from usuarios")
	defer rows.Close()

	for rows.Next() {
		var u usuario

		rows.Scan(&u.id, &u.nome)

		fmt.Println(u)
	}
}
