package db

import (
	"database/sql"
	"log"
)

func InitSQLiteDB(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatalf("Erro ao abrir o banco de dados: %v", err)
	}
	// Configura o modo WAL para melhorar o suporte a concorrência
	_, err = db.Exec("PRAGMA journal_mode=WAL;")
	if err != nil {
		log.Fatalf("Erro ao configurar o modo WAL: %v", err)
	}
	createTables(db)
	return db
}

func createTables(db *sql.DB) {
	createProjectsTable := `
	CREATE TABLE IF NOT EXISTS projects (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		squad TEXT,
		identifier TEXT,
		source TEXT,
		consultation_date DATETIME
	);
	`
	_, err := db.Exec(createProjectsTable)
	if err != nil {
		log.Fatalf("Erro ao criar a tabela projects: %v", err)
	}

	createMetricsTable := `
	CREATE TABLE IF NOT EXISTS metrics (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		project_id INTEGER,
		metric TEXT,
		value TEXT,
		best_value BOOLEAN,
		FOREIGN KEY(project_id) REFERENCES projects(id)
	);
	`
	_, err = db.Exec(createMetricsTable)
	if err != nil {
		log.Fatalf("Erro ao criar a tabela metrics: %v", err)
	}
}
