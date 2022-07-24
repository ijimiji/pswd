package commands

import (
	"fmt"
	"pswd/pkg/db"
	"pswd/pkg/hash"
)

func Add() {
	var passphrase, key, password string
	fmt.Scanf("%s %s %s", &passphrase, &key, &password)
	db.Add(key, hash.HashPassword(password, passphrase))
}

func Get() {
	var passphrase string
	fmt.Scanf("%s", &passphrase)

	hashedPasswords := db.GetAll()

	for _, password := range hashedPasswords {
		fmt.Printf("%s: [%s]\n", password.Key, hash.UnhashPassword(password.Password, passphrase))
	}
}
