package routes

import (
	"drokkit/handlers"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenCookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := tokenCookie.Value
		claims := &handlers.Claims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return handlers.JwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/register", handlers.RegisterPlayer).Methods("POST")
	router.HandleFunc("/login", handlers.LoginPlayer).Methods("POST")

	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(AuthMiddleware)
	protected.HandleFunc("/match", handlers.CreateMatch).Methods("POST")
	protected.HandleFunc("/faction", handlers.CreateFaction).Methods("POST")
	protected.HandleFunc("/alliance", handlers.CreateAlliance).Methods("POST")
	protected.HandleFunc("/resource", handlers.UpdateResource).Methods("POST")

	router.HandleFunc("/ws/play", handlers.WebSocketHandler).Methods("GET")

	admin := router.PathPrefix("/admin").Subrouter()
	admin.Use(AuthMiddleware)
	admin.HandleFunc("/create", handlers.CreateAdmin).Methods("POST")
	admin.HandleFunc("/delete-player", handlers.DeletePlayer).Methods("DELETE")

	router.HandleFunc("/leaderboard", handlers.GetLeaderboard).Methods("GET")

	return router
}
