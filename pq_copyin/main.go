package main

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
)

func main() {
	records := [][]string{
		{"9", ""},
	}
	
	/*
	Needs test_datatype table set up in postgres database.
	schema for test_datatype:
	postgres=# \d+ test_datatype
                                             Table "public.test_datatype"
	 Column |  Type   |                         Modifiers                          | Storage | Stats target | Description 
	--------+---------+------------------------------------------------------------+---------+--------------+-------------
	 id     | integer | not null default nextval('test_datatype_id_seq'::regclass) | plain   |              | 
	 d      | date    | not null default ('now'::text)::date                       | plain   |              | 
	*/
	db, err := sql.Open("postgres", "dbname=postgres user=postgres password=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("open: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("open ping: %v", err)
	}
	defer db.Close()

	txn, err := db.Begin()
	if err != nil {
		log.Fatalf("begin: %v", err)
	}

	stmt, err := txn.Prepare(pq.CopyIn("test_datatype", "id", "d"))
	if err != nil {
		log.Fatalf("prepare: %v", err)
	}

	for _, r := range records {
		_, err = stmt.Exec(r[0], r[1])
		if err != nil {
			log.Fatalf("exec: %v", err)
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatalf("exec: %v", err)
	}

	err = stmt.Close()
	if err != nil {
		log.Fatalf("stmt close: %v", err)
	}

	err = txn.Commit()
	if err != nil {
		log.Fatalf("commit: %v", err)
	}
}
