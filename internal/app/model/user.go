package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID 	  			  uuid.UUID `json:"id"`
	Email 			  string	`json:"email"`
	Username 		  string	`json:"username"`
	Password		  string	`json:"password,omitempty"`
	EncryptedPassword string	`json:"-"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Username, validation.Required),
		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(8, 100)),
	)
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}


	return nil
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil;
}

func (u *User) CreateJWT(key []byte) (string, error) {
	exp := time.Now().Add(24 * time.Hour)
	c := &Claims{
		u.ID,
		jwt.StandardClaims{
			Subject: u.Email,
			ExpiresAt: exp.Unix(),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(key);
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}