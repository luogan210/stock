package storage

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"server/utils"
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
		utils.LogInfo("正在创建数据库目录: %s", filepath.Dir(path))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			initError = fmt.Errorf("mkdir data dir: %w", err)
			utils.LogError("创建数据库目录失败: %v", err)
			return
		}
		dsn := fmt.Sprintf("file:%s?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)", path)
		utils.LogInfo("正在打开SQLite数据库: %s", path)
		db, err := sql.Open("sqlite", dsn)
		if err != nil {
			initError = fmt.Errorf("open sqlite: %w", err)
			utils.LogError("打开SQLite数据库失败: %v", err)
			return
		}
		if err := db.Ping(); err != nil {
			initError = fmt.Errorf("ping sqlite: %w", err)
			utils.LogError("数据库连接测试失败: %v", err)
			return
		}
		globalDB = &DB{SQL: db}
		utils.LogInfo("SQLite数据库连接成功")
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
            title TEXT,
            plan_name TEXT,
            stock_code TEXT NOT NULL,
            stock_name TEXT,
            type TEXT NOT NULL,
            trading_time TEXT NOT NULL,
            price REAL NOT NULL,
            quantity INTEGER NOT NULL,
            strategy TEXT,
            remark TEXT,
            status TEXT DEFAULT 'pending',
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

	// 执行创建表语句
	utils.LogInfo("正在创建数据库表...")
	for i, s := range stmts {
		tableNames := []string{"stocks", "plans", "logs", "reviews"}
		if i < len(tableNames) {
			utils.LogInfo("正在创建表: %s", tableNames[i])
		}
		if _, err := d.SQL.Exec(s); err != nil {
			utils.LogError("创建表失败: %v", err)
			return fmt.Errorf("migrate: %w", err)
		}
	}
	utils.LogInfo("数据库表创建完成")

	// 执行表结构更新语句（如果列不存在则添加）
	utils.LogInfo("正在检查并更新表结构...")
	alterStmts := []string{
		`ALTER TABLE logs ADD COLUMN title TEXT;`,
		`ALTER TABLE logs ADD COLUMN status TEXT DEFAULT 'pending';`,
	}

	for _, s := range alterStmts {
		// 使用 PRAGMA table_info 检查列是否存在
		if err := d.addColumnIfNotExists(s); err != nil {
			utils.LogError("更新表结构失败: %v", err)
			return fmt.Errorf("alter table: %w", err)
		}
	}
	utils.LogInfo("表结构检查完成")

	return nil
}

// addColumnIfNotExists 检查列是否存在，如果不存在则添加
func (d *DB) addColumnIfNotExists(alterStmt string) error {
	// 提取表名和列名
	var tableName, columnName string
	if alterStmt == `ALTER TABLE logs ADD COLUMN title TEXT;` {
		tableName = "logs"
		columnName = "title"
	} else if alterStmt == `ALTER TABLE logs ADD COLUMN status TEXT DEFAULT 'pending';` {
		tableName = "logs"
		columnName = "status"
	} else {
		return fmt.Errorf("unsupported alter statement: %s", alterStmt)
	}

	// 检查列是否存在
	checkQuery := `SELECT COUNT(*) FROM pragma_table_info(?) WHERE name = ?`
	var count int
	err := d.SQL.QueryRow(checkQuery, tableName, columnName).Scan(&count)
	if err != nil {
		return fmt.Errorf("check column exists: %w", err)
	}

	// 如果列不存在，则添加
	if count == 0 {
		utils.LogInfo("表 %s 中缺少列 %s，正在添加...", tableName, columnName)
		if _, err := d.SQL.Exec(alterStmt); err != nil {
			utils.LogError("添加列失败: %v", err)
			return fmt.Errorf("add column: %w", err)
		}
		utils.LogInfo("成功添加列 %s.%s", tableName, columnName)
	} else {
		utils.LogDebug("列 %s.%s 已存在，跳过", tableName, columnName)
	}

	return nil
}

func (d *DB) Close() error {
	if err := d.SQL.Close(); err != nil {
		utils.LogError("关闭数据库连接时出错: %v", err)
		return err
	}
	utils.LogInfo("数据库连接已关闭")
	return nil
}

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
