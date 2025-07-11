// user.go - Kullanıcı iş mantığı
package service

import (
	"eros/user-service/model"
	"eros/user-service/repository"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo  *repository.UserRepository
	photoRepo *repository.PhotoRepository
}

func NewUserService(userRepo *repository.UserRepository, photoRepo *repository.PhotoRepository) *UserService {
	return &UserService{
		userRepo:  userRepo,
		photoRepo: photoRepo,
	}
}

// CreateUser - Yeni kullanıcı oluştur
func (s *UserService) CreateUser(user *model.User) error {
	// Email kontrolü
	exists, err := s.userRepo.EmailExists(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email already exists")
	}

	// Minimum profil bilgisi kontrolü
	if user.Name == "" || user.Email == "" {
		return errors.New("name and email are required")
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return s.userRepo.CreateUser(user)
}

// AuthenticateUser - Kullanıcı kimlik doğrulama
func (s *UserService) AuthenticateUser(email, password string) (*model.User, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

// GetUserByID - ID ile kullanıcı getir
func (s *UserService) GetUserByID(userID int) (*model.User, error) {
	return s.userRepo.GetUserByID(userID)
}

// UpdateUser - Kullanıcı bilgilerini güncelle
func (s *UserService) UpdateUser(user *model.User) error {
	user.UpdatedAt = time.Now()
	return s.userRepo.UpdateUser(user)
}

// AddPhoto - Fotoğraf ekle
func (s *UserService) AddPhoto(photo *model.Photo) error {
	photo.CreatedAt = time.Now()
	repoPhoto := &repository.Photo{
		ID:         photo.ID,
		UserID:     photo.UserID,
		URL:        photo.URL,
		IsPrimary:  photo.IsPrimary,
		OrderIndex: photo.OrderIndex,
		AIScore:    photo.AIScore,
		IsVerified: photo.IsVerified,
		CreatedAt:  photo.CreatedAt,
	}
	return s.photoRepo.AddPhoto(repoPhoto)
}

// GetUserPhotos - Kullanıcının fotoğraflarını getir
func (s *UserService) GetUserPhotos(userID int) ([]model.Photo, error) {
	repoPhotos, err := s.photoRepo.GetPhotosByUser(userID)
	if err != nil {
		return nil, err
	}
	var photos []model.Photo
	for _, rp := range repoPhotos {
		photos = append(photos, model.Photo{
			ID:         rp.ID,
			UserID:     rp.UserID,
			URL:        rp.URL,
			IsPrimary:  rp.IsPrimary,
			OrderIndex: rp.OrderIndex,
			AIScore:    rp.AIScore,
			IsVerified: rp.IsVerified,
			CreatedAt:  rp.CreatedAt,
		})
	}
	return photos, nil
}

// GetPotentialMatches - Potansiyel eşleşmeleri getir
func (s *UserService) GetPotentialMatches(userID int, limit int) ([]model.User, error) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return s.userRepo.GetPotentialMatches(user, limit)
}

// GetUserByEmail - Email ile kullanıcı getir
func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	return s.userRepo.GetUserByEmail(email)
}

// EmailExists - Email kontrolü
func (s *UserService) EmailExists(email string) (bool, error) {
	return s.userRepo.EmailExists(email)
}

// Kullanıcının fotoğraf sayısını döndür (stub)
func (s *UserService) GetUserPhotoCount(userID int) (int, error) {
	return 0, nil
}

// Fotoğrafı analiz et (stub)
func (s *UserService) AnalyzePhoto(photoURL string) (float64, error) {
	return 0.0, nil
}

// Fotoğrafları sırala (stub)
func (s *UserService) ReorderPhotos(userID int, order []int) error {
	return nil
}

// Fotoğrafı sil (stub)
func (s *UserService) DeletePhoto(photoID int) error {
	return nil
}

// Kullanıcıyı sil (stub)
func (s *UserService) DeleteUser(userID int) error {
	return nil
}
