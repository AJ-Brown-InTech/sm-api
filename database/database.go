package main

import (
      "context"
      "database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func Libre_Db() (*sql.DB, error){
	db, err := sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	

	ctx := context.Background()
	_, err = db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS test (id INTEGER PRIMARY KEY, test TEXT)")
	if err != nil {
		panic(err)
	}
	
	return db, nil

}