package users

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

type Userdata struct {
	ID          int
	Username    string
	Name        string
	Surname     string
	Description string
}

var (
	Hostname = ""
	Port     = 2345
	Username = ""
	Password = ""
	Database = ""
)

func openConnection() (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Hostname, Port, Username, Password, Database)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		fmt.Println("Open(): ", err)
		return nil, err
	}
	return db, nil
}

func exists(username string) int {
	username = strings.ToLower(username)

	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer db.Close()
	userId := -1
	statement := fmt.Sprintf(`select "id" from "users" where username='%s'`, username)
	rows, err := db.Query(statement)
	if err != nil {
		fmt.Println("Scan", err)
		return -1
	}

	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println("Scan", err)
			return -1
		}
		userId = id
	}
	defer rows.Close()
	return userId
}

func AddUser(d Userdata) int {
	d.Username = strings.ToLower(d.Username)
	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer db.Close()

	userId := exists(d.Username)

	if userId != -1 {
		fmt.Println("User already exist:", d.Username)
		return -1
	}

	insertStatements := `insert into "users" ("username") values ($1)`

	_, err = db.Exec(insertStatements, d.Username)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	userId = exists(d.Username)

	if userId == -1 {
		return userId
	}

	insertStatements = `insert into "userdata" ("userid", "name", "surname", "description") values ($1, $2, $3, $4)`

	_, err = db.Exec(insertStatements, userId, d.Name, d.Surname, d.Description)

	if err != nil {
		fmt.Println("db.Exex()", err)
		return -1
	}

	return userId

}

func ListUsers() ([]Userdata, error) {
	Data := []Userdata{}

}
