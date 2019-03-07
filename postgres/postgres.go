package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Db is used for interacting with the database
type Db struct {
	*sql.DB
}

// New makes a new database using the conStr
func New(conStr string) (*Db, error) {
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}

// ConnString returns a connection string when passed in parameters
func ConStr(host string, port int, user string, dbName string) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbName,
	)
}

type User struct {
	ID         int
	Name       string
	Age        int
}

// GetUser is called within our user query for graphql
func (d *Db) GetUser(name string) []User {
	stmt, err := d.Prepare("SELECT * FROM users WHERE name=$1")
	if err != nil {
		fmt.Println("GetUser Preperation Err: ", err)
	}

	rows, err := stmt.Query(name)
	if err != nil {
		fmt.Println("GetUser Query Err: ", err)
	}

	// Create User to hold data
	var r User
	users := []User{}
	// Copy the columns into values point to User 
	for rows.Next() {
		err = rows.Scan(
			&r.ID,
			&r.Name,
			&r.Age,
		)
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}
		users = append(users, r)
	}

	return users
}