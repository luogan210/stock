package db

import (
	"database/sql"
	"fmt"
	"sync"

	"server/storage"
)

var (
	// GlobalDB 全局数据库连接
	GlobalDB *storage.DB
	// once 确保只初始化一次
	once sync.Once
	// initError 初始化错误
	initError error
)

// Init 初始化全局数据库连接
func Init(db *storage.DB) error {
	once.Do(func() {
		if db == nil {
			initError = fmt.Errorf("数据库连接不能为空")
			return
		}
		GlobalDB = db
	})
	return initError
}

// GetDB 获取全局数据库连接
func GetDB() *storage.DB {
	return GlobalDB
}

// MustGetDB 获取全局数据库连接，如果未初始化则panic
func MustGetDB() *storage.DB {
	if GlobalDB == nil {
		panic("数据库未初始化，请先调用 db.Init()")
	}
	return GlobalDB
}

// GetSQL 获取底层的sql.DB连接
func GetSQL() *sql.DB {
	if GlobalDB == nil {
		return nil
	}
	return GlobalDB.SQL
}
