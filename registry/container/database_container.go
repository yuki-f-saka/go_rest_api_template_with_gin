package container

import (
	"database/sql"
	"fmt"
	"sync"

	"go_rest_api_template_with_gin/packages/env"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once // 本当に1回だけ実行されたかどうかの記録を保持しておくためにグローバル変数としている
)

// GetDatabase はデータベース接続のシングルトンインスタンスを返します。
func GetDatabase() (*sql.DB, error) {
	var err error
	once.Do(func() {
		cfg := env.Env()
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DBUser,
			cfg.DBPassword,
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBName,
		)
		db, err = sql.Open(cfg.DBDriver, dsn)
		if err != nil {
			return
		}
		if err = db.Ping(); err != nil {
			return
		}
	})
	return db, err
}

/*
sync.Onceを使用しない場合は、以下のようにinit()関数で初期化でも良い
*/
// var db *sql.DB

// func init() {
//     var err error
//     db, err = sql.Open(...)
//     if err != nil {
//         panic(err)
//     }
// }

// func GetDatabase() (*sql.DB, error) {
//     return db, nil
// }
