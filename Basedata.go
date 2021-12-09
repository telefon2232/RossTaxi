package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	//"strings"
)
const (
	userName = "Andrey14045"
	password = "1488"
	ip       = "127.0.0.1"
	port     = "3306"
)
func main(){

	//connect mysql Anrey
//dbName := "rosstaxi"
//path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
//db, err := sql.Open("mysql", path)

	// connect mysql Pavel(terminal)
db, err := sql.Open("mysql", "root:87654321@/") 
if err != nil {
panic(err)
}
defer db.Close()

//create database
createdb,err := db.Query("CREATE DATABASE rosstaxi;")
if err != nil {
	panic(err)
}
defer createdb.Close()

//select database
use,err := db.Query("USE rosstaxi;")
if err != nil {
	panic(err)
}
defer use.Close()

//create teable users
createtable,err := db.Query("CREATE TABLE users ( ID CHAR(5), role VARCHAR(10), login VARCHAR(20), password VARCHAR(12),PRIMARY KEY(ID));")
if err != nil {
	panic(err)
}
defer createtable.Close()

//insert in table users
insert, err := db.Query("INSERT INTO users VALUES (00001, 'driver', 'Pavel', '123');")
if err != nil {
	panic(err)
}
defer insert.Close()


}