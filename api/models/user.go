package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/badoux/checkmail"
)

//user struct fields used in db
type User struct {
	UserID    uint64    `gorm:"primary_key;auto_increment" json:"user_id"`
	Name      string    `gorm:"size:255;not null;" json:"name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

//user prepare values to insert or update
func (u *User) Prepare() {
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Password = html.EscapeString(strings.TrimSpace(u.Password))
	u.CreatedAt = time.Now()
}

// validate values
func (u *User) Validate(action string) error {
	if u.Name == "" {
		return errors.New("required Name")
	}
	if u.Password == "" {
		return errors.New("required password")
	}
	if u.Email == "" {
		return errors.New("required email")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("invalid email")
	}
	return nil
}

//Crate user
func (u *User) SaveUser(db *gorm.DB) error {
	result := db.Where("email = ?", u.Email).FirstOrCreate(&u)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrInvalidData
	}
	return nil
}

//find user from db
func (u *User) FindUser(db *gorm.DB) error {
	if result := db.Model(&u).Where("user_id = ?", u.UserID).First(&u); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("User with that id does not exist")
		}
		return result.Error
	}
	return nil
}

//Find All users
func FindAllUsers(db *gorm.DB) ([]User, error) {
	var users []User

	if result := db.Model(User{}).Find(&users); result.Error != nil {
		return []User{}, result.Error
	}
	return users, nil
}

//update user
func (u *User) UpdateUser(m map[string]interface{}, db *gorm.DB) error {

	if result := db.Model(User{}).Where("user_id = ?", u.UserID).Updates(m); result.Error != nil {

		return result.Error
	}
	return nil
}

//Delete User
func (u *User) DeleteUser(db *gorm.DB) error {

	if result := db.Model(User{}).Where("user_id = ?", u.UserID).Delete(&u); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("User with that id does not exist")
		}
		return result.Error
	}
	return nil
}
