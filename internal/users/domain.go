package users

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	ID         uint      `json:"id"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Profile    Profile   `json:"profile"`
	CratedAt   time.Time `json:"crated_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

type Profile struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Bio        string    `json:"bio"`
	CratedAt   time.Time `json:"crated_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

type CreateUser struct {
	Email    string `validate:"required,min=1,max=50" json:"email"`
	Password string `validate:"required,min=8,max=20" json:"password"`
}

type UpdateUserEmail struct {
	ID    uint   `validate:"required" json:"id"`
	Email string `validate:"required,min=1,max=50" json:"email"`
}

type UpdateUserPassword struct {
	ID           uint   `validate:"required" json:"id"`
	LastPassword string `validate:"required,min=8,max=20" json:"last_password"`
	NewPassword  string `validate:"required,min=8,max=20" json:"new_password"`
}

type UpdateUserProfile struct {
	ID   uint   `validate:"required" json:"id"`
	Name string `validate:"required,min=8,max=50" json:"name"`
	Bio  string `validate:"required,min=8,max=100" json:"bio"`
}

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user User) User
	Update(ctx context.Context, tx *sql.Tx, user User) User
	Delete(ctx context.Context, tx *sql.Tx, user User) User
	FindById(ctx context.Context, tx *sql.Tx, userId uint) (User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []User
}

type UserService interface {
	Create(ctx context.Context)
	Update(ctx context.Context)
	Delete(ctx context.Context)
	FindById(ctx context.Context)
	FindAll(ctx context.Context)
}
