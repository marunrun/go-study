package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	{
		// create user table if not exists
		query := `
    CREATE TABLE  if not exists users  (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`

		// Executes the SQL query in our database. Check err to ensure there was no error.
		if _, err = db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{
		// insert a user

		username := "marin"
		password := "123456"
		createAt := time.Now()

		result, err := db.Exec("insert into users (username, password, created_at) VALUES ( ?,?,?)", username, password, createAt)

		id, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id)
	}

	{
		// query a single user
		var (
			id       int
			usernmae string
			password string
			createAt time.Time
		)
		query := `SELECT id,username,password,created_at from users where id = ?`
		if err := db.QueryRow(query, 1).Scan(&id, &usernmae, &password, &createAt); err != nil {
			if errors.Is(sql.ErrNoRows, err) {
				log.Println(err)
			} else {
				log.Fatal(err)
			}
		}
		fmt.Println(id, usernmae, password, createAt)
	}

	{
		// query all users
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		rows, err := db.Query("select id,username,password,created_at from  users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var users []user
		for rows.Next() {
			var u user
			if err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt); err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v \n", users)
	}
	{
		// delete user 1
		result, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result.RowsAffected())
	}

}
