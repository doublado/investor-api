package main

import (
	"investor-api/db"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"investor-api/api"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Init()

	requiredEnv := map[string]string{
		"OPENAI_API_KEY": os.Getenv("OPENAI_API_KEY"),
		"API_SECRET":     os.Getenv("API_SECRET"),
	}

	for key, val := range requiredEnv {
		if val == "" || val == "YOUR_API_KEY" || val == "YOUR_API_SECRET" {
			log.Fatalf("Environment variable %s is not set or is using a placeholder", key)
		}
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/evaluate", api.RequireAuth(api.EvaluateHandler))
	mux.HandleFunc("/scenario", api.RequireAuth(api.ScenarioHandler))
	mux.HandleFunc("/leaderboard", api.RequireAuth(api.GetLeaderboardHandler))
	mux.HandleFunc("/leaderboard/submit", api.RequireAuth(api.SubmitScoreHandler))

	addr := ":8080"
	log.Printf("Server started at %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
