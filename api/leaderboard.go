package api

import (
	"encoding/json"
	"investor-api/db"
	"net/http"
	"sync"

	"investor-api/model"
)

var (
	leaderboard     = make(map[string]int) // name -> score
	leaderboardLock sync.RWMutex
)

func SubmitScoreHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var entry model.LeaderboardEntry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec(`
		INSERT INTO leaderboard (name, score)
		VALUES (?, ?)
		ON DUPLICATE KEY UPDATE score = GREATEST(score, VALUES(score))`,
		entry.Name, entry.Score,
	)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetLeaderboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.DB.Query(`
		SELECT name, score
		FROM leaderboard
		ORDER BY score DESC
		LIMIT 10
	`)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []model.LeaderboardEntry
	for rows.Next() {
		var entry model.LeaderboardEntry
		if err := rows.Scan(&entry.Name, &entry.Score); err != nil {
			http.Error(w, "Failed to parse result", http.StatusInternalServerError)
			return
		}
		results = append(results, entry)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
