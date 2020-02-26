package module

/*
create DATABASE myblog;
CREATE USER 'admin'@'localhost' IDENTIFIED WITH mysql_native_password BY 'admin';
GRANT ALL PRIVILEGES ON myblog.* To 'admin'@'localhost';
*/
/*
blog table
+--------------+------------+------+-----+---------+--------------+
| Field        | Type       | Null | Key | Default | Extra        |
+--------------+------------+------+-----+---------+--------------+
| id           | int(64)    | NO   | PRI | NULL    |AUTO_INCREMENT|
| title        | VARCHAR(50)| NO   |     | NULL    |              |
| author       | VARCHAR(20)| NO   |     | NULL    |              |
| timestamp    | int(11)    | NO   |     | NULL    |              |
| content      | MEDIUMTEXT | NO   |     | NULL    |              |
+--------------+------------+------+-----+---------+--------------+
comment
+--------------+------------+------+-----+---------+--------------+
| Field        | Type       | Null | Key | Default | Extra        |
+--------------+------------+------+-----+---------+--------------+
| id           | int(64)    | NO   |     | NULL    |			  |
| Cid          | int(64)    | NO   | PRI | NULL    |AUTO_INCREMENT|
| fromIP       | VARCHAR(20)| NO   |     | NULL    |              |
| timestamp    | int(11)    | NO   |     | NULL    |              |
| content      | MEDIUMTEXT | NO   |     | NULL    |              |
+--------------+------------+------+-----+---------+--------------+
*/
import (
	"database/sql"
	"io/ioutil"
	"log"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

type _Mysql interface {
	insert()
	update()
	query()
	delete()
}

const (
	dbname     = "myblog"
	dbuser     = "admin"
	dbpassword = "admin"
)

var db = &sql.DB{}

func init() {
	db, _ = sql.Open("mysql", dbuser+":"+dbpassword+"@/"+dbname+"?multiStatements=true")
	err := db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	db.SetConnMaxLifetime(time.Second)
	print("db open")
}
func tsStart() (*sql.Tx, error) {
	return db.Begin()
}

func createTables(filename string) (err error) {
	sqlBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	sqlTables := string(sqlBytes)
	defer db.Close()
	_, err = db.Exec(sqlTables)
	if err != nil {
		return
	}
	return nil
}
