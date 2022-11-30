package main

import (
	"database/sql"
	"fmt"

	"github.com/gomodul/dbssh"
	"github.com/gomodul/dbsshpg"
)

func main() {
	ssh, driverName, err := dbsshpg.Open(dbssh.Config{
		Host: "127.0.0.1",
		Port: "22",
		User: "root",
		Pass: "password",
	})
	if err != nil {
		panic(err)
	}
	defer ssh.Close()

	db, err := sql.Open(driverName, "host=localhost user=root password= dbname=database_name port=5432 sslmode=disable TimeZone=Asia/Jakarta")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connected to PostgreSQL via SSH")
}
