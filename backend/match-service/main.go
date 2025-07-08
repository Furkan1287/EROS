// main.go - Match Service Entry Point
package main

import (
    "log"
    "net/http"
    "os"
    "eros/match-service/handler"
    "eros/match-service/repository"
    "eros/match-service/service"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
)

func main() {
    // .env dosyasını yükle
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using default values")
    }

    // SQLite veritabanını başlat
    dbPath := os.Getenv("DB_PATH")
    if dbPath == "" {
        dbPath = "./eros_match.db"
    }

    db, err := repository.NewSQLiteDB(dbPath)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Close()

    // Veritabanı tablolarını oluştur
    if err := repository.InitMatchDatabase(db); err != nil {
        log.Fatal("Failed to initialize database:", err)
    }

    // Repository'leri oluştur
    matchRepo := repository.NewMatchRepository(db)
    userRepo := repository.NewUserRepository(db)
    
    // AI Service'i oluştur
    aiService := service.NewAIService()

    // Service'leri oluştur
    matchService := service.NewMatchService(matchRepo, userRepo, aiService)

    // Handler'ları oluştur
    swipeHandler := handler.NewSwipeHandler(matchService)
    blindHandler := handler.NewBlindHandler(matchService)

    // Router'ı oluştur
    router := mux.NewRouter()

    // Swipe routes (Klasik Tinder tarzı)
    router.HandleFunc("/api/swipe", swipeHandler.Swipe).Methods("POST")
    router.HandleFunc("/api/matches/potential", swipeHandler.GetPotentialMatches).Methods("GET")
    router.HandleFunc("/api/matches/history", swipeHandler.GetMatchHistory).Methods("GET")

    // Blind date routes
    router.HandleFunc("/api/blind/request", blindHandler.RequestBlindMatch).Methods("POST")
    router.HandleFunc("/api/blind/message", blindHandler.SendBlindMessage).Methods("POST")
    router.HandleFunc("/api/blind/messages", blindHandler.GetBlindMessages).Methods("GET")
    router.HandleFunc("/api/blind/complete", blindHandler.CompleteBlindDate).Methods("POST")
    router.HandleFunc("/api/blind/status", blindHandler.GetBlindMatchStatus).Methods("GET")

    // CORS middleware
    router.Use(func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Access-Control-Allow-Origin", "*")
            w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
            
            if r.Method == "OPTIONS" {
                w.WriteHeader(http.StatusOK)
                return
            }
            
            next.ServeHTTP(w, r)
        })
    })

    // Sunucuyu başlat
    port := os.Getenv("PORT")
    if port == "" {
        port = "8082"
    }

    log.Printf("Match Service starting on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}
