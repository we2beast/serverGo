package handlers

import (
	"net/http"
	"encoding/json"
	"fmt"
	configDB "Calday-Server/config"
)

type users struct {
	Users []User
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	usersAll := users{}

	err := queryUsers(&usersAll)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.Marshal(usersAll)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, string(out))
	return
}

func queryUsers(usersAll *users) error {
	db, err := configDB.InitDB()
	if err != nil {
		return err
	}

	rows, err := db.Query(`SELECT id, firstname, lastname, email, phone, role, subscribe, date_subscribe FROM users`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Phone,
			&user.Role,
			&user.Subscribe,
			&user.DateSubscribe,
		)
		if err != nil {
			return err
		}
		usersAll.Users = append(usersAll.Users, user)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}
