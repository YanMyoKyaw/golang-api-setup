package user

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"test/database"
	"test/model"
	"time"
)

var key = []byte("1900")

type Claims struct {
	Username string
	jwt.StandardClaims
}

func Register(user *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(hashedPassword)
	id, _ := uuid.NewUUID()
	user.ID = id
	database.CreateUser(user)
	return nil
}

func LogIn(user *model.User) (*model.JWTToken, error) {

	dbUser := database.GetUserByMail(user.Email)
	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if dbUser.Email != user.Email || err != nil {
		return &model.JWTToken{}, errors.New("Incorrect information")
	}
	expirationTime := time.Now().Add(10 * time.Minute)
	claim := &Claims{
		Username: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return &model.JWTToken{}, err
	}
	return &model.JWTToken{
		Name:   "token",
		Token:  tokenStr,
		Expire: expirationTime,
	}, nil
}
