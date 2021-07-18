package model

import (
	"database/sql"
	"gindemo/database"
)

// User 用户
type User struct {
	Id         int    `json:"id" form:"id"`
	Username   string `json:"username" form:"username"`
	Password   string `json:"password" form:"password"`
	Status     int    `json:"status" form:"status"`
	CreateTime int64  `json:"createTime" form:"createTime"`
}

// ----------------- 数据库操作 ----------------------

// ------ 查询 -------

// QueryUserWithSql 依据条件查询用户
func QueryUserWithSql(DB *sql.DB, sql string) int {
	rows, _ := database.QueryRowsDB(DB, sql)
	id := 0
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			continue
		}
		if id > 0 {
			break
		}
	}
	return id
}

// QueryIdWithUsername 依据用户 username 查询
func QueryIdWithUsername(DB *sql.DB, username string) int {
	sql := "select id from user where username='" + username + "'"
	return QueryUserWithSql(DB, sql)
}

// QueryIdWithUserPwd 依据用户名和密码查询
func QueryIdWithUserPwd(DB *sql.DB, username string, password string) int {
	sql := "select id from user where username='" + username + "' and password='" + password + "'"
	return QueryUserWithSql(DB, sql)
}

// ------ 插入 -------

// InsertUser 插入
func InsertUser(DB *sql.DB, user User) (int64, error) {
	return database.ModifyDB(DB, "insert into user(username,password,status,createTime) values (?,?,?,?)",
		user.Username, user.Password, user.Status, user.CreateTime)
}

// ------ 更新 -------
// 插入

// ------ 删除 -------
// 删除
