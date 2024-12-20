package main

import (
	"database/sql"
	"fmt"
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
	a.db, err = sql.Open("sqlite3", "./menuitems.sqlite")
	if err != nil {
		return err
	}
	defer a.db.Close()

	// 创建示例表
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS conn_collection  (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		conn_name TEXT NOT NULL,
		conn_ip TEXT NOT NULL,
		conn_port INTEGER NOT NULL,
		user_name TEXT,
		password TEXT,
		key TEXT
	);
	`
	_, err = a.db.Exec(createTableQuery)
	if err != nil {
		return err
	}

	log.Println("Database initialized successfully.")
	return nil
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
