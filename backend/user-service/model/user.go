// user.go - Kullanıcı veri modeli
package model

import "time"

type User struct {
	ID              int       `json:"id" db:"id"`
	Name            string    `json:"name" db:"name"`
	Email           string    `json:"email" db:"email"`
	Password        string    `json:"-" db:"password"`
	Bio             string    `json:"bio" db:"bio"`
	Age             int       `json:"age" db:"age"`
	AgeRange        string    `json:"age_range" db:"age_range"`
	Distance        int       `json:"distance" db:"distance"`
	Seriousness     int       `json:"seriousness" db:"seriousness"`
	Height          int       `json:"height" db:"height"` // cm cinsinden
	Weight          int       `json:"weight" db:"weight"` // kg cinsinden
	Smokes          bool      `json:"smokes" db:"smokes"`
	Drinks          bool      `json:"drinks" db:"drinks"`
	Job             string    `json:"job" db:"job"`
	JobCategory     string    `json:"job_category" db:"job_category"`
	Education       string    `json:"education" db:"education"`
	Hobbies         []string  `json:"hobbies" db:"hobbies"`
	HobbyCategories []string  `json:"hobby_categories" db:"hobby_categories"`
	Photos          []Photo   `json:"photos"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

type UserPreferences struct {
	ID                       int      `json:"id" db:"id"`
	UserID                   int      `json:"user_id" db:"user_id"`
	MinAge                   int      `json:"min_age" db:"min_age"`
	MaxAge                   int      `json:"max_age" db:"max_age"`
	MinHeight                int      `json:"min_height" db:"min_height"`
	MaxHeight                int      `json:"max_height" db:"max_height"`
	AcceptsSmokers           bool     `json:"accepts_smokers" db:"accepts_smokers"`
	AcceptsDrinkers          bool     `json:"accepts_drinkers" db:"accepts_drinkers"`
	MinSeriousness           int      `json:"min_seriousness" db:"min_seriousness"`
	MaxSeriousness           int      `json:"max_seriousness" db:"max_seriousness"`
	PreferredJobCategories   []string `json:"preferred_job_categories" db:"preferred_job_categories"`
	PreferredHobbyCategories []string `json:"preferred_hobby_categories" db:"preferred_hobby_categories"`
}
