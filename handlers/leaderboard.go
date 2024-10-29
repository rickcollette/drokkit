package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"drokkit/models"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	redisClient *redis.Client
)

// InitLeaderboard initializes Redis and database
func InitLeaderboard(database *gorm.DB, rdb *redis.Client) {
	db = database
	redisClient = rdb
}

// GetLeaderboard handles various types of leaderboard views
func GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	leaderboardType := r.URL.Query().Get("type")
	timeFrame := r.URL.Query().Get("timeframe")

	// Default to "all-time" if no timeframe is provided
	if timeFrame == "" {
		timeFrame = "all-time"
	}

	key := "leaderboard:" + leaderboardType + ":" + timeFrame
	cache, err := redisClient.Get(key).Result()

	// Check for cached leaderboard data
	if err == nil {
		w.Write([]byte(cache))
		return
	}

	// If not in cache, query and cache the result
	leaderboard := fetchLeaderboardFromDB(leaderboardType, timeFrame)
	leaderboardJSON, err := json.Marshal(leaderboard)
	if err != nil {
		http.Error(w, "Failed to encode leaderboard", http.StatusInternalServerError)
		return
	}

	redisClient.Set(key, leaderboardJSON, 15*time.Minute)
	w.Write(leaderboardJSON)
}

// fetchLeaderboardFromDB queries the database based on the leaderboard type and timeframe
func fetchLeaderboardFromDB(leaderboardType, timeFrame string) []models.Stats {
	var stats []models.Stats

	// Modify query based on leaderboard type
	switch leaderboardType {
	case "individual":
		db.Where("team_id IS NULL").Order("experience DESC").Find(&stats)
	case "team":
		// Assuming a model or stats aggregation for team leaderboard
		db.Where("team_id IS NOT NULL").Order("experience DESC").Find(&stats)
	default:
		// Default to individual leaderboard if type is unrecognized
		db.Order("experience DESC").Find(&stats)
	}

	// Further filter by timeframe
	switch timeFrame {
	case "monthly":
		db.Where("created_at >= ?", time.Now().AddDate(0, -1, 0)).Order("experience DESC").Find(&stats)
	case "weekly":
		db.Where("created_at >= ?", time.Now().AddDate(0, 0, -7)).Order("experience DESC").Find(&stats)
	default:
		db.Order("experience DESC").Find(&stats)
	}

	return stats
}
