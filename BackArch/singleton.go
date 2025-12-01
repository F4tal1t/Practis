package main

import (
	"fmt"
	"sync"
)

// Database represents a singleton database connection
type Database struct {
	connectionString string
	isConnected     bool
}

var (
	dbInstance *Database
	once       sync.Once
)

// GetDatabaseInstance returns the singleton database instance
func GetDatabaseInstance() *Database {
	once.Do(func() {
		fmt.Println("Creating database instance...")
		dbInstance = &Database{
			connectionString: "postgres://localhost:5432/mydb",
			isConnected:     false,
		}
	})
	return dbInstance
}

// Connect simulates connecting to the database
func (db *Database) Connect() {
	if !db.isConnected {
		fmt.Printf("Connecting to database: %s\n", db.connectionString)
		db.isConnected = true
	} else {
		fmt.Println("Database is already connected")
	}
}

// Query simulates executing a database query
func (db *Database) Query(sql string) {
	if db.isConnected {
		fmt.Printf("Executing query: %s\n", sql)
	} else {
		fmt.Println("Database not connected. Please connect first.")
	}
}

// Disconnect simulates disconnecting from the database
func (db *Database) Disconnect() {
	if db.isConnected {
		fmt.Println("Disconnecting from database...")
		db.isConnected = false
	} else {
		fmt.Println("Database is already disconnected")
	}
}

// GetConnectionString returns the connection string
func (db *Database) GetConnectionString() string {
	return db.connectionString
}

// IsConnected returns the connection status
func (db *Database) IsConnected() bool {
	return db.isConnected
}
