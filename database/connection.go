package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "0214807424"
	dbname   = "buku_panduan"
)

func main() {
	conn := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname= %s sslmode = disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", conn)
	CheckError(err)

	defer db.Close()

	buatkaryawan := `create table "Employee"(
		nama varchar(255),
		keterangan varchar(255)
	)`
	_, a := db.Exec(buatkaryawan)
	CheckError(a)

	insertkaryawan := `insert into "Employee"("nama", "keterangan") values('reza', 'pria tampan')`
	_, e := db.Exec(insertkaryawan)
	CheckError(e)

	insertkaryawanbeda := `insert into "Employee"("nama", "keterangan") values($1, $2)`
	_, w := db.Exec(insertkaryawanbeda, "wahyu", "pria tampan")
	CheckError(w)
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
