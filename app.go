package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type App struct {
	db *sql.DB
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup() error {
	var err error
	a.db, err = sql.Open("sqlite3", "./literm")
	if err != nil {
		return err
	}
	defer a.db.Close()

	// 创建示例表
	createConnInfo := `
			create table if not exists conn_info
		(
			id            INTEGER
				primary key autoincrement,
			conn_name     TEXT,
			conn_ip       TEXT,
			conn_port     INTEGER,
			user_name     TEXT,
			password      TEXT,
			key           TEXT,
			collection_id INTEGER
		);
	`
	_, err = a.db.Exec(createConnInfo)
	if err != nil {
		return err
	}

	createCollection := `
		create table if not exists conn_collection
		(
			id            INTEGER
				primary key autoincrement,
			collection_name TEXT
		);
	`
	_, err = a.db.Exec(createCollection)
	if err != nil {
		return err
	}

	log.Println("Database initialized successfully.")
	return nil
}
