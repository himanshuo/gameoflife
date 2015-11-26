package  main 

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"log"

)

var db *sql.DB
func main(){
	startDB()
}
func startDB() error{
	var err error
	//sqlite 3 database is stored in /data/data.db file
	db, err = sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return err
	}

	//this makes sure we can actually query the database.
	if err := db.Ping(); err != nil {
  		return err
	}
	//check to see if tables in database already exist
	
	//statement creates Task table which only has a PK id and name textfields
	sqlStmt = "create table Task (id integer not null primary key autoincrement, name text);"
	//run statement to create table. return object of Result type is not needed.
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err
	}
	return nil
}