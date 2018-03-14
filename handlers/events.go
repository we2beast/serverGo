package handlers

import (
	"net/http"
	"fmt"
	configDB "Calday-Server/config"
	"encoding/json"
	"time"
	"strconv"
)

type Event struct {
	ID int `json:"id"`
	UserId int `json:"user_id"`
	Title string `json:"title"`
	Text string `json:"text"`
	ListNotifications string `json:"list_notifications"`
	Notifications string `json:"notifications"`
	CreateAt string `json:"create_at"`
	Date string `json:"date"`
	Complete bool `json:"complete"`
	Important bool `json:"important"`
}

type Events struct {
	Event []Event
}

func EventHandler(w http.ResponseWriter, r *http.Request) {
	event := Event{}
	events := Events{}
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

	userId, err := parseParams(r, "/api/events/", 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	sqlStatement := `SELECT * FROM events WHERE user_id = $1 ORDER BY date ASC`

	rows, err := db.Query(sqlStatement, userId)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&event.ID,
			&event.UserId,
			&event.Title,
			&event.Text,
			&event.ListNotifications,
			&event.Notifications,
			&event.CreateAt,
			&event.Date,
			&event.Complete,
			&event.Important,
		)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		events.Event = append(events.Event, event)
	}
	err = rows.Err()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(events)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(out))
	return
}

func EventInsertHandler(w http.ResponseWriter, r *http.Request) {
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

	user_id := r.PostForm.Get("user_id")
	title := r.PostForm.Get("title")
	text := r.PostForm.Get("text")
	list_notifications := r.PostForm.Get("list_notifications")
	notification := r.PostForm.Get("notifications")
	date := r.PostForm.Get("date")
	complete := r.PostForm.Get("complete")
	important := r.PostForm.Get("important")

	if user_id == "" || title == "" || text == "" || list_notifications == "" || notification == "" || date == "" || complete == "" || important == "" {
		http.Error(w, "Fields are empty", http.StatusBadRequest)
		return
	}

	t, _ := time.Parse("2006-01-02 15:04:05.000000", date)

	sqlStatement := `
	INSERT INTO events (user_id, title, text, list_notifications, notifications, create_at, date, complete, important)
	VALUES ($1, $2, $3, $4, $5, NOW()::date, $6, $7, $8)
	RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, user_id, title, text, list_notifications, notification, t, complete, important).Scan(&id)
	if err != nil {
		http.Error(w, "Internal Server Error 2", http.StatusInternalServerError)
		return
	} else {
		event := Event{}
		sqlStatement := `SELECT id, user_id, title, text, list_notifications, notifications, create_at, date, complete, important FROM events WHERE id = $1`
		err = db.QueryRow(sqlStatement, id).Scan(&event.ID, &event.UserId, &event.Title, &event.Text, &event.ListNotifications, &event.Notifications, &event.CreateAt, &event.Date, &event.Complete, &event.Important)
		if err != nil {
			http.Error(w, "Internal Server Error 2", http.StatusInternalServerError)
			return
		}

		responseJSON, _ := json.Marshal(event)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(responseJSON))
	}
	return
}

func EventUpdateHandler(w http.ResponseWriter, r *http.Request) {
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

	id, err := strconv.ParseInt(r.PostForm.Get("id"), 10, 0)
	user_id := r.PostForm.Get("user_id")
	title := r.PostForm.Get("title")
	text := r.PostForm.Get("text")
	list_notifications := r.PostForm.Get("list_notifications")
	notification := r.PostForm.Get("notifications")
	date := r.PostForm.Get("date")
	complete := r.PostForm.Get("complete")
	important := r.PostForm.Get("important")

	if user_id == "" || title == "" || text == "" || list_notifications == "" || notification == "" || date == "" || complete == "" || important == "" {
		http.Error(w, "Fields are empty", http.StatusBadRequest)
		return
	}

	t, _ := time.Parse("2006-01-02 15:04:05.000000", date)

	sqlStatement := `
	UPDATE events SET title = $2, text = $3, list_notifications = $4, notifications = $5, date = $6, complete = $7, important = $8
	WHERE id = $1 AND user_id = $9
	RETURNING id`
	idUpdate := 0
	err = db.QueryRow(sqlStatement, id, title, text, list_notifications, notification, t, complete, important, user_id).Scan(&idUpdate)
	if err != nil {
		http.Error(w, "Event not found", http.StatusBadRequest)
		return
	} else {
		event := Event{}
		sqlStatement := `SELECT * FROM events WHERE id = $1`
		err = db.QueryRow(sqlStatement, idUpdate).Scan(&event.ID, &event.UserId, &event.Title, &event.Text, &event.ListNotifications, &event.Notifications, &event.CreateAt, &event.Date, &event.Complete, &event.Important)

		responseJSON, _ := json.Marshal(event)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(responseJSON))
	}
	return
}
