package repository

import (
	"database/sql"
	"encoding/json"
	"eros/user-service/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	hobbiesJSON, err := json.Marshal(user.Hobbies)
	if err != nil {
		return err
	}

	hobbyCategoriesJSON, err := json.Marshal(user.HobbyCategories)
	if err != nil {
		return err
	}

	result, err := r.db.Exec(`
		INSERT INTO users (name, email, password, bio, age, age_range, distance, seriousness, height, weight, smokes, drinks, job, job_category, education, hobbies, hobby_categories, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, user.Name, user.Email, user.Password, user.Bio, user.Age, user.AgeRange, user.Distance, user.Seriousness, user.Height, user.Weight, user.Smokes, user.Drinks, user.Job, user.JobCategory, user.Education, string(hobbiesJSON), string(hobbyCategoriesJSON), user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)
	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	var hobbiesJSON, hobbyCategoriesJSON string

	err := r.db.QueryRow(`
		SELECT id, name, email, password, bio, age, age_range, distance, seriousness, height, weight, smokes, drinks, job, job_category, education, hobbies, hobby_categories, created_at, updated_at
		FROM users WHERE email = ?
	`, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Bio, &user.Age, &user.AgeRange, &user.Distance, &user.Seriousness, &user.Height, &user.Weight, &user.Smokes, &user.Drinks, &user.Job, &user.JobCategory, &user.Education, &hobbiesJSON, &hobbyCategoriesJSON, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(hobbiesJSON), &user.Hobbies); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(hobbyCategoriesJSON), &user.HobbyCategories); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByID(userID int) (*model.User, error) {
	var user model.User
	var hobbiesJSON, hobbyCategoriesJSON string

	err := r.db.QueryRow(`
		SELECT id, name, email, password, bio, age, age_range, distance, seriousness, height, weight, smokes, drinks, job, job_category, education, hobbies, hobby_categories, created_at, updated_at
		FROM users WHERE id = ?
	`, userID).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Bio, &user.Age, &user.AgeRange, &user.Distance, &user.Seriousness, &user.Height, &user.Weight, &user.Smokes, &user.Drinks, &user.Job, &user.JobCategory, &user.Education, &hobbiesJSON, &hobbyCategoriesJSON, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(hobbiesJSON), &user.Hobbies); err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(hobbyCategoriesJSON), &user.HobbyCategories); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdateUser(user *model.User) error {
	hobbiesJSON, err := json.Marshal(user.Hobbies)
	if err != nil {
		return err
	}

	hobbyCategoriesJSON, err := json.Marshal(user.HobbyCategories)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(`
		UPDATE users SET name = ?, email = ?, bio = ?, age = ?, age_range = ?, distance = ?, seriousness = ?, height = ?, weight = ?, smokes = ?, drinks = ?, job = ?, job_category = ?, education = ?, hobbies = ?, hobby_categories = ?, updated_at = ?
		WHERE id = ?
	`, user.Name, user.Email, user.Bio, user.Age, user.AgeRange, user.Distance, user.Seriousness, user.Height, user.Weight, user.Smokes, user.Drinks, user.Job, user.JobCategory, user.Education, string(hobbiesJSON), string(hobbyCategoriesJSON), user.UpdatedAt, user.ID)

	return err
}

func (r *UserRepository) GetPotentialMatches(user *model.User, limit int) ([]model.User, error) {
	rows, err := r.db.Query(`
		SELECT id, name, email, password, bio, age, age_range, distance, seriousness, height, weight, smokes, drinks, job, job_category, education, hobbies, hobby_categories, created_at, updated_at
		FROM users 
		WHERE id != ? AND seriousness BETWEEN ? AND ?
		LIMIT ?
	`, user.ID, user.Seriousness-2, user.Seriousness+2, limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		var hobbiesJSON, hobbyCategoriesJSON string
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Bio, &u.Age, &u.AgeRange, &u.Distance, &u.Seriousness, &u.Height, &u.Weight, &u.Smokes, &u.Drinks, &u.Job, &u.JobCategory, &u.Education, &hobbiesJSON, &hobbyCategoriesJSON, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(hobbiesJSON), &u.Hobbies); err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(hobbyCategoriesJSON), &u.HobbyCategories); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

func (r *UserRepository) EmailExists(email string) (bool, error) {
	var exists int
	err := r.db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&exists)
	return exists > 0, err
}
