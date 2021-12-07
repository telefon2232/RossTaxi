package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)
const (
	userName = "Andrey14045"
	password = "1488"
	ip       = "127.0.0.1"
	port     = "3306"
)
func main(){

dbName := "rosstaxi"
path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
db, err := sql.Open("mysql", path)
if err != nil {
panic(err)
}

defer db.Close()

insert, err := db.Query("INSERT  INTO users VALUES('VlaDick')")

if err != nil {
panic(err)
}
defer insert.Close()


}