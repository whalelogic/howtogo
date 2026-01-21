// Package database provides functions to initialize and interact with the analytics database.
package database

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
)

// AnalyticsStore handles all database interactions
type AnalyticsStore struct {
	DB *sql.DB
}

// InitAnalyticsDB opens the database, sets WAL mode, and creates tables.
func InitAnalyticsDB(filepath string) (*AnalyticsStore, error) {
	db, err := sql.Open("sqlite", filepath)
	if err != nil {
		return nil, err
	}

	// Allows concurrent reads/writes.
	if _, err := db.Exec("PRAGMA journal_mode=WAL;"); err != nil {
		return nil, err
	}

	// 2. Set Busy Timeout
	if _, err := db.Exec("PRAGMA busy_timeout=5000;"); err != nil {
		return nil, err
	}

	// 3. Create Table
	query := `
	CREATE TABLE IF NOT EXISTS visits (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ip_address TEXT NOT NULL,
		path TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	CREATE INDEX IF NOT EXISTS idx_ip ON visits(ip_address);
	`
	if _, err := db.Exec(query); err != nil {
		return nil, err
	}

	return &AnalyticsStore{DB: db}, nil
}

// RecordVisit runs asynchronously to prevent blocking the HTTP response.
func (s *AnalyticsStore) RecordVisit(ip string, path string) {
	go func() {
		_, err := s.DB.Exec("INSERT INTO visits (ip_address, path) VALUES (?, ?)", ip, path)
		if err != nil {
			log.Printf("Background DB Error: Failed to record visit: %v", err)
		}
	}()
}

// GetStats returns total views and unique visitors.
func (s *AnalyticsStore) GetStats() (int, int) {
	var totalViews int
	var uniqueVisitors int

	row := s.DB.QueryRow(`
		SELECT 
			(SELECT COUNT(*) FROM visits),
			(SELECT COUNT(DISTINCT ip_address) FROM visits)
	`)
	
	if err := row.Scan(&totalViews, &uniqueVisitors); err != nil {
		log.Printf("Error reading stats: %v", err)
		return 0, 0
	}

	return totalViews, uniqueVisitors
}
