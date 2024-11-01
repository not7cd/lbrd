package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"log"
	"time"
)

type LeaderBoard struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (t *LeaderBoard) Scan(value interface{}) error {
	return json.Unmarshal([]byte(value.(string)), t)
}

func (t *LeaderBoard) Marshal() (driver.Value, error) {
	b, err := json.Marshal(t)
	return string(b), err
}

func InsertLeaderBoard(db *sql.DB, lb LeaderBoard) error {

	_, err := db.Exec("INSERT INTO leaderboard (name) VALUES (?, ?)", lb.Name)
	return err
}

type Score struct {
	Id       int       `json:"id"`
	Player   string    `json:"player"`
	Value    int       `json:"value"`
	LastEdit time.Time `json:"lastEdit"`
}

func (t *Score) Scan(value interface{}) error {
	return json.Unmarshal([]byte(value.(string)), t)
}

func (t *Score) Marshal() (driver.Value, error) {
	b, err := json.Marshal(t)
	return string(b), err
}

func nowUnixTimestamp() int32 {
	return int32(time.Now().Unix())
}

func InsertScore(db *sql.DB, score Score) error {
	_, err := db.Exec(
		"INSERT INTO score (player, value, last_change) VALUES (?, ?, ?)",
		score.Player, score.Value, nowUnixTimestamp(),
	)
	return err
}

func GetScores(db *sql.DB, leaderboard_id int) ([]Score, error) {
	rows, err := db.Query("SELECT id, player, value, last_edit FROM score WHERE leaderboard_id = ?", leaderboard_id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var scores []Score

	for rows.Next() {
		var score Score
		var timestamp string
		err = rows.Scan(&score.Id, &score.Player, &score.Value, &timestamp)
		if err != nil {
			log.Fatal(err)
		}
		score.LastEdit, err = time.Parse(time.DateTime, timestamp)
		if err != nil {
			log.Fatal(err)
		}
		scores = append(scores, score)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return scores, err
}
