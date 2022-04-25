package utils

import (
	"log"

	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/configs"
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/errors"
	"golang.org/x/crypto/bcrypt"
)

var SECRET_KEY string

func init() {
	SECRET_KEY = configs.EnvJWTSecretKey()
}

type SignedDetails struct {
	UserId  string `json:"userId"`
	Name    string `json:"name"`
	EmailId string `json:"emailId"`
	IsAdmin bool   `json:"isAdmin"`
	jwt.StandardClaims
}

func CreateToken(id string, emailId, name string, isAdmin bool) (string, *errors.AppError) {
	claims := &SignedDetails{
		UserId:  id,
		Name:    name,
		EmailId: emailId,
		IsAdmin: isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		return "", errors.NewBadRequestError("Error while creating token")
	}

	return token, nil
}

func ValidateToken(tokenReceived string) (bool, *errors.AppError) {
	token, err := jwt.ParseWithClaims(
		tokenReceived,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		return false, errors.NewUnexpectedError("Error while validating token")
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		err_ := errors.NewBadRequestError("Error while validating token")
		return false, err_
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err_ := errors.NewBadRequestError("Your session has expired. Please re-login")
		return false, err_
	}

	return true, nil
}

func GetClaimsFromToken(tokenReceived string) (SignedDetails, *errors.AppError) {
	var claims SignedDetails
	token, err := jwt.ParseWithClaims(
		tokenReceived,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		return claims, errors.NewExpectationFailed("Error while validating token")
	}

	if token.Valid {
		return claims, nil
	}

	return claims, errors.NewBadRequestError("Token is invalid")
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, *errors.AppError) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))

	if err != nil {
		return false, errors.NewUnauthorisedError("Invalid password")
	}
	return true, nil
}
