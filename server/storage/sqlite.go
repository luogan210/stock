package storage

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	_ "modernc.org/sqlite"
)

var (
	// globalDB 全局数据库实例
	globalDB *DB
	// once 确保只初始化一次
	once sync.Once
	// initError 初始化错误
	initError error
)

type DB struct {
	SQL *sql.DB
}

// OpenSQLite opens (and creates if missing) a SQLite database at the given path.
func OpenSQLite(path string) (*DB, error) {
	once.Do(func() {
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			initError = fmt.Errorf("mkdir data dir: %w", err)
			return
		}
		dsn := fmt.Sprintf("file:%s?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)", path)
		db, err := sql.Open("sqlite", dsn)
		if err != nil {
			initError = fmt.Errorf("open sqlite: %w", err)
			return
		}
		if err := db.Ping(); err != nil {
			initError = fmt.Errorf("ping sqlite: %w", err)
			return
		}
		globalDB = &DB{SQL: db}
	})
	return globalDB, initError
}

// Migrate applies minimal schema needed for the app.
func (d *DB) Migrate() error {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS stocks (
            id TEXT PRIMARY KEY,
            code TEXT NOT NULL,
            name TEXT NOT NULL,
            region TEXT,
            currency TEXT,
            category TEXT,
            enabled INTEGER DEFAULT 1,
            remark TEXT,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
        );`,
		`CREATE TABLE IF NOT EXISTS plans (
            id TEXT PRIMARY KEY,
            name TEXT NOT NULL,
            type TEXT,
            stock_code TEXT,
            stock_name TEXT,
            strategy TEXT,
            trading_strategy TEXT,
            target_price REAL,
            quantity INTEGER,
            stop_loss REAL,
            take_profit REAL,
            start_time TEXT,
            end_time TEXT,
            risk_level TEXT,
            description TEXT,
            remark TEXT,
            status TEXT DEFAULT 'active',
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
        );`,
		`CREATE TABLE IF NOT EXISTS logs (
            id TEXT PRIMARY KEY,
            plan_name TEXT,
            stock_code TEXT NOT NULL,
            stock_name TEXT,
            type TEXT NOT NULL,
            trading_time TEXT NOT NULL,
            price REAL NOT NULL,
            quantity INTEGER NOT NULL,
            strategy TEXT,
            remark TEXT,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
        );`,
		`CREATE TABLE IF NOT EXISTS reviews (
            id TEXT PRIMARY KEY,
            period TEXT NOT NULL,
            review_date TEXT NOT NULL,
            title TEXT NOT NULL,
            buy_count INTEGER DEFAULT 0,
            sell_count INTEGER DEFAULT 0,
            total_profit REAL DEFAULT 0,
            summary TEXT NOT NULL,
            improvements TEXT,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
        );`,
	}
	for _, s := range stmts {
		if _, err := d.SQL.Exec(s); err != nil {
			return fmt.Errorf("migrate: %w", err)
		}
	}
	return nil
}

func (d *DB) Close() error { return d.SQL.Close() }

// Exec 执行SQL语句
func (d *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.SQL.Exec(query, args...)
}

// QueryRow 执行查询单行
func (d *DB) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.SQL.QueryRow(query, args...)
}

// Query 执行查询多行
func (d *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return d.SQL.Query(query, args...)
}

// GetDB 获取全局数据库实例
func GetDB() *DB {
	if globalDB == nil {
		panic("数据库未初始化，请先调用 storage.OpenSQLite()")
	}
	return globalDB
}
