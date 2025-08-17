package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitSQLite(filepath string) *sql.DB {
	db, err := sql.Open("sqlite", filepath)
	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}

	// Single connection mode for SQLite
	db.SetMaxOpenConns(1)

	// Enable WAL for durability
	_, _ = db.Exec("PRAGMA journal_mode=WAL;")

	// Run migrations
	runMigrations(db)

	return db
}

func runMigrations(db *sql.DB) {
	// Example: create tables if not exist
	schema := `
    CREATE TABLE IF NOT EXISTS auth_users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT UNIQUE,
        password TEXT,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		price REAL NOT NULL,
		cgst_rate REAL NOT NULL DEFAULT 0,   -- CGST % for this product
		sgst_rate REAL NOT NULL DEFAULT 0  -- SGST % for this product
	);

	CREATE TABLE IF NOT EXISTS sales (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		product_id INTEGER NOT NULL,
		quantity INTEGER NOT NULL,
		unit_price REAL NOT NULL,            -- snapshot of product price at sale time
		cgst_rate REAL NOT NULL,             -- snapshot of CGST % at sale time
		sgst_rate REAL NOT NULL,             -- snapshot of SGST % at sale time
		cgst_amount REAL NOT NULL,           -- calculated CGST amount
		sgst_amount REAL NOT NULL,           -- calculated SGST amount
		line_total REAL NOT NULL,            -- (unit_price * quantity + taxes)
		subtotal REAL NOT NULL,              -- (unit_price * quantity)
		sold_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(product_id) REFERENCES products(id)
	);

    CREATE TABLE IF NOT EXISTS settings (
        key TEXT PRIMARY KEY,
        value TEXT
    );
    `
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatalf("failed to migrate DB: %v", err)
	}
}
