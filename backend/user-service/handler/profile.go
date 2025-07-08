// profile.go - Kullanıcı profil işlemleri
package handler

import (
	"encoding/json"
	"eros/user-service/model"
	"eros/user-service/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProfileHandler struct {
	userService *service.UserService
}

func NewProfileHandler(userService *service.UserService) *ProfileHandler {
	return &ProfileHandler{
		userService: userService,
	}
}

// GetProfile - Kullanıcı profilini getir
func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateProfile - Kullanıcı profilini güncelle
func (h *ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user.ID = userID
	if err := h.userService.UpdateUser(&user); err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Profile updated successfully"})
}

// DeleteProfile - Kullanıcı profilini sil
func (h *ProfileHandler) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.userService.DeleteUser(userID); err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Profile deleted successfully"})
}

// GetUserPreferences - Kullanıcı tercihlerini getir
func (h *ProfileHandler) GetUserPreferences(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	preferences := map[string]interface{}{
		"age_range":   user.AgeRange,
		"distance":    user.Distance,
		"seriousness": user.Seriousness,
		"hobbies":     user.Hobbies,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(preferences)
}

// UpdateUserPreferences - Kullanıcı tercihlerini güncelle
func (h *ProfileHandler) UpdateUserPreferences(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var preferences map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&preferences); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Tercihleri güncelle
	if ageRange, ok := preferences["age_range"].(string); ok {
		user.AgeRange = ageRange
	}
	if distance, ok := preferences["distance"].(int); ok {
		user.Distance = distance
	}
	if seriousness, ok := preferences["seriousness"].(float64); ok {
		user.Seriousness = int(seriousness)
	}
	if hobbies, ok := preferences["hobbies"].([]string); ok {
		user.Hobbies = hobbies
	}

	if err := h.userService.UpdateUser(user); err != nil {
		http.Error(w, "Failed to update preferences", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Preferences updated successfully"})
}
