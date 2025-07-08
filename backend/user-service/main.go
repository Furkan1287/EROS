// main.go - User Service Entry Point
package main

import (
	"eros/user-service/handler"
	"eros/user-service/repository"
	"eros/user-service/service"
	"log"
	"net/http"
	"os"

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
		dbPath = "./eros.db"
	}

	db, err := repository.NewSQLiteDB(dbPath)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Veritabanı tablolarını oluştur
	if err := repository.InitDatabase(db); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Repository'leri oluştur
	userRepo := repository.NewUserRepository(db)
	photoRepo := repository.NewPhotoRepository(db)

	// Service'leri oluştur
	userService := service.NewUserService(userRepo, photoRepo)

	// Handler'ları oluştur
	authHandler := handler.NewAuthHandler(userService)
	photosHandler := handler.NewPhotosHandler(userService)
	profileHandler := handler.NewProfileHandler(userService)

	// Router'ı oluştur
	router := mux.NewRouter()

	// Auth routes
	router.HandleFunc("/api/auth/register", authHandler.Register).Methods("POST")
	router.HandleFunc("/api/auth/simple-register", authHandler.SimpleRegister).Methods("POST")
	router.HandleFunc("/api/auth/login", authHandler.Login).Methods("POST")

	// Form data routes
	router.HandleFunc("/api/form/hobby-categories", authHandler.GetHobbyCategories).Methods("GET")
	router.HandleFunc("/api/form/education-levels", authHandler.GetEducationLevels).Methods("GET")
	router.HandleFunc("/api/form/job-categories", authHandler.GetJobCategories).Methods("GET")

	// Photo routes
	router.HandleFunc("/api/photos/upload", photosHandler.UploadPhoto).Methods("POST")
	router.HandleFunc("/api/photos", photosHandler.GetUserPhotos).Methods("GET")
	router.HandleFunc("/api/photos/reorder", photosHandler.ReorderPhotos).Methods("PUT")
	router.HandleFunc("/api/photos", photosHandler.DeletePhoto).Methods("DELETE")

	// User routes
	router.HandleFunc("/api/users/{id}", profileHandler.GetProfile).Methods("GET")
	router.HandleFunc("/api/users/{id}", profileHandler.UpdateProfile).Methods("PUT")
	router.HandleFunc("/api/users/{id}", profileHandler.DeleteProfile).Methods("DELETE")
	router.HandleFunc("/api/users/{id}/preferences", profileHandler.GetUserPreferences).Methods("GET")
	router.HandleFunc("/api/users/{id}/preferences", profileHandler.UpdateUserPreferences).Methods("PUT")

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
		port = "8081"
	}

	log.Printf("User Service starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
