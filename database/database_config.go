package database

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"reflect"
)

// Init 用来初始化和测试数据库
func Init() error {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4")
	if err != nil {
		return err
	}
	err = CreateBase(db)
	if err != nil {
		return err
	}
	return nil
}

// CreateBase 用来创建一个database并且测试
func CreateBase(db *sql.DB) error {
	err := CreateTest(db)
	if err != nil {
		return err
	}
	err = InsertTest(db)
	if err != nil {
		return err
	}
	err = SelectTest(db)
	if err != nil {
		return err
	}
	err = DeleteTest(db)
	if err != nil {
		return err
	}
	return nil
}

// CreateTest 创建数据库和表并测试
func CreateTest(db *sql.DB) error {
	// if _, err := db.Query("drop database if exists test"); err != nil {
	// 	return err
	// // }
	// fmt.Print("1")
	// if _, err := db.Query("create database test"); err != nil {
	// 	return err
	// }
	// fmt.Print("2")
	return nil

}

// InsertTest 插入数据并测试
func InsertTest(db *sql.DB) error {
	if _, err := db.Query("create table testtable (col1 int,col2 varchar(20),col3 varchar(10))"); err != nil {
		return err
	}
	if _, err := db.Query("insert into testtable values (101,'小明','北京市'),(102,'小红','天津市')"); err != nil {
		return err
	}
	return nil
}

// SelectTest 查表测试
func SelectTest(db *sql.DB) error {
	selectResult, err := db.Query("select * from testtable")
	if err != nil {
		return err
	}
	values := reflect.ValueOf(selectResult)
	fmt.Println(values)
	return nil
}

// DeleteTest 删表，删base测试
func DeleteTest(db *sql.DB) error {
	if _, err := db.Query("delete from testtable where col1=101"); err != nil {
		return err
	}
	err := SelectTest(db)
	if err != nil {
		return err
	}
	if _, err := db.Query("drop table testtable "); err != nil {
		return err
	}
	// if _, err := db.Query("drop base test"); err != nil {
	// 	return err
	// }
	return nil
}
