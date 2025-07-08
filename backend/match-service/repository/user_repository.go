// user_repository.go - Kullanıcı veritabanı işlemleri (Match Service için)
package repository

import (
    "database/sql"
    "eros/match-service/model"
    "encoding/json"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db: db}
}

// GetUserByID - ID ile kullanıcı getir
func (r *UserRepository) GetUserByID(userID int) (*model.User, error) {
    query := `
        SELECT id, name, email, bio, age, age_range, distance, seriousness,
               height, weight, smokes, drinks, job, job_category, education,
               hobbies, hobby_categories, created_at, updated_at
        FROM users WHERE id = ?
    `
    
    user := &model.User{}
    var hobbiesStr, hobbyCategoriesStr string
    
    err := r.db.QueryRow(query, userID).Scan(
        &user.ID, &user.Name, &user.Email, &user.Bio, &user.Age, &user.AgeRange,
        &user.Distance, &user.Seriousness, &user.Height, &user.Weight,
        &user.Smokes, &user.Drinks, &user.Job, &user.JobCategory, &user.Education,
        &hobbiesStr, &hobbyCategoriesStr, &user.CreatedAt, &user.UpdatedAt,
    )
    
    if err != nil {
        return nil, err
    }
    
    // JSON string'leri array'e çevir
    if hobbiesStr != "" {
        json.Unmarshal([]byte(hobbiesStr), &user.Hobbies)
    }
    if hobbyCategoriesStr != "" {
        json.Unmarshal([]byte(hobbyCategoriesStr), &user.HobbyCategories)
    }
    
    return user, nil
}

// GetPotentialMatches - Potansiyel eşleşmeleri getir
func (r *UserRepository) GetPotentialMatches(user *model.User, limit int) ([]model.User, error) {
    query := `
        SELECT id, name, email, bio, age, age_range, distance, seriousness,
               height, weight, smokes, drinks, job, job_category, education,
               hobbies, hobby_categories, created_at, updated_at
        FROM users 
        WHERE id != ? AND age BETWEEN ? AND ?
        ORDER BY RANDOM() LIMIT ?
    `
    
    rows, err := r.db.Query(query, user.ID, 18, 100, limit)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var users []model.User
    for rows.Next() {
        var user model.User
        var hobbiesStr, hobbyCategoriesStr string
        
        err := rows.Scan(
            &user.ID, &user.Name, &user.Email, &user.Bio, &user.Age, &user.AgeRange,
            &user.Distance, &user.Seriousness, &user.Height, &user.Weight,
            &user.Smokes, &user.Drinks, &user.Job, &user.JobCategory, &user.Education,
            &hobbiesStr, &hobbyCategoriesStr, &user.CreatedAt, &user.UpdatedAt,
        )
        
        if err != nil {
            return nil, err
        }
        
        // JSON string'leri array'e çevir
        if hobbiesStr != "" {
            json.Unmarshal([]byte(hobbiesStr), &user.Hobbies)
        }
        if hobbyCategoriesStr != "" {
            json.Unmarshal([]byte(hobbyCategoriesStr), &user.HobbyCategories)
        }
        
        users = append(users, user)
    }
    
    return users, nil
}

// GetUsersByIDs - ID listesi ile kullanıcıları getir
func (r *UserRepository) GetUsersByIDs(userIDs []int) ([]model.User, error) {
    if len(userIDs) == 0 {
        return []model.User{}, nil
    }
    
    // IN clause için placeholder'lar oluştur
    placeholders := ""
    args := make([]interface{}, len(userIDs))
    for i, id := range userIDs {
        if i > 0 {
            placeholders += ","
        }
        placeholders += "?"
        args[i] = id
    }
    
    query := `
        SELECT id, name, email, bio, age, age_range, distance, seriousness,
               height, weight, smokes, drinks, job, job_category, education,
               hobbies, hobby_categories, created_at, updated_at
        FROM users WHERE id IN (` + placeholders + `)
    `
    
    rows, err := r.db.Query(query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var users []model.User
    for rows.Next() {
        var user model.User
        var hobbiesStr, hobbyCategoriesStr string
        
        err := rows.Scan(
            &user.ID, &user.Name, &user.Email, &user.Bio, &user.Age, &user.AgeRange,
            &user.Distance, &user.Seriousness, &user.Height, &user.Weight,
            &user.Smokes, &user.Drinks, &user.Job, &user.JobCategory, &user.Education,
            &hobbiesStr, &hobbyCategoriesStr, &user.CreatedAt, &user.UpdatedAt,
        )
        
        if err != nil {
            return nil, err
        }
        
        // JSON string'leri array'e çevir
        if hobbiesStr != "" {
            json.Unmarshal([]byte(hobbiesStr), &user.Hobbies)
        }
        if hobbyCategoriesStr != "" {
            json.Unmarshal([]byte(hobbyCategoriesStr), &user.HobbyCategories)
        }
        
        users = append(users, user)
    }
    
    return users, nil
} 