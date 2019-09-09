package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"

	"github.com/obasajujoshua31/blogos/api/security"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null;" json:"name"`
	Email     string    `gorm:"type:varchar(255);not null;unique" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp" json:"updated_at"`
	Posts     []Post    `gorm:"foreignkey:AuthorID" json:"posts,omitempty"`
}

func (u *User) BeforeSave() error {
	hashPassword, err := security.Hash(u.Password)

	if err != nil {
		return err
	}
	u.Password = string(hashPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Password = html.EscapeString(strings.TrimSpace(u.Password))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	case "login":
		if u.Email == "" {
			return errors.New("Required Email")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		if u.Password == "" {
			return errors.New("Password is required")
		}
		return nil

	default:
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *User) ConfirmPassword(password string) error {
	return security.VerifyPassword(u.Password, password)

}
