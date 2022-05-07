package main

import (
	"database/sql"
	"log"

	"fmt"

	_ "github.com/proullon/ramsql/driver"
)

func main() {
	db, err := sql.Open("ramsql", "goimdb")
	if err != nil {
		log.Fatal("error", err)
		return
	}

	createTb := `
	CREATE TABLE IF NOT EXISTS goimdb (
	id INT AUTO_INCREMENT,
	imdbID TEXT NOT NULL UNIQUE,
	title TEXT NOT NULL,
	year INT NOT NULL,
	rating FLOAT NOT NULL,
	isSuperHero BOOLEAN NOT NULL,
	PRIMARY KEY (id)
	);
	`

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("error:", err)
		return
	}

	fmt.Println("table created.")

	insert := `
	INSERT INTO goimdb(imdbID,title,year,rating,isSuperHero)
	VALUES (?,?,?,?,?);
	`

	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal("prepare statement error:", err)
	}
	r, err := stmt.Exec("tt4154796", "Avengers: Endgame", 2019, 8.4, true)
	if err != nil {
		log.Fatal("insert error:", err)
	}
	l, err := r.LastInsertId()
	fmt.Println("lastINsertId", l, "err:", err)
	ef, err := r.RowsAffected()
	fmt.Println("RowsAffecred", ef, "err:", err)

	rows, err := db.Query(`
	SELECT id, imdbID, title, year, rating, isSuperHero
	FROM goimdb
	`)
	if err != nil {
		log.Fatal("query error", err)
	}

	for rows.Next() {
		var id int
		var imdbID, title string
		var year int
		var rating float32
		var isSuperHero bool

		err := rows.Scan(&id, &imdbID, &title, &year, &rating, &isSuperHero)
		if err != nil {
			log.Fatal("for rows error", err)
		}
		fmt.Println("row:", id, imdbID, title, year, rating, isSuperHero)
	}

	stmt2, err := db.Prepare(`
	UPDATE goimdb
	SET rating=?
	WHERE imdbID=?
	`)
	_, err = stmt2.Exec(9.2, "tt4154796")
	if err != nil {
		log.Fatal("update error", err)
	}

	rowx := db.QueryRow(`SELECT id, imdbID, title, year, rating, isSuperHero 
	FROM goimdb WHERE imdbID=?`, "tt4154796")
	var id int
	var imdbID, title string
	var year int
	var rating float32
	var isSuperHero bool
	err = rowx.Scan(&id, &imdbID, &title, &year, &rating, &isSuperHero)
	if err != nil {
		log.Fatal("scan one rowx error", err)
	}

	fmt.Println("one rowx:", id, imdbID, title, year, rating, isSuperHero)
}
