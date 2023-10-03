package main

import (
	"database/sql"
	"fmt"
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

func (r *Repo) Connect() {
	connStr := "user=ziyan password=postgres dbname=logdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("fail open connection to db: %v", err)
	}
	r.db = db
}

func (r *Repo) Close() {
	r.db.Close()
}

func (r *Repo) Insert(createdAt, logTyp, logMsg string) (int, error) {
	sqlStmt := `
	INSERT INTO dslog (created_at, logType, logMsg) 
	VALUES ($1, $2, $3)
	RETURNING id`

	id := 0
	err := r.db.QueryRow(sqlStmt, createdAt, logTyp, logMsg).Scan(&id)
	if err != nil {
		fmt.Errorf("error insert item: %v", err)
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

		// date, err := time.Parse(time.RFC3339, createdAt)
		// if err != nil {
		// 	log.Fatal("parse date error", err)
		// }

		// fmt.Println(id, date, logType, logMsg)
	}
	return entries
}

func main() {

	var repo Repo
	repo.Connect()
	defer repo.Close()

	// repo.Insert(time.Now().Format(time.RFC1123), "ERROR", "INSERT ERROR MESSAGE")
	myEntries := repo.GetAll()
	for _, e := range myEntries {
		fmt.Printf("%+v\n", e)
	}

	// connStr := "user=ziyan password=postgres dbname=logdb sslmode=disable"
	// MyDb, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	log.Fatalf("fail open connect to db: %v", err)
	// }
	// defer MyDb.Close()

	// pingErr := db.Ping()
	// if pingErr != nil {
	// 	log.Fatalf("fail ping db: %v", err)
	// }

	// rows, err := db.Query("SELECT * FROM dslog")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var id int
	// 	var createdAt, logType, logMsg string
	// 	if err := rows.Scan(&id, &createdAt, &logType, &logMsg); err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	date, err := time.Parse(time.RFC3339, createdAt)
	// 	if err != nil {
	// 		log.Fatal("parse date error", err)
	// 	}

	// 	fmt.Println(id, date, logType, logMsg)
	// }

	// sqlStmt := `
	// 	INSERT INTO
	// `

	// id := 0
	// db.Query(sqlStmt, )

}
