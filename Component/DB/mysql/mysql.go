package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id   int
	age  int
	name string
}

// 定义一个全局对象db
var db *sql.DB

// InitDB 定义一个初始化数据库的函数
func InitDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:123456@tcp(192.168.11.239:3306)/test?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// QueryRow 查询
func QueryRow(id int) {
	sqlStr := "select id,name,age from user where id=?"
	var u user
	err := db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed,err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}

func QueryMultiRow(id int) {
	sql := "select id, name, age from user where id > ?"
	rows, err := db.Query(sql, id)
	if err != nil {
		fmt.Printf("query failed!err=%v", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

func Insert() {
	sql := "insert into user(name,age) values (?,?)"
	ret, err := db.Exec(sql, "王五", 38)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed,err:%v\n", err)
		return
	}
	fmt.Printf("insert success,the id is %d\n", theID)
}

func Update() {
	sql := "update user set age=? where id=?"
	ret, err := db.Exec(sql, 39, 3)
	if err != nil {
		fmt.Printf("update failed,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed,err:%v\n", err)
		return
	}
	fmt.Printf("update success,affected rows:%d\n", n)
}

func Delete() {
	sql := "delete from user where id=?"
	ret, err := db.Exec(sql, 3)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed,err:%v\n", err)
		return
	}
	fmt.Printf("delete success,affected rows:%d\n", n)
}

func Transaction() {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("begin trans failed,err:%v\n", err)
		return
	}
	sql := "update user set age=30 where id=?"
	ret1, err := tx.Exec(sql, 2)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec sql failed,err:%v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec ret1.RowsAffected failed,err:%v\n", err)
		return
	}

	sqlStr2 := "Update user set age =40 where id =?"
	ret2, err := db.Exec(sqlStr2, 2)
	if err != nil {
		fmt.Printf("exec sql2 failed,err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec ret2.RowsAffected() failed,err:%v\n", err)
		return
	}

	fmt.Println(affRow1, affRow2)
	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("提交事务...")
		tx.Commit()
	} else {
		tx.Rollback()
		fmt.Println("事务回滚...")
	}
	fmt.Printf("exec trans success!")

}
