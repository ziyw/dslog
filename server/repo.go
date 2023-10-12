// TODO:
// - GetAllAsync() -> multiple threads, each one query the ddb at the same time
// - Unit tests

// How fast it can go?

package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type LogEntry struct {
	id        int
	createdAt time.Time
	logType   string
	logMsg    string
}

func (l LogEntry) nicePrintln() {
	fmt.Printf("%d %s %s %s\n", l.id, l.createdAt.Format(time.RFC3339), l.logType, l.logMsg)
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

func (r *Repo) GetById(id int) []LogEntry {
	sqlStmt := "SELECT * FROM dslog WHERE id = $1"

	rows, err := r.db.Query(sqlStmt, id)
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

func (r *Repo) GetByType(logType string) []LogEntry {
	sql := "SELECT * FROM dslog WHERE logType = $1;"
	rows, err := r.db.Query(sql, logType)
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

func (r *Repo) GetByTimeRange(startTime, endTime time.Time) []LogEntry {
	sql := "SELECT * FROM dslog WHERE created_at BETWEEN $1 AND $2"
	rows, err := r.db.Query(sql, startTime, endTime)
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

func (r *Repo) GetAll() []LogEntry {
	sql := "SELECT * FROM dslog ORDER BY id ASC"
	rows, err := r.db.Query(sql)
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

func (r *Repo) asyncGetAll(out chan LogEntry) {

	sql := "SELECT * FROM dslog ORDER BY id ASC"
	rows, err := r.db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var entry LogEntry
		if err := rows.Scan(&entry.id, &entry.createdAt, &entry.logType, &entry.logMsg); err != nil {
			log.Fatal(err)
		}
		out <- entry
	}
	defer close(out)
}

func (r Repo) asyncMessageProcess(in <-chan LogEntry, bufsize int) {
	var wg sync.WaitGroup
	for {
		for i := 0; i < bufsize; i++ {
			if entry, ok := <-in; ok {
				wg.Add(1)
				entry.nicePrintln()
				wg.Done()
			} else {
				wg.Wait()
				return
			}
			wg.Wait()
		}
	}
}

func main() {
	var repo Repo
	repo.Connect()
	defer repo.Close()

	result := make(chan LogEntry, 20)
	go repo.asyncGetAll(result)
	repo.asyncMessageProcess(result, 20)
}
