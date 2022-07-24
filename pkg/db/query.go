package db

import (
	"fmt"
	"log"
)

type Password struct {
	Key      string
	Password string
}

var db = GetDatabase()

func GetAll() []Password {
	passwords := []Password{}
	rows, err := db.Query(`SELECT * FROM PASSWORDS;`)

	if err == nil {
		for rows.Next() {
			password := Password{}
			rows.Scan(&password.Key, &password.Password)
			passwords = append(passwords, password)
		}
	}

	return passwords
}

func Add(key string, value string) {
	_, err := db.Exec(fmt.Sprintf(`INSERT INTO PASSWORDS VALUES ("%s", "%s");`, key, value))
	if err != nil {
		log.Fatal(err)
	}
}
