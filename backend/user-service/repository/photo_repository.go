package repository

import (
	"database/sql"
	"time"
)

type Photo struct {
	ID         int
	UserID     int
	URL        string
	IsPrimary  bool
	OrderIndex int
	AIScore    float64
	IsVerified bool
	CreatedAt  time.Time
}

type PhotoRepository struct {
	db *sql.DB
}

func NewPhotoRepository(db *sql.DB) *PhotoRepository {
	return &PhotoRepository{db: db}
}

func (r *PhotoRepository) AddPhoto(photo *Photo) error {
	_, err := r.db.Exec(`
		INSERT INTO photos (user_id, url, is_primary, order_index, ai_score, is_verified, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, photo.UserID, photo.URL, photo.IsPrimary, photo.OrderIndex, photo.AIScore, photo.IsVerified, photo.CreatedAt)
	return err
}

func (r *PhotoRepository) GetPhotosByUser(userID int) ([]Photo, error) {
	rows, err := r.db.Query(`
		SELECT id, user_id, url, is_primary, order_index, ai_score, is_verified, created_at
		FROM photos WHERE user_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photos []Photo
	for rows.Next() {
		var p Photo
		err := rows.Scan(&p.ID, &p.UserID, &p.URL, &p.IsPrimary, &p.OrderIndex, &p.AIScore, &p.IsVerified, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		photos = append(photos, p)
	}
	return photos, nil
}
