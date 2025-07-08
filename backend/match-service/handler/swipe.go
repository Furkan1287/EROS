// swipe.go - Swipe işlemleri (Klasik Tinder tarzı)
package handler

import (
    "encoding/json"
    "net/http"
    "strconv"
    "eros/match-service/service"
    "eros/match-service/model"
)

type SwipeHandler struct {
    matchService *service.MatchService
}

func NewSwipeHandler(matchService *service.MatchService) *SwipeHandler {
    return &SwipeHandler{matchService: matchService}
}

// SwipeRequest - Swipe isteği
type SwipeRequest struct {
    UserID     int  `json:"user_id"`
    TargetID   int  `json:"target_id"`
    Direction  string `json:"direction"` // "right" veya "left"
}

// SwipeResponse - Swipe yanıtı
type SwipeResponse struct {
    IsMatch     bool   `json:"is_match"`
    Message     string `json:"message"`
    DateTask    *model.DateTask `json:"date_task,omitempty"`
}

// Swipe - Kullanıcı swipe yapar
func (h *SwipeHandler) Swipe(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req SwipeRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Swipe yönü kontrolü
    if req.Direction != "right" && req.Direction != "left" {
        http.Error(w, "Invalid direction", http.StatusBadRequest)
        return
    }

    // Swipe işlemini gerçekleştir
    isMatch, err := h.matchService.ProcessSwipe(req.UserID, req.TargetID, req.Direction)
    if err != nil {
        http.Error(w, "Swipe processing failed", http.StatusInternalServerError)
        return
    }

    response := SwipeResponse{
        IsMatch: isMatch,
        Message: "Swipe processed successfully",
    }

    // Eğer eşleşme varsa, date görevi öner
    if isMatch {
        dateTask, err := h.matchService.GenerateDateTask(req.UserID, req.TargetID)
        if err == nil {
            response.DateTask = dateTask
        }
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

// GetPotentialMatches - Potansiyel eşleşmeleri getir
func (h *SwipeHandler) GetPotentialMatches(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    userIDStr := r.URL.Query().Get("user_id")
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    limitStr := r.URL.Query().Get("limit")
    limit := 10 // varsayılan
    if limitStr != "" {
        if l, err := strconv.Atoi(limitStr); err == nil {
            limit = l
        }
    }

    matches, err := h.matchService.GetPotentialMatches(userID, limit)
    if err != nil {
        http.Error(w, "Failed to get potential matches", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(matches)
}

// GetMatchHistory - Eşleşme geçmişini getir
func (h *SwipeHandler) GetMatchHistory(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    userIDStr := r.URL.Query().Get("user_id")
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    history, err := h.matchService.GetMatchHistory(userID)
    if err != nil {
        http.Error(w, "Failed to get match history", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(history)
}
