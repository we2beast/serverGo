package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	configDB "Calday-Server/config"
	"github.com/dgrijalva/jwt-go"
	"time"
	"strings"
	"strconv"
	"os"
)

var secretKey = []byte(os.Getenv("SESSION_SECRET"))

type UserUrl struct {
	ID int64
}

type User struct {
	ID int `json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Role int `json:"role"`
	Subscribe bool `json:"subscribe"`
	DateSubscribe string `json:"date_subscribe"`
}

type Edit struct {
	ID int `json:"id"`
	Success bool `json:"success"`
}

type Response struct {
	Token string `json:"token"`
	Status string `json:"status"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request)  {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
		return
	}

	db, err := configDB.InitDB()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	firstname := r.PostForm.Get("firstname")
	lastname := r.PostForm.Get("lastname")
	email := r.PostForm.Get("email")
	phone := r.PostForm.Get("phone")

	if firstname == "" || lastname == "" || email == "" || phone == "" {
		http.Error(w, "Fields are empty", http.StatusBadRequest)
		return
	}

	sqlStatement := `SELECT phone FROM users WHERE phone = $1`

	phoneDB := ""

	err = db.QueryRow(sqlStatement, phone).Scan(&phoneDB)

	if phone == phoneDB {
		http.Error(w, "Phone number is exists", http.StatusBadRequest)
		return
	} else {
		sqlStatement := `
		INSERT INTO users (firstname, lastname, email, phone, date_register, role, subscribe, date_subscribe)
		VALUES ($1, $2, $3, $4, NOW(), 0, false, null)
		RETURNING id`
		id := 0
		err = db.QueryRow(sqlStatement, firstname, lastname, email, phone).Scan(&id)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		} else {
			claims := jwt.MapClaims{
				"id": id,
				"ExpiresAt": 15000,
				"IssuedAt": time.Now().Unix(),
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(secretKey)
			if err != nil {
				w.WriteHeader(http.StatusBadGateway)
				w.Write([]byte(err.Error()))
			}
			response := Response{Token: tokenString, Status: "success"}
			responseJSON, _ := json.Marshal(response)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(responseJSON))
			return
		}
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request)  {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
		return
	}

	db, err := configDB.InitDB()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	phone := r.PostForm.Get("phone")
	codePhone, err := strconv.ParseInt(r.PostForm.Get("codePhone"), 10, 0)

	if phone == "" || err != nil {
		http.Error(w, "Fields are empty", http.StatusBadRequest)
		return
	}

	var randomCode int64 = 12345

	if codePhone !=  randomCode {
		http.Error(w, "Code is not right", http.StatusBadRequest)
		return
	} else {
		sqlStatement := `SELECT id FROM users WHERE phone = $1`
		id := ""
		err = db.QueryRow(sqlStatement, phone).Scan(&id)
		if id == "" {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		} else {
			id, err := strconv.ParseInt(id, 10, 0)
			claims := jwt.MapClaims{
				"id":        id,
				"ExpiresAt": 15000,
				"IssuedAt":  time.Now().Unix(),
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(secretKey)
			if err != nil {
				w.WriteHeader(http.StatusBadGateway)
				w.Write([]byte(err.Error()))
			}

			sqlStatement := `SELECT count(user_id) FROM auth_token WHERE user_id = $1`
			var idCount string
			err = db.QueryRow(sqlStatement, id).Scan(&idCount)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			subscribeQuery := `SELECT subscribe FROM users WHERE id = $1`
			var subscribe string
			err = db.QueryRow(subscribeQuery, id).Scan(&subscribe)

			if idCount > "2" && subscribe == "false" {
				http.Error(w, "Unauthorized. Subscribe not found", http.StatusUnauthorized)
				return
			} else {
				sqlStatement := `
				INSERT INTO auth_token (token, user_id)
				VALUES ($1, $2)`
				_, err = db.Exec(sqlStatement, tokenString, id)
				if err != nil {
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}

				response := Response{Token: tokenString, Status: "success"}
				responseJSON, _ := json.Marshal(response)
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, string(responseJSON))
			}
			return
		}
	}
}

//func getTokenHandler(w http.ResponseWriter, r *http.Request) {
//	err := r.ParseForm()
//	if err != nil {
//		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
//		return
//	}
//	username := r.PostForm.Get("username")
//	password := r.PostForm.Get("password")
//	if originalPassword, ok := usersOne[username]; ok {
//		if password == originalPassword {
//			claims := jwt.MapClaims{
//				"username": username,
//				"id": 1,
//				"ExpiresAt": 15000,
//				"IssuedAt": time.Now().Unix(),
//			}
//			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//			tokenString, err := token.SignedString(secretKey)
//			if err != nil {
//				w.WriteHeader(http.StatusBadGateway)
//				w.Write([]byte(err.Error()))
//			}
//			response := Response{Token: tokenString, Status: "success"}
//			responseJSON, _ := json.Marshal(response)
//			w.WriteHeader(http.StatusOK)
//			w.Header().Set("Content-Type", "application/json")
//			w.Write(responseJSON)
//		} else {
//			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
//			return
//		}
//	} else {
//		http.Error(w, "User is not found", http.StatusNotFound)
//		return
//	}
//}

func parseParams(r *http.Request, prefix string, num int) (int64, error) {
	url := strings.TrimPrefix(r.URL.Path, prefix)
	params := strings.Split(url, "/")
	if len(params) != num || len(params[0]) == 0 {
		return 0, fmt.Errorf("Bad format. Expecting exactly %d params", num)
	}
	if i, err := strconv.ParseInt(params[0], 10, 0); err == nil {
		return i, nil
	}
	return 0, nil
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	user := UserUrl{}
	userOne := User{}
	params, err := parseParams(r, "/api/user/", 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	user.ID = params
	err = queryUser(&userOne, params)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(userOne)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(out))
	return
}

func queryUser(userOne *User, params int64) error {
	db, err := configDB.InitDB()
	if err != nil {
		return err
	}

	rows, err := db.Query(`SELECT id, firstname, lastname, email, phone, role, subscribe, date_subscribe FROM users WHERE id = $1`, int(params))
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&userOne.ID,
			&userOne.FirstName,
			&userOne.LastName,
			&userOne.Email,
			&userOne.Phone,
			&userOne.Role,
			&userOne.Subscribe,
			&userOne.DateSubscribe,
		)
		if err != nil {
			return err
		}
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

func UserEditHandler(w http.ResponseWriter, r *http.Request) {
	db, err := configDB.InitDB()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
		return
	}
	update := Edit{}
	id, err := strconv.ParseInt(r.PostForm.Get("id"), 10, 0)
	firstname := r.PostForm.Get("firstname")
	lastname := r.PostForm.Get("lastname")
	email := r.PostForm.Get("email")

	if firstname == "" || lastname == "" || email == "" {
		http.Error(w, "Fields are empty", http.StatusBadRequest)
		return
	}

	sqlStatement := `
	UPDATE users SET firstname = $2, lastname = $3, email = $4
	WHERE id = $1`
	_, err = db.Exec(sqlStatement, id, firstname, lastname, email)
	if err != nil {
		update.ID = int(id)
		update.Success = false
	} else {
		update.ID = int(id)
		update.Success = true
	}

	out, err := json.Marshal(update)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, string(out))
	return
}