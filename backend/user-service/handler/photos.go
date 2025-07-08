// photos.go - Kullanıcı fotoğraf işlemleri
package handler

import (
    "encoding/json"
    "net/http"
    "eros/user-service/model"
    "eros/user-service/service"
    "strconv"
)

type PhotosHandler struct {
    userService *service.UserService
}

func NewPhotosHandler(userService *service.UserService) *PhotosHandler {
    return &PhotosHandler{userService: userService}
}

// UploadPhotoRequest - Fotoğraf yükleme isteği
type UploadPhotoRequest struct {
    UserID    int    `json:"user_id"`
    ImageData string `json:"image_data"` // Base64 encoded image
    OrderIndex int   `json:"order_index"`
}

// ReorderPhotosRequest - Fotoğraf sıralama isteği
type ReorderPhotosRequest struct {
    UserID int   `json:"user_id"`
    PhotoIDs []int `json:"photo_ids"` // Yeni sıralama
}

// UploadPhoto - Fotoğraf yükleme (maksimum 6 fotoğraf)
func (h *PhotosHandler) UploadPhoto(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req UploadPhotoRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Fotoğraf sayısı kontrolü
    photoCount, err := h.userService.GetUserPhotoCount(req.UserID)
    if err != nil {
        http.Error(w, "Failed to get photo count", http.StatusInternalServerError)
        return
    }

    if photoCount >= 6 {
        http.Error(w, "Maximum 6 photos allowed", http.StatusBadRequest)
        return
    }

    // AI analizi için fotoğrafı işle
    aiScore, err := h.userService.AnalyzePhoto(req.ImageData)
    if err != nil {
        http.Error(w, "Photo analysis failed", http.StatusInternalServerError)
        return
    }

    photo := &model.Photo{
        UserID:    req.UserID,
        URL:       req.ImageData, // Gerçek uygulamada cloud storage'a yüklenir
        OrderIndex: req.OrderIndex,
        AIScore:   aiScore,
        IsVerified: true, // TODO: AI ile doğrulama
    }

    if err := h.userService.AddPhoto(photo); err != nil {
        http.Error(w, "Failed to save photo", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Photo uploaded successfully",
        "photo_id": photo.ID,
        "ai_score": photo.AIScore,
    })
}

// GetUserPhotos - Kullanıcının fotoğraflarını getir
func (h *PhotosHandler) GetUserPhotos(w http.ResponseWriter, r *http.Request) {
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

    photos, err := h.userService.GetUserPhotos(userID)
    if err != nil {
        http.Error(w, "Failed to get photos", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(photos)
}

// ReorderPhotos - Fotoğraf sıralamasını değiştir
func (h *PhotosHandler) ReorderPhotos(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var req ReorderPhotosRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if err := h.userService.ReorderPhotos(req.UserID, req.PhotoIDs); err != nil {
        http.Error(w, "Failed to reorder photos", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Photos reordered successfully",
    })
}

// DeletePhoto - Fotoğraf sil
func (h *PhotosHandler) DeletePhoto(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    photoIDStr := r.URL.Query().Get("photo_id")
    photoID, err := strconv.Atoi(photoIDStr)
    if err != nil {
        http.Error(w, "Invalid photo ID", http.StatusBadRequest)
        return
    }

    if err := h.userService.DeletePhoto(photoID); err != nil {
        http.Error(w, "Failed to delete photo", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Photo deleted successfully",
    })
}
