// user.go - Kullanıcı veri modeli (Match Service için)
package model

import "time"

type User struct {
    ID              int       `json:"id" db:"id"`
    Name            string    `json:"name" db:"name"`
    Email           string    `json:"email" db:"email"`
    Bio             string    `json:"bio" db:"bio"`
    Age             int       `json:"age" db:"age"`
    AgeRange        string    `json:"age_range" db:"age_range"`
    Distance        int       `json:"distance" db:"distance"`
    Seriousness     string    `json:"seriousness" db:"seriousness"`
    Height          int       `json:"height" db:"height"`
    Weight          int       `json:"weight" db:"weight"`
    Smokes          bool      `json:"smokes" db:"smokes"`
    Drinks          bool      `json:"drinks" db:"drinks"`
    Job             string    `json:"job" db:"job"`
    JobCategory     string    `json:"job_category" db:"job_category"`
    Education       string    `json:"education" db:"education"`
    Hobbies         []string  `json:"hobbies" db:"hobbies"`
    HobbyCategories []string  `json:"hobby_categories" db:"hobby_categories"`
    CreatedAt       time.Time `json:"created_at" db:"created_at"`
    UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
} 