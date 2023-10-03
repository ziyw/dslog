package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type LogEntry struct {
	id        int
	createdAt time.Time
	logType   string
	logMsg    string
}

type Repo struct {
	db *sql.DB
}

func (r *Repo) Connect() error {
	connStr := "user=ziyan password=postgres dbname=logdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	r.db = db
	return nil
}

func (r *Repo) Close() {
	r.db.Close()
}

func (r *Repo) Insert(createdAt time.Time, logTyp, logMsg string) (int, error) {
	sqlStmt := `
	INSERT INTO dslog (created_at, logType, logMsg) 
	VALUES ($1, $2, $3)
	RETURNING id`

	id := 0
	err := r.db.QueryRow(sqlStmt, createdAt, logTyp, logMsg).Scan(&id)
	if err != nil {
		log.Fatal("error insert item: ", err)
		return 0, err
	}
	return id, nil
}

func (r *Repo) GetAll() []LogEntry {
	rows, err := r.db.Query("SELECT * FROM dslog")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	entries := []LogEntry{}

	for rows.Next() {
		var entry LogEntry
		if err := rows.Scan(&entry.id, &entry.createdAt, &entry.logType, &entry.logMsg); err != nil {
			log.Fatal(err)
		}
		entries = append(entries, entry)
	}
	return entries
}
