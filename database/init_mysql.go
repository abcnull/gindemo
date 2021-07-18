package database

import (
	"database/sql"
	"gindemo/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// DB 数据库
var DB *sql.DB

// InitMysql 初始化 mysql 数据库
func InitMysql() {
	// 连接数据库
	if DB == nil {
		DB, _ = sql.Open("mysql", config.DB_USER+":"+config.DB_PWD+"@tcp("+config.DB_ADD+":"+config.DB_PORT+")/"+config.DB_NAME)
		// 如果表已经存在则不会创建或覆盖
		CreateTableWithUser(DB)
		CreateTableWithArticle(DB)
	}
}

// CreateTableWithUser 创建用户表如果不存在
func CreateTableWithUser(DB *sql.DB) (int64, error) {
	// 创建表结构 sql
	sql := `CREATE TABLE IF NOT EXISTS user(
        id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
        username VARCHAR(64),
        password VARCHAR(64),
        status INT(4),
        createTime INT(10),
		lastUpdateTime INT(10)
        );`
	// 执行 sql
	return ModifyDB(DB, sql)
}

// CreateTableWithArticle 创建文章表如果不存在
func CreateTableWithArticle(DB *sql.DB) (int64, error) {
	sql := `create table if not exists article(
        id int(4) primary key auto_increment not null,
        title varchar(30),
        author varchar(20),
        tags varchar(30),
        short varchar(255),
        content longtext,
        createtime int(10)
        );`
	// 执行 sql
	return ModifyDB(DB, sql)
}

// ModifyDB 执行 sql
func ModifyDB(DB *sql.DB, sql string, args ...interface{}) (int64, error) {
	// Exec 执行 sql
	result, err := DB.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	// 拿到影响的行数
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

// QueryRowsDB 基本查询
func QueryRowsDB(DB *sql.DB, sql string) (*sql.Rows, error) {
	return DB.Query(sql)
}
